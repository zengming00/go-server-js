package db

import (
	"database/sql"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib"
)

func NewStmt(runtime *goja.Runtime, stmt *sql.Stmt) *goja.Object {
	o := runtime.NewObject()
	o.Set("exec", func(call goja.FunctionCall) goja.Value {
		args := lib.GetAllArgs(&call)
		result, err := stmt.Exec(args...)
		if err != nil {
			return lib.MakeErrorValue(runtime, err)
		}
		return lib.MakeReturnValue(runtime, NewResult(runtime, result))
	})
	return o
}
