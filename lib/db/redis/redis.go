package redis

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/garyburd/redigo/redis"
	"github.com/zengming00/go-server-js/lib"
)

type _redis struct {
	runtime *goja.Runtime
}

func (This *_redis) stringFunc(call goja.FunctionCall) goja.Value {
	repl := call.Argument(0).Export()
	str, err := redis.String(repl, nil)
	return lib.MakeReturnValue(This.runtime, str, err)
}

func (This *_redis) dial(call goja.FunctionCall) goja.Value {
	network := call.Argument(0).String()
	address := call.Argument(1).String()
	retVal := This.runtime.NewObject()
	conn, err := redis.Dial(network, address)
	if err != nil {
		retVal.Set("err", err.Error())
		return retVal
	}
	retVal.Set("value", NewConn(This.runtime, conn))
	return retVal
}

func init() {
	require.RegisterNativeModule("redis", func(runtime *goja.Runtime, module *goja.Object) {
		This := &_redis{
			runtime: runtime,
		}
		o := module.Get("exports").(*goja.Object)
		o.Set("dial", This.dial)
		o.Set("string", This.stringFunc)
	})
}
