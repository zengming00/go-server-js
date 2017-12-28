package image

import (
	"errors"
	"image"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
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
	return This.runtime.ToValue(&rect)
}

func (This *_image) newRGBA(call goja.FunctionCall) goja.Value {
	r := call.Argument(0).Export()
	if rect, ok := r.(*image.Rectangle); ok {
		rgba := image.NewRGBA(*rect)
		return This.runtime.ToValue(NewRGBA(This.runtime, rgba))
	}
	panic(errors.New("p0 is not a Rectangle type"))
}

func init() {
	require.RegisterNativeModule("image", func(runtime *goja.Runtime, module *goja.Object) {
		This := &_image{
			runtime: runtime,
		}
		o := module.Get("exports").(*goja.Object)
		o.Set("rect", This.rect)
		o.Set("newRGBA", This.newRGBA)
	})
}
