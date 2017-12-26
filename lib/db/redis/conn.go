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
	length := len(call.Arguments)
	for i := 1; i < length; i++ {
		args = append(args, call.Argument(i).Export())
	}
	retVal := This.runtime.NewObject()
	reply, err := (*This.conn).Do(commandName, args...)
	if err != nil {
		retVal.Set("err", err.Error())
		return retVal
	}
	retVal.Set("reply", This.runtime.ToValue(reply))
	return retVal
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
