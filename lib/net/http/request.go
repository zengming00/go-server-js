package http

import (
	"io/ioutil"
	"net/http"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib"
	"github.com/zengming00/go-server-js/lib/net/url"
)

func NewRequest(runtime *goja.Runtime, r *http.Request) *goja.Object {
	o := runtime.NewObject()

	o.Set("isMissingFile", func(call goja.FunctionCall) goja.Value {
		// todo get pro
		p0 := call.Argument(0).Export()
		if err, ok := p0.(error); ok {
			return runtime.ToValue(err == http.ErrMissingFile)
		}
		panic(runtime.NewTypeError("p0 is not error type:%T", p0))
	})

	o.Set("getContentLength", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(r.ContentLength)
	})

	o.Set("getMethod", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(r.Method)
	})

	o.Set("getHost", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(r.Host)
	})

	o.Set("getBody", func(call goja.FunctionCall) goja.Value {
		//todo
		return runtime.ToValue(r.Body)
	})

	o.Set("getHeader", func(call goja.FunctionCall) goja.Value {
		return NewHeader(runtime, r.Header)
	})

	o.Set("getHeaders", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(map[string][]string(r.Header))
	})

	o.Set("getUri", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(r.RequestURI)
	})

	o.Set("getUrl", func(call goja.FunctionCall) goja.Value {
		return url.NewURL(runtime, r.URL)
	})

	o.Set("getRemoteAddr", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(r.RemoteAddr)
	})

	o.Set("getForm", func(call goja.FunctionCall) goja.Value {
		return url.NewValues(runtime, r.Form)
	})

	o.Set("formValue", func(call goja.FunctionCall) goja.Value {
		key := call.Argument(0).String()
		value := r.FormValue(key)
		return runtime.ToValue(value)
	})

	o.Set("formFile", func(call goja.FunctionCall) goja.Value {
		key := call.Argument(0).String()
		file, fileHeader, err := r.FormFile(key)
		if err != nil {
			return lib.MakeErrorValue(runtime, err)
		}
		return runtime.ToValue(map[string]interface{}{
			"file":   NewMultipartFile(runtime, file),
			"name":   fileHeader.Filename,
			"header": map[string][]string(fileHeader.Header),
		})
	})

	o.Set("userAgent", func(call goja.FunctionCall) goja.Value {
		value := r.UserAgent()
		return runtime.ToValue(value)
	})

	o.Set("parseForm", func(call goja.FunctionCall) goja.Value {
		err := r.ParseForm()
		if err != nil {
			return lib.MakeErrorValue(runtime, err)
		}
		return nil
	})

	o.Set("parseMultipartForm", func(call goja.FunctionCall) goja.Value {
		maxMemory := call.Argument(0).ToInteger()
		err := r.ParseMultipartForm(maxMemory)
		if err != nil {
			return lib.MakeErrorValue(runtime, err)
		}
		return nil
	})

	o.Set("cookie", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		c, err := r.Cookie(name)
		if err != nil {
			return lib.MakeErrorValue(runtime, err)
		}
		return lib.MakeReturnValue(runtime, NewCookie(runtime, c))
	})

	o.Set("cookies", func(call goja.FunctionCall) goja.Value {
		cks := r.Cookies()
		arr := make([]*goja.Object, len(cks))
		for i, v := range cks {
			arr[i] = NewCookie(runtime, v)
		}
		return runtime.ToValue(arr)
	})

	o.Set("getRawBody", func(call goja.FunctionCall) goja.Value {
		bts, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return lib.MakeErrorValue(runtime, err)
		}
		return lib.MakeReturnValue(runtime, bts)
	})

	return o
}
