package png

import (
	"image"
	"image/png"
	"io"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type _png struct {
	runtime *goja.Runtime
}

func (This *_png) encode(call goja.FunctionCall) goja.Value {
	p0 := call.Argument(0).ToObject(This.runtime)
	p1 := call.Argument(1).ToObject(This.runtime)

	wproto, ok := goja.AssertFunction(p0.Get("getPrototype"))
	if !ok {
		panic(This.runtime.NewTypeError("p0 not have getPrototype() function"))
	}
	wo, err := wproto(p0)
	if err != nil {
		panic(err)
	}
	w, ok := wo.Export().(io.Writer)
	if !ok {
		panic(This.runtime.NewTypeError("p0 can not convert to *io.Writer"))
	}

	mproto, ok := goja.AssertFunction(p1.Get("getPrototype"))
	if !ok {
		panic(This.runtime.NewTypeError("p1 not have getPrototype() function"))
	}
	mo, err := mproto(p1)
	if err != nil {
		panic(err)
	}
	m, ok := mo.Export().(image.Image)
	if !ok {
		panic(This.runtime.NewTypeError("p0 can not convert to *image.Image"))
	}
	err = png.Encode(w, m)
	if err != nil {
		return This.runtime.ToValue(err.Error())
	}
	return nil
}

func init() {
	// TODO require存在bug windows 下不能使用/
	// require.RegisterNativeModule("image/png", func(runtime *goja.Runtime, module *goja.Object) {
	require.RegisterNativeModule("png", func(runtime *goja.Runtime, module *goja.Object) {
		This := &_png{
			runtime: runtime,
		}
		o := module.Get("exports").(*goja.Object)
		o.Set("encode", This.encode)
	})
}
