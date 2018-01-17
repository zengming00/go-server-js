package lib

import (
	"errors"
	"fmt"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/nodejs/require"
)

func init() {
	require.RegisterNativeModule("types", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)

		o.Set("newInt", func(call goja.FunctionCall) goja.Value {
			return runtime.ToValue(new(int))
		})

		o.Set("intValue", func(call goja.FunctionCall) goja.Value {
			v := call.Argument(0).Export()
			if vv, ok := v.(*int); ok {
				return runtime.ToValue(*vv)
			}
			panic(runtime.NewTypeError("p0 is not int type:%T", v))
		})

		o.Set("newBool", func(call goja.FunctionCall) goja.Value {
			return runtime.ToValue(new(bool))
		})

		o.Set("boolValue", func(call goja.FunctionCall) goja.Value {
			v := call.Argument(0).Export()
			if vv, ok := v.(*bool); ok {
				return runtime.ToValue(*vv)
			}
			panic(runtime.NewTypeError("p0 is not bool type:%T", v))
		})

		o.Set("newString", func(call goja.FunctionCall) goja.Value {
			return runtime.ToValue(new(string))
		})

		o.Set("stringValue", func(call goja.FunctionCall) goja.Value {
			v := call.Argument(0).Export()
			if vv, ok := v.(*string); ok {
				return runtime.ToValue(*vv)
			}
			panic(runtime.NewTypeError("p0 is not string type:%T", v))
		})

		o.Set("makeByteSlice", func(call goja.FunctionCall) goja.Value {
			mLen := call.Argument(0).ToInteger()
			mCap := mLen
			if len(call.Arguments) != 1 {
				mCap = call.Argument(1).ToInteger()
			}
			v := make([]byte, mLen, mCap)
			return runtime.ToValue(v)
		})

		o.Set("test", func(call goja.FunctionCall) goja.Value {
			v := call.Argument(0).Export()
			fmt.Printf("%T %[1]v\n", v)
			return nil
		})

		o.Set("err", func(call goja.FunctionCall) goja.Value {
			return MakeErrorValue(runtime, errors.New("terr"))
		})

		o.Set("retNil", func(call goja.FunctionCall) goja.Value {
			// nil 和 goja.Undefined() 效果相同，在js中都是undefined
			return nil
		})

		o.Set("retUndefined", func(call goja.FunctionCall) goja.Value {
			return goja.Undefined()
		})

		o.Set("retNull", func(call goja.FunctionCall) goja.Value {
			// 在js中是null
			return goja.Null()
		})

	})
}
