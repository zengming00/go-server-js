package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"

	_ "github.com/zengming00/go-server-js/lib"
	_ "github.com/zengming00/go-server-js/lib/db"
	_ "github.com/zengming00/go-server-js/lib/db/redis"
	_ "github.com/zengming00/go-server-js/lib/image"
	_ "github.com/zengming00/go-server-js/lib/image/png"
	mhttp "github.com/zengming00/go-server-js/lib/net/http"
	_ "github.com/zengming00/go-server-js/lib/net/url"
	_ "github.com/zengming00/go-server-js/lib/path"

	_ "net/http/pprof"

	"github.com/dop251/goja"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zengming00/go-server-js/nodejs/console"
	"github.com/zengming00/go-server-js/nodejs/require"
)

var configFile = flag.String("config", "config.json", "set the config.json")

var _cwd string
var _sessionMgr *SessionMgr
var _cacheMgr *CacheMgr

type _server struct {
	runtime          *goja.Runtime
	registry         *require.Registry
	config           *config
	writeResultValue bool
}

type config struct {
	IndexFile             string
	Port                  string
	WorkDir               *string
	SessionMaxLifeTimeSec int64
	CacheGcIntervalSec    int64
}

func (This *_server) handler(w http.ResponseWriter, r *http.Request) {
	// 加上这行后，内存的使用情况好了一些
	// runtime.GC()
	// defer runtime.GC()

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
			runtime.Set("response", mhttp.NewResponse(runtime, w))
			runtime.Set("request", mhttp.NewRequest(runtime, r))
			runtime.Set("cache", NewCache(runtime, _cacheMgr))
			runtime.Set("session", NewSession(runtime, _sessionMgr, w, r))

			ret, err := runFile(file, runtime, registry)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				switch err := err.(type) {
				case *goja.Exception:
					fmt.Println("*goja.Exception:", err.String())
					if v := err.Value().ToObject(runtime).Get("nativeType"); v != nil {
						fmt.Printf("%T:%[1]v\n", v.Export())
					}
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
			SessionMaxLifeTimeSec: 60 * 15,
			CacheGcIntervalSec:    60,
		},
	}

	initConfig(*configFile, s.config)

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

	_sessionMgr = NewSessionMgr("sid", s.config.SessionMaxLifeTimeSec)
	log.Printf("SessionMaxLifeTimeSec: %d\r\n", s.config.SessionMaxLifeTimeSec)
	_cacheMgr = NewCacheMgr(s.config.CacheGcIntervalSec)
	log.Printf("CacheGcIntervalSec: %d\r\n", s.config.CacheGcIntervalSec)

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
	flag.Parse()
	debug.SetGCPercent(1)
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
