package lib

import (
	"io"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/nodejs/require"
)

type _io struct {
	runtime *goja.Runtime
}

func (This *_io) copy(call goja.FunctionCall) goja.Value {
	p0 := call.Argument(0).ToObject(This.runtime)
	p0proto, ok := goja.AssertFunction(p0.Get("getPrototype"))
	if !ok {
		panic(This.runtime.NewTypeError("p0 not have getPrototype() function"))
	}
	p0obj, err := p0proto(p0)
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}
	dst, ok := p0obj.Export().(io.Writer)
	if !ok {
		panic(This.runtime.NewTypeError("p0 is not Writer: %T", p0))
	}

	p1 := call.Argument(1).ToObject(This.runtime)
	p1proto, ok := goja.AssertFunction(p1.Get("getPrototype"))
	if !ok {
		panic(This.runtime.NewTypeError("p1 not have getPrototype() function"))
	}
	p1obj, err := p1proto(p1)
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}
	src, ok := p1obj.Export().(io.Reader)
	if !ok {
		panic(This.runtime.NewTypeError("p1 is not Reader: %T", p1))
	}

	written, err := io.Copy(dst, src)
	return MakeReturnValue(This.runtime, written, err)
}

func init() {
	require.RegisterNativeModule("io", func(runtime *goja.Runtime, module *goja.Object) {
		This := &_io{
			runtime: runtime,
		}
		o := module.Get("exports").(*goja.Object)
		o.Set("copy", This.copy)
	})
}
