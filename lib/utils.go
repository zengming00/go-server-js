package lib

import (
	"crypto/md5"
	"crypto/sha1"
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
	panic(This.runtime.NewTypeError("data is not a byte array"))
}

func (This *_utils) toBase64(call goja.FunctionCall) goja.Value {
	p0 := call.Argument(0).Export()
	var str string
	switch data := p0.(type) {
	case []byte:
		str = base64.StdEncoding.EncodeToString(data)
	case string:
		str = base64.StdEncoding.EncodeToString([]byte(data))
	default:
		panic(This.runtime.NewTypeError("data is not a byte array or string"))
	}
	return This.runtime.ToValue(str)
}

func (This *_utils) md5(call goja.FunctionCall) goja.Value {
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
		panic(This.runtime.NewTypeError("data is not a byte array or string"))
	}
	return This.runtime.ToValue(hex.EncodeToString(r))
}

func (This *_utils) sha1(call goja.FunctionCall) goja.Value {
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
		panic(This.runtime.NewTypeError("data is not a byte array or string"))
	}
	return This.runtime.ToValue(hex.EncodeToString(r))
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
		o.Set("sha1", This.sha1)
	})
}
