package db

import (
	"database/sql"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib"
)

func NewDB(runtime *goja.Runtime, db *sql.DB) *goja.Object {
	o := runtime.NewObject()
	o.Set("query", func(call goja.FunctionCall) goja.Value {
		args := lib.GetAllArgs(&call)
		if query, ok := args[0].(string); ok {
			var rows *sql.Rows
			var err error
			if len(args) == 1 {
				rows, err = db.Query(query)
			} else {
				rows, err = db.Query(query, args[1:]...)
			}
			if err != nil {
				return lib.MakeErrorValue(runtime, err)
			}
			return lib.MakeReturnValue(runtime, NewRows(runtime, rows))
		}
		panic(runtime.NewTypeError("p0 is not a string type:%T", args[0]))
	})

	o.Set("close", func(call goja.FunctionCall) goja.Value {
		err := db.Close()
		if err != nil {
			return runtime.ToValue(lib.NewError(runtime, err))
		}
		return nil
	})

	o.Set("exec", func(call goja.FunctionCall) goja.Value {
		args := lib.GetAllArgs(&call)

		if query, ok := args[0].(string); ok {
			var result sql.Result
			var err error
			if len(args) == 1 {
				result, err = db.Exec(query)
			} else {
				result, err = db.Exec(query, args[1:]...)
			}
			if err != nil {
				return lib.MakeErrorValue(runtime, err)
			}
			return lib.MakeReturnValue(runtime, NewResult(runtime, result))
		}
		panic(runtime.NewTypeError("p0 is not a string type:%T", args[0]))
	})

	o.Set("prepare", func(call goja.FunctionCall) goja.Value {
		query := call.Argument(0).String()
		stmt, err := db.Prepare(query)
		if err != nil {
			return lib.MakeErrorValue(runtime, err)
		}
		return lib.MakeReturnValue(runtime, NewStmt(runtime, stmt))
	})

	o.Set("stats", func(call goja.FunctionCall) goja.Value {
		st := db.Stats()
		o := runtime.NewObject()
		o.Set("openConnections", st.OpenConnections)
		return o
	})

	o.Set("begin", func(call goja.FunctionCall) goja.Value {
		tx, err := db.Begin()
		if err != nil {
			return lib.MakeErrorValue(runtime, err)
		}
		return lib.MakeReturnValue(runtime, NewTx(runtime, tx))
	})

	o.Set("setMaxOpenConns", func(call goja.FunctionCall) goja.Value {
		n := call.Argument(0).ToInteger()
		db.SetMaxOpenConns(int(n))
		return nil
	})

	o.Set("setMaxIdleConns", func(call goja.FunctionCall) goja.Value {
		n := call.Argument(0).ToInteger()
		db.SetMaxIdleConns(int(n))
		return nil
	})
	return o
}
