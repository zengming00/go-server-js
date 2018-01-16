package redis

import (
	"github.com/dop251/goja"
	"github.com/garyburd/redigo/redis"
	"github.com/zengming00/go-server-js/lib"
	"github.com/zengming00/go-server-js/nodejs/require"
)

func init() {
	require.RegisterNativeModule("redis", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("dial", func(call goja.FunctionCall) goja.Value {
			network := call.Argument(0).String()
			address := call.Argument(1).String()
			conn, err := redis.Dial(network, address)
			if err != nil {
				return lib.MakeErrorValue(runtime, err)
			}
			return lib.MakeReturnValue(runtime, NewConn(runtime, conn))
		})

		o.Set("string", func(call goja.FunctionCall) goja.Value {
			repl := call.Argument(0).Export()
			str, err := redis.String(repl, nil)
			if err != nil {
				return lib.MakeErrorValue(runtime, err)
			}
			return lib.MakeReturnValue(runtime, str)
		})
	})
}
