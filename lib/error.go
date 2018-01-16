package lib

import (
	"errors"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/nodejs/require"
)

func NewError(runtime *goja.Runtime, err error) *goja.Object {
	o := runtime.NewObject()
	o.Set("error", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(err.Error())
	})
	o.Set("getPrototype", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(err)
	})
	return o
}

func init() {
	require.RegisterNativeModule("error", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("new", func(call goja.FunctionCall) goja.Value {
			text := call.Argument(0).String()
			err := errors.New(text)
			return NewError(runtime, err)
		})
	})
}
