package url

import (
	"net/url"

	"github.com/dop251/goja"
)

type _values struct {
	runtime *goja.Runtime
	values  *url.Values
}

func (This *_values) del(call goja.FunctionCall) goja.Value {
	key := call.Argument(0).String()
	This.values.Del(key)
	return nil
}

func (This *_values) add(call goja.FunctionCall) goja.Value {
	key := call.Argument(0).String()
	value := call.Argument(1).String()
	This.values.Add(key, value)
	return nil
}

func (This *_values) encode(call goja.FunctionCall) goja.Value {
	str := This.values.Encode()
	return This.runtime.ToValue(str)
}

func (This *_values) get(call goja.FunctionCall) goja.Value {
	key := call.Argument(0).String()
	str := This.values.Get(key)
	return This.runtime.ToValue(str)
}

func (This *_values) set(call goja.FunctionCall) goja.Value {
	key := call.Argument(0).String()
	value := call.Argument(1).String()
	This.values.Set(key, value)
	return nil
}

func NewValues(runtime *goja.Runtime, values *url.Values) *goja.Object {
	This := &_values{
		runtime: runtime,
		values:  values,
	}
	o := runtime.NewObject()
	o.Set("del", This.del)
	o.Set("add", This.add)
	o.Set("encode", This.encode)
	o.Set("get", This.get)
	o.Set("set", This.set)
	return o
}
