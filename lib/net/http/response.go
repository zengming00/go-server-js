package http

import (
	"net/http"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib"
)

func NewResponse(runtime *goja.Runtime, w http.ResponseWriter) *goja.Object {
	o := runtime.NewObject()
	o.Set("header", func(call goja.FunctionCall) goja.Value {
		return NewHeader(runtime, w.Header())
	})

	o.Set("write", func(call goja.FunctionCall) goja.Value {
		p0 := call.Argument(0).Export()

		var data []byte
		switch t := p0.(type) {
		case []interface{}:
			data = make([]byte, len(t))
			for i, v := range t {
				if val, ok := v.(int64); ok {
					if val >= 0 && val <= 255 {
						data[i] = byte(val)
						continue
					}
				}
				panic(runtime.NewTypeError("can not convert to byte `data[%d]:%T`", i, v))
			}
		case []byte:
			data = t
		case string:
			data = []byte(t)
		default:
			panic(runtime.NewTypeError("data is not a []byte or string type:%T", t))
		}

		n, err := w.Write(data)
		if err != nil {
			return lib.MakeErrorValue(runtime, err)
		}
		return lib.MakeReturnValue(runtime, n)
	})

	o.Set("writeHeader", func(call goja.FunctionCall) goja.Value {
		n := call.Argument(0).ToInteger()
		w.WriteHeader(int(n))
		return nil
	})

	o.Set("setCookie", func(call goja.FunctionCall) goja.Value {
		cookie := &http.Cookie{}
		cookie.Name = call.Argument(0).String()
		cookie.Value = call.Argument(1).String()
		cookie.Path = call.Argument(2).String()
		cookie.MaxAge = int(call.Argument(3).ToInteger())
		cookie.HttpOnly = call.Argument(4).ToBoolean()
		http.SetCookie(w, cookie)
		return nil
	})

	o.Set("getPrototype", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(w)
	})

	return o
}
