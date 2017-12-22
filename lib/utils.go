package lib

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type utilsRuntime struct {
	runtime *goja.Runtime
}

func (This *utilsRuntime) print(call goja.FunctionCall) goja.Value {
	fmt.Print(call.Argument(0).String())
	return nil
}

func (This *utilsRuntime) toString(call goja.FunctionCall) goja.Value {
	data := call.Argument(0).Export()
	if bts, ok := data.([]byte); ok {
		return This.runtime.ToValue(string(bts))
	}
	panic(This.runtime.NewTypeError("toString data is not a byte array"))
}

func init() {
	require.RegisterNativeModule("utils", func(runtime *goja.Runtime, module *goja.Object) {
		This := &utilsRuntime{
			runtime: runtime,
		}

		o := module.Get("exports").(*goja.Object)
		o.Set("print", This.print)
		o.Set("toString", This.toString)
	})
}
