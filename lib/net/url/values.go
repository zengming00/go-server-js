package url

import (
	"net/url"

	"github.com/dop251/goja"
)

func NewValues(runtime *goja.Runtime, values url.Values) *goja.Object {
	o := runtime.NewObject()
	o.Set("del", func(call goja.FunctionCall) goja.Value {
		key := call.Argument(0).String()
		values.Del(key)
		return nil
	})

	o.Set("add", func(call goja.FunctionCall) goja.Value {
		key := call.Argument(0).String()
		value := call.Argument(1).String()
		values.Add(key, value)
		return nil
	})

	o.Set("encode", func(call goja.FunctionCall) goja.Value {
		str := values.Encode()
		return runtime.ToValue(str)
	})

	o.Set("get", func(call goja.FunctionCall) goja.Value {
		key := call.Argument(0).String()
		str := values.Get(key)
		return runtime.ToValue(str)
	})

	o.Set("gets", func(call goja.FunctionCall) goja.Value {
		key := call.Argument(0).String()
		values := values[key]
		return runtime.ToValue(values)
	})

	o.Set("getAll", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(map[string][]string(values))
	})

	o.Set("set", func(call goja.FunctionCall) goja.Value {
		key := call.Argument(0).String()
		value := call.Argument(1).String()
		values.Set(key, value)
		return nil
	})
	return o
}
