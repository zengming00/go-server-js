package main

import (
	"reflect"
	"sync"

	"github.com/dop251/goja"
)

type _cache struct {
	runtime *goja.Runtime
}

var cache = make(map[string]interface{})
var mu sync.RWMutex

var validType = []reflect.Kind{
	reflect.Bool,
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Float32,
	reflect.Float64,
	reflect.Complex64,
	reflect.Complex128,
	reflect.String,
}

func isValidType(val interface{}) bool {
	k := reflect.TypeOf(val).Kind()
	for _, v := range validType {
		if v == k {
			return true
		}
	}
	return false
}

func (This *_cache) set(call goja.FunctionCall) goja.Value {
	s := call.Argument(0).String()
	v := call.Argument(1).Export()
	if isValidType(v) {
		mu.Lock()
		defer mu.Unlock()
		cache[s] = v
	}
	panic(This.runtime.NewTypeError("set value type is not permitted"))
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
