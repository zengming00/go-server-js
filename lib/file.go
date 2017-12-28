package lib

import (
	"io/ioutil"
	"os"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type _file struct {
	runtime *goja.Runtime
	file    *os.File
}

func (This *_file) write(call goja.FunctionCall) goja.Value {
	filename := call.Argument(0).String()
	data := call.Argument(1).Export()
	if bts, ok := data.([]byte); ok {
		err := ioutil.WriteFile(filename, bts, 0666)
		if err != nil {
			return This.runtime.ToValue(err)
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
		retVal.Set("err", err.Error())
		return retVal
	}
	retVal.Set("data", data)
	return retVal
}

func (This *_file) close(call goja.FunctionCall) goja.Value {
	err := This.file.Close()
	return This.runtime.ToValue(err)
}

func (This *_file) getPrototype(call goja.FunctionCall) goja.Value {
	return This.runtime.ToValue(This.file)
}

func (This *_file) writeString(call goja.FunctionCall) goja.Value {
	s := call.Argument(0).String()
	retVal := This.runtime.NewObject()

	n, err := This.file.WriteString(s)
	if err != nil {
		retVal.Set("err", err.Error())
		return retVal
	}
	retVal.Set("n", n)
	return retVal
}

func NewFile(runtime *goja.Runtime, file *os.File) *goja.Object {
	This := &_file{
		runtime: runtime,
		file:    file,
	}
	o := runtime.NewObject()
	o.Set("writeString", This.writeString)
	o.Set("close", This.close)
	o.Set("getPrototype", This.getPrototype)
	return o
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
