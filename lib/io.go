package lib

import (
	"io"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/nodejs/require"
)

func init() {
	require.RegisterNativeModule("io", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("copy", func(call goja.FunctionCall) goja.Value {
			p0 := GetNativeType(runtime, &call, 0)
			dst, ok := p0.(io.Writer)
			if !ok {
				panic(runtime.NewTypeError("p0 is not io.Writer type:%T", p0))
			}

			p1 := GetNativeType(runtime, &call, 1)
			src, ok := p1.(io.Reader)
			if !ok {
				panic(runtime.NewTypeError("p1 is not io.Reader type:%T", p1))
			}

			written, err := io.Copy(dst, src)
			if err != nil {
				return MakeErrorValue(runtime, err)
			}
			return MakeReturnValue(runtime, written)
		})
	})
}
