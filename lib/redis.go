package lib

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/garyburd/redigo/redis"
)

type redisRuntime struct {
	runtime *goja.Runtime
}

type redisObj struct {
	runtime *goja.Runtime
	conn    *redis.Conn
}

func (This *redisObj) close(call goja.FunctionCall) goja.Value {
	err := (*This.conn).Close()
	return This.runtime.ToValue(err)
}

func (This *redisObj) do(call goja.FunctionCall) goja.Value {
	var args []interface{}
	commandName := call.Argument(0).String()
	len := len(call.Arguments)
	for i := 1; i < len; i++ {
		args = append(args, call.Argument(i).Export())
	}
	reply, err := (*This.conn).Do(commandName, args...)
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}
	return This.runtime.ToValue(reply)
}

func (This *redisRuntime) stringFunc(call goja.FunctionCall) goja.Value {
	repl := call.Argument(0).Export()
	str, err := redis.String(repl, nil)
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}
	return This.runtime.ToValue(str)
}

func (This *redisRuntime) dial(call goja.FunctionCall) goja.Value {
	network := call.Argument(0).String()
	address := call.Argument(1).String()
	conn, err := redis.Dial(network, address)
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}
	obj := &redisObj{
		runtime: This.runtime,
		conn:    &conn,
	}
	o := This.runtime.NewObject()
	o.Set("do", obj.do)
	o.Set("close", obj.close)
	return o
}

func init() {
	require.RegisterNativeModule("redis", func(runtime *goja.Runtime, module *goja.Object) {
		This := &redisRuntime{
			runtime: runtime,
		}

		o := module.Get("exports").(*goja.Object)
		o.Set("dial", This.dial)
		o.Set("string", This.stringFunc)
	})
}
