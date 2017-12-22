package lib

import (
	"database/sql"
	"errors"
	"reflect"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/go-sql-driver/mysql"
)

type sqlRuntime struct {
	runtime *goja.Runtime
}

type sqlObj struct {
	runtime *goja.Runtime
	db      *sql.DB
}

type rowsObj struct {
	runtime *goja.Runtime
	rows    *sql.Rows
}

func (This *rowsObj) err(call goja.FunctionCall) goja.Value {
	err := This.rows.Err()
	return This.runtime.ToValue(err)
}

func (This *rowsObj) scan(call goja.FunctionCall) goja.Value {
	rows := This.rows
	cols, err := rows.Columns()
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}
	ct, err := rows.ColumnTypes()
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}

	arr := make([]interface{}, len(ct))
	for i, v := range ct {
		t := v.ScanType()
		arr[i] = reflect.New(t).Interface()
	}

	err = rows.Scan(arr...)
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}

	m := make(map[string]interface{})
	for i, col := range cols {
		switch vv := arr[i].(type) {
		case *int32:
			m[col] = *vv
		case *sql.NullString:
			m[col] = *vv
		case *sql.NullBool:
			m[col] = *vv
		case *sql.NullFloat64:
			m[col] = *vv
		case *sql.NullInt64:
			m[col] = *vv
		case *sql.RawBytes:
			m[col] = string(*vv)
		case *mysql.NullTime:
			m[col] = *vv
		default:
			panic(This.runtime.NewGoError(errors.New("unknown type")))
		}
	}
	return This.runtime.ToValue(m)
}

func (This *rowsObj) next(call goja.FunctionCall) goja.Value {
	r := This.rows.Next()
	return This.runtime.ToValue(r)
}

func (This *rowsObj) close(call goja.FunctionCall) goja.Value {
	err := This.rows.Close()
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}
	return nil
}

func (This *sqlObj) query(call goja.FunctionCall) goja.Value {
	query := call.Argument(0).String()
	rows, err := This.db.Query(query)
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}
	obj := &rowsObj{
		runtime: This.runtime,
		rows:    rows,
	}
	o := This.runtime.NewObject()
	o.Set("close", obj.close)
	o.Set("next", obj.next)
	o.Set("scan", obj.scan)
	o.Set("err", obj.err)
	return o
}

func (This *sqlObj) close(call goja.FunctionCall) goja.Value {
	err := This.db.Close()
	return This.runtime.ToValue(err)
}

func (This *sqlRuntime) newFunc(call goja.FunctionCall) goja.Value {
	driverName := call.Argument(0).String()
	dataSourceName := call.Argument(1).String()
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}
	obj := &sqlObj{
		runtime: This.runtime,
		db:      db,
	}
	o := This.runtime.NewObject()
	o.Set("query", obj.query)
	o.Set("close", obj.close)
	return o
}

func init() {
	require.RegisterNativeModule("sql", func(runtime *goja.Runtime, module *goja.Object) {
		This := &sqlRuntime{
			runtime: runtime,
		}

		o := module.Get("exports").(*goja.Object)
		o.Set("new", This.newFunc)
	})
}
