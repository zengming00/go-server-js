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

func (This *requestRuntime) parseForm(call goja.FunctionCall) goja.Value {
	err := This.r.ParseForm()
	if err != nil {
		return This.runtime.NewGoError(err)
	}
	return nil
}

func (This *requestRuntime) cookie(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	c, err := This.r.Cookie(name)
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}
	return newCookie(This.runtime, c)
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
	o.Set("host", r.Host)
	o.Set("requestURI", r.RequestURI)

	o.Set("formValue", This.formValue)
	o.Set("userAgent", This.userAgent)
	o.Set("parseForm", This.parseForm)
	o.Set("cookie", This.cookie)

	return o
}
