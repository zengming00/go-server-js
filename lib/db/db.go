package db

import (
	"database/sql"
	"fmt"

	"github.com/dop251/goja"
)

type _db struct {
	runtime *goja.Runtime
	db      *sql.DB
}

func (This *_db) query(call goja.FunctionCall) goja.Value {
	query := call.Argument(0).String()
	retVal := This.runtime.NewObject()
	rows, err := This.db.Query(query)
	if err != nil {
		retVal.Set("err", err.Error())
		return retVal
	}
	retVal.Set("rows", NewRows(This.runtime, rows))
	return retVal
}

func (This *_db) close(call goja.FunctionCall) goja.Value {
	err := This.db.Close()
	return This.runtime.ToValue(err)
}

func (This *_db) exec(call goja.FunctionCall) goja.Value {
	// todo
	query := call.Argument(0).String()
	retVal := This.runtime.NewObject()
	r, err := This.db.Exec(query)
	if err != nil {
		fmt.Println(err)
		retVal.Set("err", err.Error())
		return retVal
	}
	fmt.Println(r)
	return nil
}

func (This *_db) prepare(call goja.FunctionCall) goja.Value {
	query := call.Argument(0).String()
	retVal := This.runtime.NewObject()
	stmt, err := This.db.Prepare(query)
	if err != nil {
		retVal.Set("err", err.Error())
		return retVal
	}
	retVal.Set("stmt", NewStmt(This.runtime, stmt))
	return retVal
}

func NewDB(runtime *goja.Runtime, db *sql.DB) *goja.Object {
	obj := &_db{
		runtime: runtime,
		db:      db,
	}
	o := runtime.NewObject()
	o.Set("query", obj.query)
	o.Set("close", obj.close)
	o.Set("exec", obj.exec)
	o.Set("prepare", obj.prepare)
	return o
}
