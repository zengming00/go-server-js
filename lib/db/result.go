package db

import (
	"database/sql"

	"github.com/dop251/goja"
)

type _result struct {
	runtime *goja.Runtime
	result  *sql.Result
}

func (This *_result) lastInsertId(call goja.FunctionCall) goja.Value {
	retVal := This.runtime.NewObject()
	id, err := (*This.result).LastInsertId()
	if err != nil {
		retVal.Set("err", err.Error())
		return retVal
	}
	retVal.Set("value", id)
	return retVal
}

func (This *_result) rowsAffected(call goja.FunctionCall) goja.Value {
	retVal := This.runtime.NewObject()
	rows, err := (*This.result).RowsAffected()
	if err != nil {
		retVal.Set("err", err.Error())
		return retVal
	}
	retVal.Set("value", rows)
	return retVal
}

func NewResult(runtime *goja.Runtime, result *sql.Result) *goja.Object {
	This := &_result{
		runtime: runtime,
		result:  result,
	}
	o := runtime.NewObject()
	o.Set("lastInsertId", This.lastInsertId)
	o.Set("rowsAffected", This.rowsAffected)
	return o
}
