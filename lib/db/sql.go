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
	})
}
