package lib

import (
	"os"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type _os struct {
	runtime *goja.Runtime
}

func (This *_os) create(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	retVal := This.runtime.NewObject()

	file, err := os.Create(name)
	if err != nil {
		retVal.Set("err", err.Error())
		return retVal
	}
	retVal.Set("value", NewFile(This.runtime, file))
	return retVal
}

func init() {
	require.RegisterNativeModule("os", func(runtime *goja.Runtime, module *goja.Object) {
		This := &_os{
			runtime: runtime,
		}
		o := module.Get("exports").(*goja.Object)
		o.Set("create", This.create)
	})
}
