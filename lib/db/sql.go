package db

import (
	"database/sql"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib"
	"github.com/zengming00/go-server-js/nodejs/require"
)

var reUseDb *sql.DB

func init() {
	require.RegisterNativeModule("sql", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("open", func(call goja.FunctionCall) goja.Value {
			driverName := call.Argument(0).String()
			dataSourceName := call.Argument(1).String()
			reUse := call.Argument(2).ToBoolean()

			if reUse && reUseDb != nil {
				return runtime.ToValue(map[string]interface{}{
					"value":   NewDB(runtime, reUseDb),
					"isReUse": true,
				})
			}
			db, err := sql.Open(driverName, dataSourceName)
			if err != nil {
				return lib.MakeErrorValue(runtime, err)
			}
			if reUse && reUseDb == nil {
				reUseDb = db
			}
			return lib.MakeReturnValue(runtime, NewDB(runtime, db))
		})

		o.Set("drivers", func(call goja.FunctionCall) goja.Value {
			arr := sql.Drivers()
			return runtime.ToValue(arr)
		})

		o.Set("newNullString", func(call goja.FunctionCall) goja.Value {
			return runtime.ToValue(new(sql.NullString))
		})

		o.Set("nullStringValue", func(call goja.FunctionCall) goja.Value {
			v := call.Argument(0).Export()
			if vv, ok := v.(*sql.NullString); ok {
				return runtime.ToValue(map[string]interface{}{
					"value": vv.String,
					"valid": vv.Valid,
				})
			}
			panic(runtime.NewTypeError("p0 is not sql.NullString type:%T", v))
		})

		o.Set("newNullBool", func(call goja.FunctionCall) goja.Value {
			return runtime.ToValue(new(sql.NullBool))
		})

		o.Set("nullBoolValue", func(call goja.FunctionCall) goja.Value {
			v := call.Argument(0).Export()
			if vv, ok := v.(*sql.NullBool); ok {
				return runtime.ToValue(map[string]interface{}{
					"value": vv.Bool,
					"valid": vv.Valid,
				})
			}
			panic(runtime.NewTypeError("p0 is not sql.NullBool type:%T", v))
		})

		o.Set("newNullFloat64", func(call goja.FunctionCall) goja.Value {
			return runtime.ToValue(new(sql.NullFloat64))
		})

		o.Set("nullFloat64Value", func(call goja.FunctionCall) goja.Value {
			v := call.Argument(0).Export()
			if vv, ok := v.(*sql.NullFloat64); ok {
				return runtime.ToValue(map[string]interface{}{
					"value": vv.Float64,
					"valid": vv.Valid,
				})
			}
			panic(runtime.NewTypeError("p0 is not sql.NullFloat64 type:%T", v))
		})

		o.Set("newNullInt64", func(call goja.FunctionCall) goja.Value {
			return runtime.ToValue(new(sql.NullInt64))
		})

		o.Set("nullInt64Value", func(call goja.FunctionCall) goja.Value {
			v := call.Argument(0).Export()
			if vv, ok := v.(*sql.NullInt64); ok {
				return runtime.ToValue(map[string]interface{}{
					"value": vv.Int64,
					"valid": vv.Valid,
				})
			}
			panic(runtime.NewTypeError("p0 is not sql.NullInt64 type:%T", v))
		})

	})
}
