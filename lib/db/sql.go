package db

import (
	"database/sql"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type _sql struct {
	runtime *goja.Runtime
}

func (This *_sql) open(call goja.FunctionCall) goja.Value {
	driverName := call.Argument(0).String()
	dataSourceName := call.Argument(1).String()
	retVal := This.runtime.NewObject()
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		retVal.Set("err", err.Error())
		return retVal
	}
	retVal.Set("value", NewDB(This.runtime, db))
	return retVal
}

func (This *_sql) drivers(call goja.FunctionCall) goja.Value {
	arr := sql.Drivers()
	return This.runtime.ToValue(arr)
}

func init() {
	require.RegisterNativeModule("sql", func(runtime *goja.Runtime, module *goja.Object) {
		This := &_sql{
			runtime: runtime,
		}
		o := module.Get("exports").(*goja.Object)
		o.Set("open", This.open)
		o.Set("drivers", This.drivers)
	})
}
