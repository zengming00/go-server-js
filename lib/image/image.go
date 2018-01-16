package image

import (
	"image"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib/image/lib"
	"github.com/zengming00/go-server-js/nodejs/require"
)

func init() {
	require.RegisterNativeModule("image", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("rect", func(call goja.FunctionCall) goja.Value {
			x0 := int(call.Argument(0).ToInteger())
			y0 := int(call.Argument(1).ToInteger())
			x1 := int(call.Argument(2).ToInteger())
			y1 := int(call.Argument(3).ToInteger())
			// TODO
			rect := image.Rect(x0, y0, x1, y1)
			return runtime.ToValue(rect)
		})

		o.Set("newRGBA", func(call goja.FunctionCall) goja.Value {
			r := call.Argument(0).Export()
			if rect, ok := r.(image.Rectangle); ok {
				rgba := image.NewRGBA(rect)
				return runtime.ToValue(NewRGBA(runtime, rgba))
			}
			panic(runtime.NewTypeError("p0 is not a image.Rectangle type:%T", r))
		})

		o.Set("makeCapcha", func(call goja.FunctionCall) goja.Value {
			rgba := lib.MakeCapcha()
			return runtime.ToValue(NewRGBA(runtime, rgba))
		})

	})
}
