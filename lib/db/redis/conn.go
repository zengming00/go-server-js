package redis

import (
	"github.com/dop251/goja"
	"github.com/garyburd/redigo/redis"
	"github.com/zengming00/go-server-js/lib"
)

type _conn struct {
	runtime *goja.Runtime
	conn    redis.Conn
}

func (This *_conn) close(call goja.FunctionCall) goja.Value {
	err := This.conn.Close()
	if err != nil {
		return This.runtime.ToValue(err.Error())
	}
	return nil
}

func (This *_conn) do(call goja.FunctionCall) goja.Value {
	commandName := call.Argument(0).String()
	args := lib.GetAllArgs(&call)
	reply, err := This.conn.Do(commandName, args[1:]...)
	return lib.MakeReturnValue(This.runtime, reply, err)
}

func NewConn(runtime *goja.Runtime, conn redis.Conn) *goja.Object {
	This := &_conn{
		runtime: runtime,
		conn:    conn,
	}
	o := runtime.NewObject()
	o.Set("do", This.do)
	o.Set("close", This.close)
	return o
}
