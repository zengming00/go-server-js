package image

import (
	"image"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/nodejs/require"
	"github.com/zengming00/go-server-js/lib/image/lib"
)

type _image struct {
	runtime *goja.Runtime
}

func (This *_image) rect(call goja.FunctionCall) goja.Value {
	x0 := int(call.Argument(0).ToInteger())
	y0 := int(call.Argument(1).ToInteger())
	x1 := int(call.Argument(2).ToInteger())
	y1 := int(call.Argument(3).ToInteger())
	// TODO
	rect := image.Rect(x0, y0, x1, y1)
	return This.runtime.ToValue(rect)
}

func (This *_image) newRGBA(call goja.FunctionCall) goja.Value {
	r := call.Argument(0).Export()
	if rect, ok := r.(image.Rectangle); ok {
		rgba := image.NewRGBA(rect)
		return This.runtime.ToValue(NewRGBA(This.runtime, rgba))
	}
	panic(This.runtime.NewTypeError("p0 is not a Rectangle type"))
}

func (This *_image) makeCapcha(call goja.FunctionCall) goja.Value {
	rgba := lib.MakeCapcha()
	return This.runtime.ToValue(NewRGBA(This.runtime, rgba))
}

func init() {
	require.RegisterNativeModule("image", func(runtime *goja.Runtime, module *goja.Object) {
		This := &_image{
			runtime: runtime,
		}
		o := module.Get("exports").(*goja.Object)
		o.Set("rect", This.rect)
		o.Set("newRGBA", This.newRGBA)
		o.Set("makeCapcha", This.makeCapcha)

	})
}
