package lib

import (
	"io/ioutil"
	"os"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/nodejs/require"
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
			return MakeErrorValue(This.runtime, err)
		}
		return nil
	}
	panic(This.runtime.NewTypeError("p1 is not []byte type:%T", data))
}

func (This *_file) read(call goja.FunctionCall) goja.Value {
	filename := call.Argument(0).String()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return MakeErrorValue(This.runtime, err)
	}
	return MakeReturnValue(This.runtime, data)
}

func (This *_file) close(call goja.FunctionCall) goja.Value {
	err := This.file.Close()
	if err != nil {
		return MakeErrorValue(This.runtime, err)
	}
	return nil
}

func (This *_file) getPrototype(call goja.FunctionCall) goja.Value {
	return This.runtime.ToValue(This.file)
}

func (This *_file) writeString(call goja.FunctionCall) goja.Value {
	s := call.Argument(0).String()
	n, err := This.file.WriteString(s)
	if err != nil {
		return MakeErrorValue(This.runtime, err)
	}
	return MakeReturnValue(This.runtime, n)
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
