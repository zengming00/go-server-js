package db

import (
	"database/sql"
	"reflect"

	"github.com/dop251/goja"
	"github.com/go-sql-driver/mysql"
	"github.com/zengming00/go-server-js/lib"
)

func NewRows(runtime *goja.Runtime, rows *sql.Rows) *goja.Object {
	o := runtime.NewObject()
	o.Set("err", func(call goja.FunctionCall) goja.Value {
		err := rows.Err()
		if err != nil {
			return runtime.ToValue(lib.NewError(runtime, err))
		}
		return nil
	})

	o.Set("getData", func(call goja.FunctionCall) goja.Value {
		rows := rows
		cols, err := rows.Columns()
		if err != nil {
			panic(runtime.NewGoError(err))
		}
		ct, err := rows.ColumnTypes()
		if err != nil {
			panic(runtime.NewGoError(err))
		}

		arr := make([]interface{}, len(ct))
		for i, v := range ct {
			t := v.ScanType()
			arr[i] = reflect.New(t).Interface()
		}

		err = rows.Scan(arr...)
		if err != nil {
			panic(runtime.NewGoError(err))
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
				panic(runtime.NewTypeError("unknown type %T", vv))
			}
		}
		return runtime.ToValue(m)
	})

	o.Set("scan", func(call goja.FunctionCall) goja.Value {
		args := lib.GetAllArgs(&call)
		err := rows.Scan(args...)
		if err != nil {
			return runtime.ToValue(lib.NewError(runtime, err))
		}
		return nil
	})

	o.Set("next", func(call goja.FunctionCall) goja.Value {
		r := rows.Next()
		return runtime.ToValue(r)
	})

	o.Set("close", func(call goja.FunctionCall) goja.Value {
		err := rows.Close()
		if err != nil {
			return runtime.ToValue(lib.NewError(runtime, err))
		}
		return nil
	})
	return o
}
