package http

import (
	"bytes"
	"net/http"

	"github.com/dop251/goja"
)

type _header struct {
	runtime *goja.Runtime
	header  http.Header
}

func (This *_header) add(call goja.FunctionCall) goja.Value {
	key := call.Argument(0).String()
	value := call.Argument(1).String()
	This.header.Add(key, value)
	return nil
}

func (This *_header) del(call goja.FunctionCall) goja.Value {
	key := call.Argument(0).String()
	This.header.Del(key)
	return nil
}

func (This *_header) get(call goja.FunctionCall) goja.Value {
	key := call.Argument(0).String()
	value := This.header.Get(key)
	return This.runtime.ToValue(value)
}

func (This *_header) gets(call goja.FunctionCall) goja.Value {
	key := call.Argument(0).String()
	value := This.header[key]
	return This.runtime.ToValue(value)
}

func (This *_header) set(call goja.FunctionCall) goja.Value {
	key := call.Argument(0).String()
	value := call.Argument(1).String()
	This.header.Set(key, value)
	return nil
}

func (This *_header) getRaw(call goja.FunctionCall) goja.Value {
	byteBuf := &bytes.Buffer{}
	This.header.Write(byteBuf)
	return This.runtime.ToValue(byteBuf.Bytes())
}

func NewHeader(runtime *goja.Runtime, header http.Header) *goja.Object {
	This := &_header{
		runtime: runtime,
		header:  header,
	}
	o := runtime.NewObject()
	o.Set("add", This.add)
	o.Set("del", This.del)
	o.Set("set", This.set)
	o.Set("get", This.get)
	o.Set("gets", This.gets)
	o.Set("getRaw", This.getRaw)
	return o
}
