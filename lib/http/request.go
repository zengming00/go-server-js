package http

import (
	"net/http"

	"github.com/dop251/goja"
)

type _req struct {
	runtime *goja.Runtime
	r       *http.Request
}

func (This *_req) formValue(call goja.FunctionCall) goja.Value {
	key := call.Argument(0).String()
	value := This.r.FormValue(key)
	return This.runtime.ToValue(value)
}

func (This *_req) userAgent(call goja.FunctionCall) goja.Value {
	value := This.r.UserAgent()
	return This.runtime.ToValue(value)
}

func (This *_req) parseForm(call goja.FunctionCall) goja.Value {
	err := This.r.ParseForm()
	return This.runtime.ToValue(err)
}

func (This *_req) cookie(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	retVal := This.runtime.NewObject()

	c, err := This.r.Cookie(name)
	if err != nil {
		retVal.Set("err", err.Error())
	} else {
		retVal.Set("cookie", NewCookie(This.runtime, c))
	}
	return retVal
}

func NewRequest(runtime *goja.Runtime, r *http.Request) *goja.Object {
	This := &_req{
		runtime: runtime,
		r:       r,
	}

	o := runtime.NewObject()
	o.Set("contentLength", r.ContentLength)
	o.Set("method", r.Method)
	o.Set("header", NewHeader(runtime, &r.Header))
	o.Set("host", r.Host)
	o.Set("requestURI", r.RequestURI)

	o.Set("formValue", This.formValue)
	o.Set("userAgent", This.userAgent)
	o.Set("parseForm", This.parseForm)
	o.Set("cookie", This.cookie)

	return o
}
