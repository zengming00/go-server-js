package lib

import (
	"errors"
	"fmt"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/nodejs/require"
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
	// 无法在js中使用
	return This.runtime.ToValue(err)
}

func (This *_types) err2(call goja.FunctionCall) goja.Value {
	err := errors.New("test err2")
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

func (This *_types) makeByteSlice(call goja.FunctionCall) goja.Value {
	mLen := call.Argument(0).ToInteger()
	mCap := mLen
	if len(call.Arguments) != 1 {
		mCap = call.Argument(1).ToInteger()
	}
	v := make([]byte, mLen, mCap)
	return This.runtime.ToValue(v)
}

func (This *_types) test(call goja.FunctionCall) goja.Value {
	v := call.Argument(0).Export()
	fmt.Printf("%T %[1]v\n", v)
	return nil
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

		o.Set("makeByteSlice", This.makeByteSlice)

		o.Set("scan", This.scan)
		o.Set("test", This.test)
		o.Set("err", This.err)
		o.Set("err2", This.err2)

		o.Set("retNil", func(call goja.FunctionCall) goja.Value {
			// nil 和 goja.Undefined() 效果相同，在js中都是undefined
			return nil
		})

		o.Set("retNull", func(call goja.FunctionCall) goja.Value {
			// 在js中是null
			return goja.Null()
		})

		o.Set("retUndefined", func(call goja.FunctionCall) goja.Value {
			return goja.Undefined()
		})
	})
}
