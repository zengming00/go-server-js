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
			p0 := lib.GetNativeType(runtime, &call, 0)
			p1 := lib.GetNativeType(runtime, &call, 1)

			w, ok := p0.(io.Writer)
			if !ok {
				panic(runtime.NewTypeError("p0 is not io.Writer type:%T", p0))
			}

			m, ok := p1.(image.Image)
			if !ok {
				panic(runtime.NewTypeError("p1 is not image.Image type:%T", p1))
			}
			err := png.Encode(w, m)
			if err != nil {
				return lib.MakeErrorValue(runtime, err)
			}
			return nil
		})
	})
}
