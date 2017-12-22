package main

import "github.com/dop251/goja"
import "net/http"
import "fmt"

type responseRuntime struct {
	runtime *goja.Runtime
	w       *http.ResponseWriter
}

type headerObj struct {
	runtime *goja.Runtime
	header  *http.Header
}

func (This *headerObj) add(call goja.FunctionCall) goja.Value {
	key := call.Argument(0).String()
	value := call.Argument(1).String()
	This.header.Add(key, value)
	return nil
}

func (This *headerObj) del(call goja.FunctionCall) goja.Value {
	key := call.Argument(0).String()
	This.header.Del(key)
	return nil
}

func (This *headerObj) get(call goja.FunctionCall) goja.Value {
	key := call.Argument(0).String()
	value := This.header.Get(key)
	return This.runtime.ToValue(value)
}

func (This *headerObj) set(call goja.FunctionCall) goja.Value {
	key := call.Argument(0).String()
	value := call.Argument(1).String()
	This.header.Set(key, value)
	return nil
}

func (This *responseRuntime) header(call goja.FunctionCall) goja.Value {
	w := *This.w
	hd := w.Header()
	obj := &headerObj{
		runtime: This.runtime,
		header:  &hd,
	}
	o := This.runtime.NewObject()
	o.Set("add", obj.add)
	o.Set("del", obj.del)
	o.Set("set", obj.set)
	o.Set("get", obj.get)
	return o
}

func (This *responseRuntime) write(call goja.FunctionCall) goja.Value {
	p1 := call.Argument(0).Export()

	var data []byte
	switch t := p1.(type) {
	case []interface{}:
		data = make([]byte, len(t))
		for i, v := range t {
			if val, ok := v.(int64); ok {
				if val >= 0 && val <= 255 {
					data[i] = byte(val)
					continue
				}
			}
			panic(This.runtime.NewTypeError(fmt.Sprintf("response.write() can not convert to byte `data[%d] = %v`", i, v)))
		}
	case []byte:
		data = t
	default:
		panic(This.runtime.NewTypeError("response.write() data is not a byte array"))
	}

	w := *This.w
	n, err := w.Write(data)
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}
	return This.runtime.ToValue(n)
}

func (This *responseRuntime) writeHeader(call goja.FunctionCall) goja.Value {
	n := call.Argument(0).ToInteger()
	w := *This.w
	w.WriteHeader(int(n))
	return nil
}

func getResponse(runtime *goja.Runtime, w *http.ResponseWriter) *goja.Object {
	This := &responseRuntime{
		runtime: runtime,
		w:       w,
	}

	o := runtime.NewObject()
	o.Set("header", This.header)
	o.Set("write", This.write)
	o.Set("writeHeader", This.writeHeader)
	return o
}
