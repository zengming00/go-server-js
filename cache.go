package main

import (
	"sync"

	"github.com/dop251/goja"
)

type _cache struct {
	runtime *goja.Runtime
}

var cache = make(map[string]interface{})
var mu sync.RWMutex

func (This *_cache) set(call goja.FunctionCall) goja.Value {
	mu.Lock()
	defer mu.Unlock()
	s := call.Argument(0).String()
	v := call.Argument(1).Export()
	cache[s] = v
	return nil
}

func (This *_cache) get(call goja.FunctionCall) goja.Value {
	mu.RLock()
	defer mu.RUnlock()
	s := call.Argument(0).String()
	return This.runtime.ToValue(cache[s])
}

func NewCache(runtime *goja.Runtime) *goja.Object {
	This := &_cache{
		runtime: runtime,
	}

	o := runtime.NewObject()
	o.Set("set", This.set)
	o.Set("get", This.get)
	return o
}
