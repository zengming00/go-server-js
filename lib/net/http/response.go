package http

import (
	"fmt"
	"net/http"

	"github.com/dop251/goja"
)

type _resp struct {
	runtime *goja.Runtime
	w       http.ResponseWriter
}

func (This *_resp) header(call goja.FunctionCall) goja.Value {
	return NewHeader(This.runtime, This.w.Header())
}

func (This *_resp) write(call goja.FunctionCall) goja.Value {
	p1 := call.Argument(0).Export()
	retVal := This.runtime.NewObject()

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
			retVal.Set("err", fmt.Sprintf("response.write() can not convert to byte `data[%d] = %v`", i, v))
			return retVal
		}
	case []byte:
		data = t
	case string:
		data = []byte(t)
	default:
		retVal.Set("err", "response.write() data is not a byte array or string")
		return retVal
	}

	n, err := This.w.Write(data)
	if err != nil {
		retVal.Set("err", err.Error())
		return retVal
	}
	retVal.Set("value", n)
	return retVal
}

func (This *_resp) writeHeader(call goja.FunctionCall) goja.Value {
	n := call.Argument(0).ToInteger()
	This.w.WriteHeader(int(n))
	return nil
}

func (This *_resp) setCookie(call goja.FunctionCall) goja.Value {
	cookie := &http.Cookie{}
	cookie.Name = call.Argument(0).String()
	cookie.Value = call.Argument(1).String()
	cookie.Path = call.Argument(2).String()
	cookie.MaxAge = int(call.Argument(3).ToInteger())
	cookie.HttpOnly = call.Argument(4).ToBoolean()
	http.SetCookie(This.w, cookie)
	return nil
}

func (This *_resp) getPrototype(call goja.FunctionCall) goja.Value {
	return This.runtime.ToValue(This.w)
}

func NewResponse(runtime *goja.Runtime, w http.ResponseWriter) *goja.Object {
	This := &_resp{
		runtime: runtime,
		w:       w,
	}
	o := runtime.NewObject()
	o.Set("header", This.header)
	o.Set("write", This.write)
	o.Set("writeHeader", This.writeHeader)
	o.Set("setCookie", This.setCookie)
	o.Set("getPrototype", This.getPrototype)
	return o
}
