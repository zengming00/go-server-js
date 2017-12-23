package lib

import (
	"net/url"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type _url struct {
	runtime *goja.Runtime
}

func (This *_url) queryEscape(call goja.FunctionCall) goja.Value {
	str := url.QueryEscape(call.Argument(0).String())
	return This.runtime.ToValue(str)
}

func init() {
	require.RegisterNativeModule("url", func(runtime *goja.Runtime, module *goja.Object) {
		This := &_url{
			runtime: runtime,
		}
		o := module.Get("exports").(*goja.Object)
		o.Set("queryEscape", This.queryEscape)
	})
}
