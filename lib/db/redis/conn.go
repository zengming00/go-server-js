package redis

import (
	"github.com/dop251/goja"
	"github.com/garyburd/redigo/redis"
	"github.com/zengming00/go-server-js/lib"
)

func NewConn(runtime *goja.Runtime, conn redis.Conn) *goja.Object {
	o := runtime.NewObject()
	o.Set("do", func(call goja.FunctionCall) goja.Value {
		commandName := call.Argument(0).String()
		args := lib.GetAllArgs(&call)
		reply, err := conn.Do(commandName, args[1:]...)
		if err != nil {
			return lib.MakeErrorValue(runtime, err)
		}
		return lib.MakeReturnValue(runtime, reply)
	})

	o.Set("close", func(call goja.FunctionCall) goja.Value {
		err := conn.Close()
		if err != nil {
			return lib.MakeErrorValue(runtime, err)
		}
		return nil
	})
	return o
}
