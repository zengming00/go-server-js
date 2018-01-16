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

func (This *_rgba) drawLineH(call goja.FunctionCall) goja.Value {
	return nil
}

func (This *_rgba) drawLineV(call goja.FunctionCall) goja.Value {
	return nil
}

func (This *_rgba) drawLine(call goja.FunctionCall) goja.Value {
	return nil
}

func (This *_rgba) drawRect(call goja.FunctionCall) goja.Value {
	return nil
}

func (This *_rgba) fillRect(call goja.FunctionCall) goja.Value {
	return nil
}

func (This *_rgba) drawCircle(call goja.FunctionCall) goja.Value {
	return nil
}

func (This *_rgba) drawChar(call goja.FunctionCall) goja.Value {
	return nil
}

func (This *_rgba) drawString(call goja.FunctionCall) goja.Value {
	return nil
}

func NewRGBA(runtime *goja.Runtime, rgba *image.RGBA) *goja.Object {
	o := runtime.NewObject()
	o.Set("nativeType", rgba)

	o.Set("setRGBA", func(call goja.FunctionCall) goja.Value {
		x := int(call.Argument(0).ToInteger())
		y := int(call.Argument(1).ToInteger())
		rgba.SetRGBA(x, y, color.RGBA{
			R: uint8(call.Argument(2).ToInteger()),
			G: uint8(call.Argument(3).ToInteger()),
			B: uint8(call.Argument(4).ToInteger()),
			A: uint8(call.Argument(5).ToInteger()),
		})
		return nil
	})
	return o
}
