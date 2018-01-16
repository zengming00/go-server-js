package url

import (
	"net/url"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib"
	"github.com/zengming00/go-server-js/nodejs/require"
)

type _url struct {
	runtime *goja.Runtime
}

func (This *_url) parse(call goja.FunctionCall) goja.Value {
	rawurl := call.Argument(0).String()
	u, err := url.Parse(rawurl)
	if err != nil {
		return lib.MakeErrorValue(This.runtime, err)
	}
	return lib.MakeReturnValue(This.runtime, NewURL(This.runtime, u))
}

func (This *_url) queryEscape(call goja.FunctionCall) goja.Value {
	str := url.QueryEscape(call.Argument(0).String())
	return This.runtime.ToValue(str)
}

func (This *_url) queryUnescape(call goja.FunctionCall) goja.Value {
	str, err := url.QueryUnescape(call.Argument(0).String())
	if err != nil {
		return lib.MakeErrorValue(This.runtime, err)
	}
	return lib.MakeReturnValue(This.runtime, str)
}

func (This *_url) parseRequestURI(call goja.FunctionCall) goja.Value {
	rawurl := call.Argument(0).String()
	mUrl, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return lib.MakeErrorValue(This.runtime, err)
	}
	return lib.MakeReturnValue(This.runtime, NewURL(This.runtime, mUrl))
}

func (This *_url) parseQuery(call goja.FunctionCall) goja.Value {
	query := call.Argument(0).String()
	values, err := url.ParseQuery(query)
	if err != nil {
		return lib.MakeErrorValue(This.runtime, err)
	}
	return lib.MakeReturnValue(This.runtime, NewValues(This.runtime, values))
}

func (This *_url) newValues(call goja.FunctionCall) goja.Value {
	return NewValues(This.runtime, make(url.Values))
}

func NewURL(runtime *goja.Runtime, u *url.URL) *goja.Object {
	// TODO
	o := runtime.NewObject()
	o.Set("getForceQuery", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(u.ForceQuery)
	})

	o.Set("getFragment", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(u.Fragment)
	})

	o.Set("getHost", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(u.Host)
	})

	o.Set("getOpaque", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(u.Opaque)
	})

	o.Set("getPath", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(u.Path)
	})

	o.Set("getRawPath", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(u.RawPath)
	})

	o.Set("getRawQuery", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(u.RawQuery)
	})

	o.Set("getScheme", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(u.Scheme)
	})

	o.Set("getPort", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(u.Port())
	})

	return o
}

func init() {
	require.RegisterNativeModule("url", func(runtime *goja.Runtime, module *goja.Object) {
		This := &_url{
			runtime: runtime,
		}
		o := module.Get("exports").(*goja.Object)
		o.Set("parse", This.parse)
		o.Set("parseQuery", This.parseQuery)
		o.Set("parseRequestURI", This.parseRequestURI)
		o.Set("queryEscape", This.queryEscape)
		o.Set("queryUnescape", This.queryUnescape)
		o.Set("newValues", This.newValues)
	})
}
