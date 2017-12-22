package lib

import (
	"io/ioutil"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type fileRuntime struct {
	runtime *goja.Runtime
}

func (This *fileRuntime) write(call goja.FunctionCall) goja.Value {
	filename := call.Argument(0).String()
	data := call.Argument(1).Export()
	if bts, ok := data.([]byte); ok {
		err := ioutil.WriteFile(filename, bts, 0666)
		if err != nil {
			panic(This.runtime.NewGoError(err))
		}
		return nil
	}
	panic(This.runtime.NewTypeError("write data is not a byte array"))
}

func (This *fileRuntime) read(call goja.FunctionCall) goja.Value {
	filename := call.Argument(0).String()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}
	return This.runtime.ToValue(data)
}

func init() {
	require.RegisterNativeModule("file", func(runtime *goja.Runtime, module *goja.Object) {
		This := &fileRuntime{
			runtime: runtime,
		}

		o := module.Get("exports").(*goja.Object)
		o.Set("write", This.write)
		o.Set("read", This.read)
	})
}
