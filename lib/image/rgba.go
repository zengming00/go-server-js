package image

import (
	"image"
	"image/color"

	"github.com/dop251/goja"
)

type _rgba struct {
	runtime *goja.Runtime
	rgba    *image.RGBA
}

func (This *_rgba) getPrototype(call goja.FunctionCall) goja.Value {
	return This.runtime.ToValue(This.rgba)
}

func (This *_rgba) setRGBA(call goja.FunctionCall) goja.Value {
	x := int(call.Argument(0).ToInteger())
	y := int(call.Argument(1).ToInteger())
	This.rgba.SetRGBA(x, y, color.RGBA{
		R: uint8(call.Argument(2).ToInteger()),
		G: uint8(call.Argument(3).ToInteger()),
		B: uint8(call.Argument(4).ToInteger()),
		A: uint8(call.Argument(5).ToInteger()),
	})
	return nil
}

func NewRGBA(runtime *goja.Runtime, rgba *image.RGBA) *goja.Object {
	This := &_rgba{
		runtime: runtime,
		rgba:    rgba,
	}
	o := runtime.NewObject()
	o.Set("setRGBA", This.setRGBA)
	o.Set("getPrototype", This.getPrototype)
	return o
}
