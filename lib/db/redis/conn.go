package redis

import (
	"github.com/dop251/goja"
	"github.com/garyburd/redigo/redis"
)

type _conn struct {
	runtime *goja.Runtime
	conn    *redis.Conn
}

func (This *_conn) close(call goja.FunctionCall) goja.Value {
	err := (*This.conn).Close()
	return This.runtime.ToValue(err)
}

func (This *_conn) do(call goja.FunctionCall) goja.Value {
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

func NewConn(runtime *goja.Runtime, conn *redis.Conn) *goja.Object {
	obj := &_conn{
		runtime: runtime,
		conn:    conn,
	}
	o := runtime.NewObject()
	o.Set("do", obj.do)
	o.Set("close", obj.close)
	return o
}
