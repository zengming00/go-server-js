package path

import (
	"path/filepath"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib"
	"github.com/zengming00/go-server-js/nodejs/require"
)

func init() {
	require.RegisterNativeModule("path/filepath", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("abs", func(call goja.FunctionCall) goja.Value {
			path := call.Argument(0).String()
			str, err := filepath.Abs(path)
			return lib.MakeReturnValue(runtime, str, err)
		})

		o.Set("join", func(call goja.FunctionCall) goja.Value {
			arr := lib.GetAllArgs_string(runtime, &call)
			str := filepath.Join(arr...)
			return runtime.ToValue(str)
		})

	})
}
