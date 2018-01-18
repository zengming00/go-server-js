package http

import (
	"bytes"
	"net/http"

	"github.com/dop251/goja"
)

func NewHeader(runtime *goja.Runtime, header http.Header) *goja.Object {
	o := runtime.NewObject()
	o.Set("add", func(call goja.FunctionCall) goja.Value {
		key := call.Argument(0).String()
		value := call.Argument(1).String()
		header.Add(key, value)
		return nil
	})

	o.Set("del", func(call goja.FunctionCall) goja.Value {
		key := call.Argument(0).String()
		header.Del(key)
		return nil
	})

	o.Set("get", func(call goja.FunctionCall) goja.Value {
		key := call.Argument(0).String()
		value := header.Get(key)
		return runtime.ToValue(value)
	})

	o.Set("gets", func(call goja.FunctionCall) goja.Value {
		key := call.Argument(0).String()
		value := header[key]
		return runtime.ToValue(value)
	})

	o.Set("set", func(call goja.FunctionCall) goja.Value {
		key := call.Argument(0).String()
		value := call.Argument(1).String()
		header.Set(key, value)
		return nil
	})

	o.Set("getRaw", func(call goja.FunctionCall) goja.Value {
		byteBuf := &bytes.Buffer{}
		header.Write(byteBuf)
		return runtime.ToValue(byteBuf.Bytes())
	})
	return o
}
