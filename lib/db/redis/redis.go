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
	str, err := redis.String(repl, nil)
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}
	return This.runtime.ToValue(str)
}

func (This *_redis) dial(call goja.FunctionCall) goja.Value {
	network := call.Argument(0).String()
	address := call.Argument(1).String()
	conn, err := redis.Dial(network, address)
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}
	return NewConn(This.runtime, &conn)
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
