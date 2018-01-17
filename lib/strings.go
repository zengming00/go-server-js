package lib

import (
	"strings"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/nodejs/require"
)

func init() {
	require.RegisterNativeModule("strings", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("hasPrefix", func(call goja.FunctionCall) goja.Value {
			s := call.Argument(0).String()
			prefix := call.Argument(1).String()
			return runtime.ToValue(strings.HasPrefix(s, prefix))
		})

		o.Set("hasSuffix", func(call goja.FunctionCall) goja.Value {
			s := call.Argument(0).String()
			suffix := call.Argument(1).String()
			return runtime.ToValue(strings.HasSuffix(s, suffix))
		})

	})
}
