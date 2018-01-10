package db

import (
	"database/sql"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib"
)

type _db struct {
	runtime *goja.Runtime
	db      *sql.DB
}

func (This *_db) query(call goja.FunctionCall) goja.Value {
	args := lib.GetAllArgs(&call)
	retVal := This.runtime.NewObject()

	if query, ok := args[0].(string); ok {
		var rows *sql.Rows
		var err error
		if len(args) == 1 {
			rows, err = This.db.Query(query)
		} else {
			rows, err = This.db.Query(query, args[1:]...)
		}
		if err != nil {
			retVal.Set("err", err.Error())
			return retVal
		}
		retVal.Set("value", NewRows(This.runtime, rows))
	} else {
		retVal.Set("err", "p0 is not a string")
	}
	return retVal
}

func (This *_db) close(call goja.FunctionCall) goja.Value {
	err := This.db.Close()
	if err != nil {
		return This.runtime.ToValue(err.Error())
	}
	return nil
}

func (This *_db) exec(call goja.FunctionCall) goja.Value {
	args := lib.GetAllArgs(&call)

	if query, ok := args[0].(string); ok {
		var result sql.Result
		var err error
		if len(args) == 1 {
			result, err = This.db.Exec(query)
		} else {
			result, err = This.db.Exec(query, args[1:]...)
		}
		if err != nil {
			return This.runtime.ToValue(map[string]interface{}{
				"err": err.Error(),
			})
		}
		return This.runtime.ToValue(map[string]interface{}{
			"value": NewResult(This.runtime, result),
		})
	}
	panic(This.runtime.NewTypeError("p0 is not a string"))
}

func (This *_db) prepare(call goja.FunctionCall) goja.Value {
	query := call.Argument(0).String()
	retVal := This.runtime.NewObject()
	stmt, err := This.db.Prepare(query)
	if err != nil {
		retVal.Set("err", err.Error())
		return retVal
	}
	retVal.Set("value", NewStmt(This.runtime, stmt))
	return retVal
}

func (This *_db) stats(call goja.FunctionCall) goja.Value {
	st := This.db.Stats()
	o := This.runtime.NewObject()
	o.Set("openConnections", st.OpenConnections)
	return o
}

func (This *_db) begin(call goja.FunctionCall) goja.Value {
	retVal := This.runtime.NewObject()
	tx, err := This.db.Begin()
	if err != nil {
		retVal.Set("err", err.Error())
		return retVal
	}
	retVal.Set("value", NewTx(This.runtime, tx))
	return retVal
}

func (This *_db) setMaxOpenConns(call goja.FunctionCall) goja.Value {
	n := call.Argument(0).ToInteger()
	This.db.SetMaxOpenConns(int(n))
	return nil
}

func (This *_db) setMaxIdleConns(call goja.FunctionCall) goja.Value {
	n := call.Argument(0).ToInteger()
	This.db.SetMaxIdleConns(int(n))
	return nil
}

func NewDB(runtime *goja.Runtime, db *sql.DB) *goja.Object {
	This := &_db{
		runtime: runtime,
		db:      db,
	}
	o := runtime.NewObject()
	o.Set("query", This.query)
	o.Set("close", This.close)
	o.Set("exec", This.exec)
	o.Set("prepare", This.prepare)
	o.Set("stats", This.stats)
	o.Set("begin", This.begin)
	o.Set("setMaxOpenConns", This.setMaxOpenConns)
	o.Set("setMaxIdleConns", This.setMaxIdleConns)
	return o
}
