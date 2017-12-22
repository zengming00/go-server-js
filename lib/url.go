package lib

import (
	"net/url"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type urlRuntime struct {
	runtime *goja.Runtime
}

func (This *urlRuntime) queryEscape(call goja.FunctionCall) goja.Value {
	str := url.QueryEscape(call.Argument(0).String())
	return This.runtime.ToValue(str)
}

func init() {
	require.RegisterNativeModule("url", func(runtime *goja.Runtime, module *goja.Object) {
		This := &urlRuntime{
			runtime: runtime,
		}

		o := module.Get("exports").(*goja.Object)
		o.Set("queryEscape", This.queryEscape)
	})
}
