package http

import "github.com/dop251/goja"
import "net/http"
import "fmt"

type responseRuntime struct {
	runtime *goja.Runtime
	w       *http.ResponseWriter
}

func (This *responseRuntime) header(call goja.FunctionCall) goja.Value {
	w := *This.w
	hd := w.Header()
	return newHeader(This.runtime, &hd)
}

func (This *responseRuntime) write(call goja.FunctionCall) goja.Value {
	p1 := call.Argument(0).Export()

	var data []byte
	switch t := p1.(type) {
	case []interface{}:
		data = make([]byte, len(t))
		for i, v := range t {
			if val, ok := v.(int64); ok {
				if val >= 0 && val <= 255 {
					data[i] = byte(val)
					continue
				}
			}
			panic(This.runtime.NewTypeError(fmt.Sprintf("response.write() can not convert to byte `data[%d] = %v`", i, v)))
		}
	case []byte:
		data = t
	case string:
		data = []byte(t)
	default:
		panic(This.runtime.NewTypeError("response.write() data is not a byte array"))
	}

	w := *This.w
	n, err := w.Write(data)
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}
	return This.runtime.ToValue(n)
}

func (This *responseRuntime) writeHeader(call goja.FunctionCall) goja.Value {
	n := call.Argument(0).ToInteger()
	w := *This.w
	w.WriteHeader(int(n))
	return nil
}

func NewResponse(runtime *goja.Runtime, w *http.ResponseWriter) *goja.Object {
	This := &responseRuntime{
		runtime: runtime,
		w:       w,
	}

	o := runtime.NewObject()
	o.Set("header", This.header)
	o.Set("write", This.write)
	o.Set("writeHeader", This.writeHeader)
	return o
}
