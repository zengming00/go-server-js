package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	_ "github.com/zengming00/go-server-js/lib"
	_ "github.com/zengming00/go-server-js/lib/db"
	_ "github.com/zengming00/go-server-js/lib/db/redis"
	_ "github.com/zengming00/go-server-js/lib/image"
	_ "github.com/zengming00/go-server-js/lib/image/png"
	mhttp "github.com/zengming00/go-server-js/lib/net/http"
	_ "github.com/zengming00/go-server-js/lib/net/url"

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
	config           *config
	writeResultValue bool
}

type config struct {
	IndexFile string
	Port      string
	WorkDir   *string
}

func (This *_server) handler(w http.ResponseWriter, r *http.Request) {
	// 加上这行后，内存的使用情况好了一些
	defer runtime.GC()

	u := r.URL

	file := u.Path
	if file == "/" {
		file = This.config.IndexFile
	}
	file = filepath.Join(_cwd, file)

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
			runtime.Set("response", response)

			request := mhttp.NewRequest(runtime, r)
			runtime.Set("request", request)

			cache := NewCache(runtime)
			runtime.Set("cache", cache)

			ret, err := runFile(file, runtime, registry)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				switch err := err.(type) {
				case *goja.Exception:
					fmt.Println("*goja.Exception:", err.String())
				case *goja.InterruptedError:
					fmt.Println("*goja.InterruptedError:", err.String())
				default:
					fmt.Println("default:", err)
				}
				return
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
	w.WriteHeader(http.StatusNotFound)
}

func server() {
	var err error
	s := &_server{
		// registry: new(require.Registry),
		config: &config{
			IndexFile: "/js/index.js",
			Port:      "8080",
		},
	}

	initConfig("config.json", s.config)

	if s.config.WorkDir != nil {
		err := os.Chdir(*s.config.WorkDir)
		if err != nil {
			panic(err)
		}
	}

	_cwd, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// http.Handle("/public", s.fileServer)
	http.HandleFunc("/", s.handler)

	log.Println("server running on " + s.config.Port)
	err = http.ListenAndServe(":"+s.config.Port, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) == 2 {
		filename := os.Args[1]
		start := time.Now()

		registry := new(require.Registry)
		runtime := goja.New()
		r, err := runFile(filename, runtime, registry)
		if err != nil {
			switch err := err.(type) {
			case *goja.Exception:
				fmt.Println("*goja.Exception:", err.String())
			case *goja.InterruptedError:
				fmt.Println("*goja.InterruptedError:", err.String())
			default:
				fmt.Println("default:", err)
			}
			os.Exit(64)
		}
		fmt.Println("result:", *r)
		fmt.Printf("done (%s)", time.Since(start))
		return
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

func initConfig(path string, v interface{}) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			bts, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				panic(err)
			}
			err = ioutil.WriteFile(path, bts, 0666)
			if err != nil {
				panic(err)
			}
			return
		}
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	err = dec.Decode(v)
	if err != nil {
		panic(err)
	}
}
