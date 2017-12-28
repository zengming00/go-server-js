package lib

import (
	"errors"
	"fmt"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type _types struct {
	runtime *goja.Runtime
}

func scan(dest ...interface{}) {
	for i, v := range dest {
		if vv, ok := v.(*int); ok {
			*vv = i
		} else if vv, ok := v.(*bool); ok {
			*vv = true
		}
		fmt.Printf("%d, %T, %v\n", i, v, v)
	}
}

func (This *_types) scan(call goja.FunctionCall) goja.Value {
	args := GetAllArgs(&call)
	scan(args...)
	return nil
}

func (This *_types) err(call goja.FunctionCall) goja.Value {
	err := errors.New("test err")
	return This.runtime.ToValue(err.Error())
}

func (This *_types) newInt(call goja.FunctionCall) goja.Value {
	return This.runtime.ToValue(new(int))
}

func (This *_types) intValue(call goja.FunctionCall) goja.Value {
	v := call.Argument(0).Export()
	if vv, ok := v.(*int); ok {
		return This.runtime.ToValue(*vv)
	}
	panic(This.runtime.NewTypeError("%T can not convert to int value", v))
}

func (This *_types) newBool(call goja.FunctionCall) goja.Value {
	return This.runtime.ToValue(new(bool))
}

func (This *_types) boolValue(call goja.FunctionCall) goja.Value {
	v := call.Argument(0).Export()
	if vv, ok := v.(*bool); ok {
		return This.runtime.ToValue(*vv)
	}
	panic(This.runtime.NewTypeError("%T can not convert to bool value", v))
}

func (This *_types) newString(call goja.FunctionCall) goja.Value {
	return This.runtime.ToValue(new(string))
}

func (This *_types) stringValue(call goja.FunctionCall) goja.Value {
	v := call.Argument(0).Export()
	if vv, ok := v.(*string); ok {
		return This.runtime.ToValue(*vv)
	}
	panic(This.runtime.NewTypeError("%T can not convert to string value", v))
}

func init() {
	require.RegisterNativeModule("types", func(runtime *goja.Runtime, module *goja.Object) {
		This := &_types{
			runtime: runtime,
		}
		o := module.Get("exports").(*goja.Object)
		o.Set("newInt", This.newInt)
		o.Set("intValue", This.intValue)

		o.Set("newBool", This.newBool)
		o.Set("boolValue", This.boolValue)

		o.Set("newString", This.newString)
		o.Set("stringValue", This.stringValue)

		o.Set("scan", This.scan)
		o.Set("err", This.err)
	})
}
