package db

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	"github.com/dop251/goja"
	"github.com/go-sql-driver/mysql"
	"github.com/zengming00/go-server-js/lib"
)

type _rows struct {
	runtime *goja.Runtime
	rows    *sql.Rows
}

func (This *_rows) err(call goja.FunctionCall) goja.Value {
	err := This.rows.Err()
	return This.runtime.ToValue(err)
}

func (This *_rows) scan2(call goja.FunctionCall) goja.Value {
	rows := This.rows
	cols, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	ct, err := rows.ColumnTypes()
	if err != nil {
		panic(err)
	}

	arr := make([]interface{}, len(ct))
	for i, v := range ct {
		t := v.ScanType()
		arr[i] = reflect.New(t).Interface()
	}

	err = rows.Scan(arr...)
	if err != nil {
		panic(err)
	}

	m := make(map[string]interface{})
	for i, col := range cols {
		switch vv := arr[i].(type) {
		case *int64:
			m[col] = *vv
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
		case *string:
			m[col] = *vv
		default:
			fmt.Printf("%T\n", vv)
			panic(errors.New("unknown type"))
		}
	}
	return This.runtime.ToValue(m)
}

func (This *_rows) scan(call goja.FunctionCall) goja.Value {
	args := lib.GetAllArgs(&call)
	err := This.rows.Scan(args...)
	return This.runtime.ToValue(err)
}

func (This *_rows) next(call goja.FunctionCall) goja.Value {
	r := This.rows.Next()
	return This.runtime.ToValue(r)
}

func (This *_rows) close(call goja.FunctionCall) goja.Value {
	err := This.rows.Close()
	return This.runtime.ToValue(err)
}

func NewRows(runtime *goja.Runtime, rows *sql.Rows) *goja.Object {
	This := &_rows{
		runtime: runtime,
		rows:    rows,
	}
	o := runtime.NewObject()
	o.Set("close", This.close)
	o.Set("next", This.next)
	o.Set("scan", This.scan)
	o.Set("err", This.err)
	return o
}
