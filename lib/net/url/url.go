package url

import (
	"net/url"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/nodejs/require"
	"github.com/zengming00/go-server-js/lib"
)

type _url struct {
	runtime *goja.Runtime
	url     *url.URL
}

func (This *_url) port(call goja.FunctionCall) goja.Value {
	str := This.url.Port()
	return This.runtime.ToValue(str)
}

func (This *_url) parse(call goja.FunctionCall) goja.Value {
	rawurl := call.Argument(0).String()
	retVal := This.runtime.NewObject()
	u, err := url.Parse(rawurl)
	if err != nil {
		retVal.Set("err", err.Error())
		return retVal
	}
	retVal.Set("value", NewURL(This.runtime, u))
	return retVal
}

func (This *_url) queryEscape(call goja.FunctionCall) goja.Value {
	str := url.QueryEscape(call.Argument(0).String())
	return This.runtime.ToValue(str)
}

func (This *_url) queryUnescape(call goja.FunctionCall) goja.Value {
	str, err := url.QueryUnescape(call.Argument(0).String())
	return lib.MakeReturnValue(This.runtime, str, err)
}

func (This *_url) newValues(call goja.FunctionCall) goja.Value {
	return NewValues(This.runtime, make(url.Values))
}

func NewURL(runtime *goja.Runtime, u *url.URL) *goja.Object {
	This := &_url{
		runtime: runtime,
		url:     u,
	}
	o := runtime.NewObject()
	o.Set("forceQuery", u.ForceQuery)
	o.Set("fragment", u.Fragment)
	o.Set("host", u.Host)
	o.Set("opaque", u.Opaque)
	o.Set("path", u.Path)
	o.Set("rawPath", u.RawPath)
	o.Set("rawQuery", u.RawQuery)
	o.Set("scheme", u.Scheme)
	// TODO
	// o.Set("", u.User)
	o.Set("port", This.port)
	return o
}

func init() {
	require.RegisterNativeModule("url", func(runtime *goja.Runtime, module *goja.Object) {
		This := &_url{
			runtime: runtime,
		}
		o := module.Get("exports").(*goja.Object)
		o.Set("parse", This.parse)
		o.Set("queryEscape", This.queryEscape)
		o.Set("queryUnescape", This.queryUnescape)
		o.Set("newValues", This.newValues)
	})
}
