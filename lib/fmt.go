package lib

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/nodejs/require"
)

func init() {
	require.RegisterNativeModule("fmt", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("sprintf", func(call goja.FunctionCall) goja.Value {
			format := call.Argument(0).String()
			args := GetAllArgs(&call)
			str := fmt.Sprintf(format, args[1:]...)
			return runtime.ToValue(str)
		})

		o.Set("printf", func(call goja.FunctionCall) goja.Value {
			format := call.Argument(0).String()
			args := GetAllArgs(&call)
			n, err := fmt.Printf(format, args[1:]...)
			if err != nil {
				return MakeErrorValue(runtime, err)
			}
			return MakeReturnValue(runtime, n)
		})

		o.Set("println", func(call goja.FunctionCall) goja.Value {
			args := GetAllArgs(&call)
			n, err := fmt.Println(args...)
			if err != nil {
				return MakeErrorValue(runtime, err)
			}
			return MakeReturnValue(runtime, n)
		})

		o.Set("print", func(call goja.FunctionCall) goja.Value {
			args := GetAllArgs(&call)
			n, err := fmt.Print(args...)
			if err != nil {
				return MakeErrorValue(runtime, err)
			}
			return MakeReturnValue(runtime, n)
		})
	})
}
