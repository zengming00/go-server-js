package lib

import (
	"io"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type _io struct {
	runtime *goja.Runtime
}

func (This *_io) copy(call goja.FunctionCall) goja.Value {
	p0 := call.Argument(0).Export()
	p1 := call.Argument(1).Export()
	dst, ok := p0.(io.Writer)
	if !ok {
		panic(This.runtime.NewTypeError("p0 is not Writer: %T", p0))
	}
	src, ok := p1.(io.Reader)
	if !ok {
		panic(This.runtime.NewTypeError("p1 is not Reader: %T", p1))
	}
	written, err := io.Copy(dst, src)
	return This.runtime.ToValue(map[string]interface{}{
		"value": written,
		"err":   err,
	})
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
