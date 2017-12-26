package db

import (
	"database/sql"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib"
)

type _stmt struct {
	runtime *goja.Runtime
	stmt    *sql.Stmt
}

func (This *_stmt) exec(call goja.FunctionCall) goja.Value {
	args := lib.GetAllArgs(&call)
	retVal := This.runtime.NewObject()
	result, err := This.stmt.Exec(args...)
	if err != nil {
		retVal.Set("err", err.Error())
		return retVal
	}
	retVal.Set("result", NewResult(This.runtime, &result))
	return retVal
}

func NewStmt(runtime *goja.Runtime, stmt *sql.Stmt) *goja.Object {
	obj := &_stmt{
		runtime: runtime,
		stmt:    stmt,
	}
	o := runtime.NewObject()
	o.Set("exec", obj.exec)
	return o
}
