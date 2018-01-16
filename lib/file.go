package lib

import (
	"io/ioutil"
	"os"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/nodejs/require"
)

func NewFile(runtime *goja.Runtime, file *os.File) *goja.Object {
	o := runtime.NewObject()
	o.Set("writeString", func(call goja.FunctionCall) goja.Value {
		s := call.Argument(0).String()
		n, err := file.WriteString(s)
		if err != nil {
			return MakeErrorValue(runtime, err)
		}
		return MakeReturnValue(runtime, n)
	})

	o.Set("close", func(call goja.FunctionCall) goja.Value {
		err := file.Close()
		if err != nil {
			return MakeErrorValue(runtime, err)
		}
		return nil
	})

	o.Set("nativeType", file)

	return o
}

func init() {
	require.RegisterNativeModule("file", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("write", func(call goja.FunctionCall) goja.Value {
			filename := call.Argument(0).String()
			data := call.Argument(1).Export()
			if bts, ok := data.([]byte); ok {
				err := ioutil.WriteFile(filename, bts, 0666)
				if err != nil {
					return MakeErrorValue(runtime, err)
				}
				return nil
			}
			panic(runtime.NewTypeError("p1 is not []byte type:%T", data))
		})

		o.Set("read", func(call goja.FunctionCall) goja.Value {
			filename := call.Argument(0).String()
			data, err := ioutil.ReadFile(filename)
			if err != nil {
				return MakeErrorValue(runtime, err)
			}
			return MakeReturnValue(runtime, data)
		})
	})
}
