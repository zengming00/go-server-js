package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/zengming00/go-server-js/lib"
	_ "github.com/zengming00/go-server-js/lib/db"
	_ "github.com/zengming00/go-server-js/lib/db/redis"
	mhttp "github.com/zengming00/go-server-js/lib/http"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

// go-sqlite3在win32下的问题
// https://github.com/mattn/go-sqlite3/issues/358

var _cwd string

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

func handler(w http.ResponseWriter, r *http.Request) {
	u := r.URL
	file := filepath.Join(_cwd, u.Path)
	if strings.HasPrefix(file, _cwd) {
		ext := filepath.Ext(file)
		if ext == ".js" {
			runtime := goja.New()
			response := mhttp.NewResponse(runtime, &w)
			request := mhttp.NewRequest(runtime, r)
			runtime.Set("response", response)
			runtime.Set("request", request)

			ret, err := runFile(file, runtime)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			if goja.IsNull(*ret) || goja.IsUndefined(*ret) {
				w.WriteHeader(http.StatusOK)
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write([]byte((*ret).String()))
			return
		}
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func server() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	handErr(err)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("please input a file name")
		os.Exit(1)
	}
	filename := os.Args[1]

	r, err := runFile(filename, nil)
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
	server()
}

func runFile(filename string, runtime *goja.Runtime) (*goja.Value, error) {
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if runtime == nil {
		runtime = goja.New()
	}

	// this can be shared by multiple runtimes
	registry := new(require.Registry)
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
