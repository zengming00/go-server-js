package png

import (
	"image"
	"image/png"
	"io"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib"
	"github.com/zengming00/go-server-js/nodejs/require"
)

func init() {
	require.RegisterNativeModule("image/png", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("encode", func(call goja.FunctionCall) goja.Value {
			p0 := call.Argument(0).ToObject(runtime)
			p1 := call.Argument(1).ToObject(runtime)

			wproto, ok := goja.AssertFunction(p0.Get("getPrototype"))
			if !ok {
				panic(runtime.NewTypeError("p0 not have getPrototype() function"))
			}
			wo, err := wproto(p0)
			if err != nil {
				panic(runtime.NewGoError(err))
			}
			w, ok := wo.Export().(io.Writer)
			if !ok {
				panic(runtime.NewTypeError("p0 is not io.Writer type:%T", wo.Export()))
			}

			mproto, ok := goja.AssertFunction(p1.Get("getPrototype"))
			if !ok {
				panic(runtime.NewTypeError("p1 not have getPrototype() function"))
			}
			mo, err := mproto(p1)
			if err != nil {
				panic(runtime.NewGoError(err))
			}
			m, ok := mo.Export().(image.Image)
			if !ok {
				panic(runtime.NewTypeError("p1 is not image.Image type:%T", mo.Export()))
			}
			err = png.Encode(w, m)
			if err != nil {
				return lib.MakeErrorValue(runtime, err)
			}
			return nil
		})
	})
}
