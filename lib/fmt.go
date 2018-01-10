package lib

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/nodejs/require"
)

type _fmt struct {
	runtime *goja.Runtime
}

func (This *_fmt) sprintf(call goja.FunctionCall) goja.Value {
	format := call.Argument(0).String()
	args := GetAllArgs(&call)
	str := fmt.Sprintf(format, args[1:]...)
	return This.runtime.ToValue(str)
}

func (This *_fmt) printf(call goja.FunctionCall) goja.Value {
	format := call.Argument(0).String()
	args := GetAllArgs(&call)
	n, err := fmt.Printf(format, args[1:]...)
	return MakeReturnValue(This.runtime, n, err)
}

func (This *_fmt) println(call goja.FunctionCall) goja.Value {
	args := GetAllArgs(&call)
	n, err := fmt.Println(args...)
	return MakeReturnValue(This.runtime, n, err)
}

func init() {
	require.RegisterNativeModule("fmt", func(runtime *goja.Runtime, module *goja.Object) {
		This := &_fmt{
			runtime: runtime,
		}
		o := module.Get("exports").(*goja.Object)
		o.Set("sprintf", This.sprintf)
		o.Set("printf", This.printf)
		o.Set("println", This.println)
		o.Set("print", func(call goja.FunctionCall) goja.Value {
			args := GetAllArgs(&call)
			n, err := fmt.Print(args...)
			return MakeReturnValue(runtime, n, err)
		})
	})
}
