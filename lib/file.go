package lib

import (
	"io/ioutil"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type _file struct {
	runtime *goja.Runtime
}

func (This *_file) write(call goja.FunctionCall) goja.Value {
	filename := call.Argument(0).String()
	data := call.Argument(1).Export()
	if bts, ok := data.([]byte); ok {
		err := ioutil.WriteFile(filename, bts, 0666)
		if err != nil {
			return This.runtime.NewGoError(err)
		}
		return nil
	}
	return This.runtime.NewTypeError("file.write() data is not a byte array")
}

func (This *_file) read(call goja.FunctionCall) goja.Value {
	filename := call.Argument(0).String()
	retVal := This.runtime.NewObject()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		retVal.Set("err", This.runtime.NewGoError(err))
		return retVal
	}
	retVal.Set("data", This.runtime.ToValue(data))
	return retVal
}

func init() {
	require.RegisterNativeModule("file", func(runtime *goja.Runtime, module *goja.Object) {
		This := &_file{
			runtime: runtime,
		}
		o := module.Get("exports").(*goja.Object)
		o.Set("write", This.write)
		o.Set("read", This.read)
	})
}
