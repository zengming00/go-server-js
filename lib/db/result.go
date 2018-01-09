package db

import (
	"database/sql"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib"
)

type _result struct {
	runtime *goja.Runtime
	result  sql.Result
}

func (This *_result) lastInsertId(call goja.FunctionCall) goja.Value {
	id, err := This.result.LastInsertId()
	return lib.MakeReturnValue(This.runtime, id, err)
}

func (This *_result) rowsAffected(call goja.FunctionCall) goja.Value {
	rows, err := This.result.RowsAffected()
	return lib.MakeReturnValue(This.runtime, rows, err)
}

func NewResult(runtime *goja.Runtime, result sql.Result) *goja.Object {
	This := &_result{
		runtime: runtime,
		result:  result,
	}
	o := runtime.NewObject()
	o.Set("lastInsertId", This.lastInsertId)
	o.Set("rowsAffected", This.rowsAffected)
	return o
}
