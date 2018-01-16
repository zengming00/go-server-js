package db

import (
	"database/sql"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib"
)

func NewTx(runtime *goja.Runtime, tx *sql.Tx) *goja.Object {
	// todo
	o := runtime.NewObject()
	o.Set("commit", func(call goja.FunctionCall) goja.Value {
		err := tx.Commit()
		if err != nil {
			return lib.MakeErrorValue(runtime, err)
		}
		return nil
	})
	return o
}
