package lib

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type _utils struct {
	runtime *goja.Runtime
}

func GetAllArgs(call *goja.FunctionCall) []interface{} {
	args := make([]interface{}, len(call.Arguments))
	for i, v := range call.Arguments {
		args[i] = v.Export()
	}
	return args
}

func (This *_utils) print(call goja.FunctionCall) goja.Value {
	fmt.Print(call.Argument(0).String())
	return nil
}

func (This *_utils) toString(call goja.FunctionCall) goja.Value {
	data := call.Argument(0).Export()
	if bts, ok := data.([]byte); ok {
		return This.runtime.ToValue(string(bts))
	}
	panic(This.runtime.NewTypeError("toString data is not a byte array"))
}

func (This *_utils) toBase64(call goja.FunctionCall) goja.Value {
	data := call.Argument(0).Export()
	if bts, ok := data.([]byte); ok {
		str := base64.StdEncoding.EncodeToString(bts)
		return This.runtime.ToValue(str)
	} else if s, ok := data.(string); ok {
		str := base64.StdEncoding.EncodeToString([]byte(s))
		return This.runtime.ToValue(str)
	}
	panic(This.runtime.NewTypeError("toBase64 data is not a byte array or string"))
}

func (This *_utils) md5(call goja.FunctionCall) goja.Value {
	data := call.Argument(0).Export()
	if bts, ok := data.([]byte); ok {
		r := md5.Sum(bts)
		str := hex.EncodeToString(r[:])
		return This.runtime.ToValue(str)
	} else if s, ok := data.(string); ok {
		r := md5.Sum([]byte(s))
		str := hex.EncodeToString(r[:])
		return This.runtime.ToValue(str)
	}
	panic(This.runtime.NewTypeError("md5 data is not a byte array or string"))
}

func init() {
	require.RegisterNativeModule("utils", func(runtime *goja.Runtime, module *goja.Object) {
		This := &_utils{
			runtime: runtime,
		}
		o := module.Get("exports").(*goja.Object)
		o.Set("print", This.print)
		o.Set("toString", This.toString)
		o.Set("toBase64", This.toBase64)
		o.Set("md5", This.md5)
	})
}
