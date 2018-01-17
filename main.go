package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
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
	config           *config
	writeResultValue bool
}

type config struct {
	IndexFile             string
	Port                  string
	WorkDir               *string
	SessionMaxLifeTimeSec int64
	CacheGcIntervalSec    int64
	ScriptTimeoutSec      int
}

func (This *_server) handler(w http.ResponseWriter, r *http.Request) {
	u := r.URL
	file := u.Path
	if file == "/" {
		file = This.config.IndexFile
	}
	file = filepath.Join(_cwd, file)

	if strings.HasPrefix(file, _cwd) {
		ext := filepath.Ext(file)
		if ext == ".js" {
			runtime := goja.New()
			registry := new(require.Registry)
			runtime.Set("response", mhttp.NewResponse(runtime, w))
			runtime.Set("request", mhttp.NewRequest(runtime, r))
			runtime.Set("cache", NewCache(runtime, _cacheMgr))
			runtime.Set("session", NewSession(runtime, _sessionMgr, w, r))

			normalEndCh := make(chan struct{})

			go func() {
				ret, err := runFile(file, runtime, registry)
				if err != nil {
					switch err := err.(type) {
					case *goja.Exception:
						w.WriteHeader(http.StatusInternalServerError)
						fmt.Println("*goja.Exception:", err.String())
						if v := err.Value().ToObject(runtime).Get("nativeType"); v != nil {
							fmt.Printf("%T:%[1]v\n", v.Export())
						}
					case *goja.InterruptedError:
						fmt.Println("*goja.InterruptedError:", err.String())
					default:
						fmt.Println("default err:", err)
					}
					return
				}

				// if goja.IsNull(*ret) || goja.IsUndefined(*ret) {
				// 	w.WriteHeader(http.StatusOK)
				// 	return
				// }
				if This.writeResultValue {
					// w.Header().Set("Content-Type", "text/html; charset=utf-8")
					w.Write([]byte((*ret).String()))
				}
				normalEndCh <- struct{}{}
			}()

			if This.config.ScriptTimeoutSec > 0 {
				timeoutCh := time.After(time.Duration(This.config.ScriptTimeoutSec) * time.Second)
				select {
				case <-timeoutCh:
					runtime.Interrupt("run code timeout, halt")
					w.WriteHeader(http.StatusInternalServerError)
				case <-normalEndCh:
				}
			} else {
				<-normalEndCh
			}
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func server() {
	var err error
	s := &_server{
		config: &config{
			IndexFile: "/index.js",
			Port:      "8080",
			SessionMaxLifeTimeSec: 60 * 15,
			CacheGcIntervalSec:    60,
			ScriptTimeoutSec:      -1,
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
	log.Printf("ScriptTimeoutSec: %d\r\n", s.config.ScriptTimeoutSec)

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
	// debug.SetGCPercent(1)
	rand.Seed(time.Now().UnixNano())

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

	// (function(){ ... })() 用来允许直接在文件顶层return
	prg, err := goja.Compile(filename, "(function(){"+string(src)+"\n})()", false)
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
