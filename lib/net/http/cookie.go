package http

import (
	"net/http"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib"
)

func NewCookie(runtime *goja.Runtime, cookie *http.Cookie) *goja.Object {
	o := runtime.NewObject()
	o.Set("string", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(cookie.String())
	})
	o.Set("getDomain", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(cookie.Domain)
	})
	o.Set("getExpires", func(call goja.FunctionCall) goja.Value {
		return lib.NewTime(runtime, &cookie.Expires)
	})
	o.Set("getHttpOnly", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(cookie.HttpOnly)
	})
	o.Set("getMaxAge", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(cookie.MaxAge)
	})
	o.Set("getName", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(cookie.Name)
	})
	o.Set("getPath", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(cookie.Path)
	})
	o.Set("getRaw", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(cookie.Raw)
	})
	o.Set("getRawExpires", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(cookie.RawExpires)
	})
	o.Set("getSecure", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(cookie.Secure)
	})
	o.Set("getUnparsed", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(cookie.Unparsed)
	})
	o.Set("getValue", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(cookie.Value)
	})
	return o
}
