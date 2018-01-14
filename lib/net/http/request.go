package http

import (
	"io/ioutil"
	"net/http"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib"
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

func (This *_req) formFile(call goja.FunctionCall) goja.Value {
	key := call.Argument(0).String()
	file, fileHeader, err := This.r.FormFile(key)
	if err != nil {
		return This.runtime.ToValue(map[string]interface{}{
			"err": err.Error(),
		})
	}
	return This.runtime.ToValue(map[string]interface{}{
		"file":   NewMultipartFile(This.runtime, file),
		"name":   fileHeader.Filename,
		"header": map[string][]string(fileHeader.Header),
	})
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

func (This *_req) parseMultipartForm(call goja.FunctionCall) goja.Value {
	maxMemory := call.Argument(0).ToInteger()
	err := This.r.ParseMultipartForm(maxMemory)
	if err != nil {
		return This.runtime.ToValue(err.Error())
	}
	return nil
}

func (This *_req) cookie(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	c, err := This.r.Cookie(name)
	if err != nil {
		return This.runtime.ToValue(map[string]interface{}{
			"err": err.Error(),
		})
	}
	return This.runtime.ToValue(map[string]interface{}{
		"value": NewCookie(This.runtime, c),
	})
}

func (This *_req) cookies(call goja.FunctionCall) goja.Value {
	cks := This.r.Cookies()
	arr := make([]*goja.Object, len(cks))
	for i, v := range cks {
		arr[i] = NewCookie(This.runtime, v)
	}
	return This.runtime.ToValue(arr)
}

func (This *_req) getRawBody(call goja.FunctionCall) goja.Value {
	bts, err := ioutil.ReadAll(This.r.Body)
	return lib.MakeReturnValue(This.runtime, bts, err)
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
	o.Set("body", r.Body)
	o.Set("header", NewHeader(runtime, r.Header))
	o.Set("headers", map[string][]string(r.Header))
	o.Set("uri", r.RequestURI)
	o.Set("url", url.NewURL(runtime, r.URL))	
	o.Set("remoteAddr", r.RemoteAddr)
	o.Set("form", url.NewValues(runtime, r.Form))

	o.Set("formValue", This.formValue)
	o.Set("formFile", This.formFile)
	o.Set("userAgent", This.userAgent)
	o.Set("parseForm", This.parseForm)
	o.Set("parseMultipartForm", This.parseMultipartForm)
	o.Set("cookie", This.cookie)
	o.Set("cookies", This.cookies)
	o.Set("getRawBody", This.getRawBody)
	return o
}
