package lib

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/nodejs/require"
)

func GetNativeType(runtime *goja.Runtime, call *goja.FunctionCall, idx int) interface{} {
	return call.Argument(idx).ToObject(runtime).Get("nativeType").Export()
}

func GetGoType(runtime *goja.Runtime, call *goja.FunctionCall, idx int) goja.Value {
	p := call.Argument(idx).ToObject(runtime)
	protoFunc, ok := goja.AssertFunction(p.Get("getGoType"))
	if !ok {
		panic(runtime.NewTypeError("p%d not have getGoType() function", idx))
	}
	obj, err := protoFunc(p)
	if err != nil {
		panic(runtime.NewGoError(err))
	}
	return obj
}

func GetAllArgs(call *goja.FunctionCall) []interface{} {
	args := make([]interface{}, len(call.Arguments))
	for i, v := range call.Arguments {
		args[i] = v.Export()
	}
	return args
}

func GetAllArgs_string(runtime *goja.Runtime, call *goja.FunctionCall) []string {
	args := make([]string, len(call.Arguments))
	for i, v := range call.Arguments {
		vv := v.Export()
		if s, ok := vv.(string); ok {
			args[i] = s
		} else {
			panic(runtime.NewTypeError("arg[%d] is not a string type:%T", i, v))
		}
	}
	return args
}

func MakeReturnValue(runtime *goja.Runtime, value interface{}) goja.Value {
	return runtime.ToValue(map[string]interface{}{
		"value": value,
	})
}

func MakeErrorValue(runtime *goja.Runtime, err error) goja.Value {
	return runtime.ToValue(map[string]interface{}{
		"err": NewError(runtime, err),
	})
}

func init() {
	require.RegisterNativeModule("utils", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("print", func(call goja.FunctionCall) goja.Value {
			fmt.Print(call.Argument(0).String())
			return nil
		})

		o.Set("panic", func(call goja.FunctionCall) goja.Value {
			panic(call.Argument(0).String())
		})

		o.Set("toString", func(call goja.FunctionCall) goja.Value {
			data := call.Argument(0).Export()
			if bts, ok := data.([]byte); ok {
				return runtime.ToValue(string(bts))
			}
			panic(runtime.NewTypeError("p0 is not []byte type:%T", data))
		})

		o.Set("toBase64", func(call goja.FunctionCall) goja.Value {
			p0 := call.Argument(0).Export()
			var str string
			switch data := p0.(type) {
			case []byte:
				str = base64.StdEncoding.EncodeToString(data)
			case string:
				str = base64.StdEncoding.EncodeToString([]byte(data))
			default:
				panic(runtime.NewTypeError("p0 is not []byte or string type:%T", p0))
			}
			return runtime.ToValue(str)
		})

		o.Set("md5", func(call goja.FunctionCall) goja.Value {
			p0 := call.Argument(0).Export()
			var r []byte
			switch data := p0.(type) {
			case []byte:
				tmp := md5.Sum(data)
				r = tmp[:]
			case string:
				tmp := md5.Sum([]byte(data))
				r = tmp[:]
			default:
				panic(runtime.NewTypeError("p0 is not []byte or string type:%T", p0))
			}
			return runtime.ToValue(hex.EncodeToString(r))
		})

		o.Set("sha1", func(call goja.FunctionCall) goja.Value {
			p0 := call.Argument(0).Export()
			var r []byte
			switch data := p0.(type) {
			case []byte:
				tmp := sha1.Sum(data)
				r = tmp[:]
			case string:
				tmp := sha1.Sum([]byte(data))
				r = tmp[:]
			default:
				panic(runtime.NewTypeError("p0 is not []byte or string type:%T", p0))
			}
			return runtime.ToValue(hex.EncodeToString(r))
		})
	})
}
