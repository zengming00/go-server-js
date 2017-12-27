package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	_ "github.com/zengming00/go-server-js/lib"
	_ "github.com/zengming00/go-server-js/lib/db"
	_ "github.com/zengming00/go-server-js/lib/db/redis"
	mhttp "github.com/zengming00/go-server-js/lib/http"

	_ "net/http/pprof"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
	_ "github.com/go-sql-driver/mysql"
)

var _cwd string

type _server struct {
	runtime          *goja.Runtime
	registry         *require.Registry
	writeResultValue bool
}

func init() {
	var err error
	_cwd, err = os.Getwd()
	handErr(err)
}

func handErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (This *_server) handler(w http.ResponseWriter, r *http.Request) {
	// 加上这行后，内存的使用情况好了一些
	runtime.GC()
	u := r.URL
	file := filepath.Join(_cwd, u.Path)
	if strings.HasPrefix(file, _cwd) {
		ext := filepath.Ext(file)
		if ext == ".js" {
			runtime := This.runtime
			registry := This.registry

			if runtime == nil {
				runtime = goja.New()
			}
			if registry == nil {
				registry = new(require.Registry)
			}
			response := mhttp.NewResponse(runtime, &w)
			request := mhttp.NewRequest(runtime, r)
			runtime.Set("response", response)
			runtime.Set("request", request)

			ret, err := runFile(file, runtime, registry)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				panic(err)
			}
			// if goja.IsNull(*ret) || goja.IsUndefined(*ret) {
			// 	w.WriteHeader(http.StatusOK)
			// 	return
			// }
			// w.Header().Set("Content-Type", "text/html; charset=utf-8")
			if This.writeResultValue {
				w.Write([]byte((*ret).String()))
			}
			return
		}
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func server() {
	s := &_server{
	// registry: new(require.Registry),
	}
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/", s.handler)
	err := http.ListenAndServe(":8080", nil)
	handErr(err)
}

func main() {
	if len(os.Args) == 2 {
		filename := os.Args[1]
		start := time.Now()
		// this can be shared by multiple runtimes
		registry := new(require.Registry)
		for i := 0; i < 1000; i++ {
			runtime := goja.New()
			r, err := runFile(filename, runtime, registry)
			if err != nil {
				switch err := err.(type) {
				case *goja.Exception:
					fmt.Println(err.String())
				case *goja.InterruptedError:
					fmt.Println(err.String())
				default:
					fmt.Println(err)
				}
				os.Exit(64)
			}
			fmt.Println(*r)
		}
		fmt.Printf("done (%s)", time.Since(start))
	}
	server()
}

func runFile(filename string, runtime *goja.Runtime, registry *require.Registry) (*goja.Value, error) {
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	registry.Enable(runtime)
	console.Enable(runtime)

	time.AfterFunc(60*time.Second, func() {
		runtime.Interrupt("run code timeout, halt")
	})

	prg, err := goja.Compile(filename, string(src), false)
	if err != nil {
		return nil, err
	}

	result, err := runtime.RunProgram(prg)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
