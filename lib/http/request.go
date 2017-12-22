package http

import (
	"net/http"

	"github.com/dop251/goja"
)

type requestRuntime struct {
	runtime *goja.Runtime
	r       *http.Request
}

func (This *requestRuntime) formValue(call goja.FunctionCall) goja.Value {
	key := call.Argument(0).String()
	value := This.r.FormValue(key)
	return This.runtime.ToValue(value)
}

func (This *requestRuntime) userAgent(call goja.FunctionCall) goja.Value {
	value := This.r.UserAgent()
	return This.runtime.ToValue(value)
}

func NewRequest(runtime *goja.Runtime, r *http.Request) *goja.Object {
	This := &requestRuntime{
		runtime: runtime,
		r:       r,
	}

	o := runtime.NewObject()
	o.Set("contentLength", r.ContentLength)
	o.Set("method", r.Method)
	o.Set("header", newHeader(runtime, &r.Header))

	o.Set("formValue", This.formValue)
	o.Set("userAgent", This.userAgent)

	return o
}
