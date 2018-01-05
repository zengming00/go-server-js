package redis

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/garyburd/redigo/redis"
)

type _redis struct {
	runtime *goja.Runtime
}

func (This *_redis) stringFunc(call goja.FunctionCall) goja.Value {
	repl := call.Argument(0).Export()
	retVal := This.runtime.NewObject()
	str, err := redis.String(repl, nil)
	if err != nil {
		retVal.Set("err", err.Error())
		return retVal
	}
	retVal.Set("value", str)
	return retVal
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
