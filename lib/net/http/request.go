package http

import (
	"net/http"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib/net/url"
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
	if err != nil {
		return This.runtime.ToValue(err.Error())
	}
	return nil
}

func (This *_req) getHeader(call goja.FunctionCall) goja.Value {
	key := call.Argument(0).String()
	value := This.r.Header.Get(key)
	return This.runtime.ToValue(value)
}

func (This *_req) cookie(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	retVal := This.runtime.NewObject()

	c, err := This.r.Cookie(name)
	if err != nil {
		retVal.Set("err", err.Error())
	} else {
		retVal.Set("value", NewCookie(This.runtime, c))
	}
	return retVal
}

func (This *_req) cookies(call goja.FunctionCall) goja.Value {
	cks := This.r.Cookies()
	arr := make([]*goja.Object, len(cks))
	for i, v := range cks {
		arr[i] = NewCookie(This.runtime, v)
	}
	return This.runtime.ToValue(arr)
}

func NewRequest(runtime *goja.Runtime, r *http.Request) *goja.Object {
	This := &_req{
		runtime: runtime,
		r:       r,
	}

	o := runtime.NewObject()
	o.Set("contentLength", r.ContentLength)
	o.Set("method", r.Method)
	o.Set("host", r.Host)
	o.Set("header", NewHeader(runtime, &r.Header))
	o.Set("uri", r.RequestURI)
	o.Set("remoteAddr", r.RemoteAddr)
	o.Set("form", url.NewValues(runtime, &r.Form))

	o.Set("getHeader", This.getHeader)
	o.Set("formValue", This.formValue)
	o.Set("userAgent", This.userAgent)
	o.Set("parseForm", This.parseForm)
	o.Set("cookie", This.cookie)
	o.Set("cookies", This.cookies)

	return o
}
