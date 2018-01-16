package db

import (
	"database/sql"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib"
)

func NewResult(runtime *goja.Runtime, result sql.Result) *goja.Object {
	o := runtime.NewObject()
	o.Set("lastInsertId", func(call goja.FunctionCall) goja.Value {
		id, err := result.LastInsertId()
		if err != nil {
			return lib.MakeErrorValue(runtime, err)
		}
		return lib.MakeReturnValue(runtime, id)
	})

	o.Set("rowsAffected", func(call goja.FunctionCall) goja.Value {
		rows, err := result.RowsAffected()
		if err != nil {
			return lib.MakeErrorValue(runtime, err)
		}
		return lib.MakeReturnValue(runtime, rows)
	})
	return o
}
