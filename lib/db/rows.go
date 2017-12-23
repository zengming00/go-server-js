package db

import (
	"database/sql"
	"errors"
	"reflect"

	"github.com/dop251/goja"
	"github.com/go-sql-driver/mysql"
)

type _rows struct {
	runtime *goja.Runtime
	rows    *sql.Rows
}

func (This *_rows) err(call goja.FunctionCall) goja.Value {
	err := This.rows.Err()
	return This.runtime.ToValue(err)
}

func (This *_rows) scan(call goja.FunctionCall) goja.Value {
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

func (This *_rows) next(call goja.FunctionCall) goja.Value {
	r := This.rows.Next()
	return This.runtime.ToValue(r)
}

func (This *_rows) close(call goja.FunctionCall) goja.Value {
	err := This.rows.Close()
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}
	return nil
}

func NewRows(runtime *goja.Runtime, rows *sql.Rows) *goja.Object {
	obj := &_rows{
		runtime: runtime,
		rows:    rows,
	}
	o := runtime.NewObject()
	o.Set("close", obj.close)
	o.Set("next", obj.next)
	o.Set("scan", obj.scan)
	o.Set("err", obj.err)
	return o
}
