package db

import (
	"database/sql"

	"github.com/dop251/goja"
)

type _db struct {
	runtime *goja.Runtime
	db      *sql.DB
}

func (This *_db) query(call goja.FunctionCall) goja.Value {
	query := call.Argument(0).String()
	rows, err := This.db.Query(query)
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}
	return NewRows(This.runtime, rows)
}

func (This *_db) close(call goja.FunctionCall) goja.Value {
	err := This.db.Close()
	return This.runtime.ToValue(err)
}

func NewDB(runtime *goja.Runtime, db *sql.DB) *goja.Object {
	obj := &_db{
		runtime: runtime,
		db:      db,
	}
	o := runtime.NewObject()
	o.Set("query", obj.query)
	o.Set("close", obj.close)
	return o
}
