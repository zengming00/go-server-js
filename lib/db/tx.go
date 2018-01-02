package db

import (
	"database/sql"

	"github.com/dop251/goja"
)

type _tx struct {
	runtime *goja.Runtime
	tx      *sql.Tx
}

func (This *_tx) commit(call goja.FunctionCall) goja.Value {
	err := This.tx.Commit()
	if err != nil {
		return This.runtime.ToValue(err.Error())
	}
	return nil
}

func NewTx(runtime *goja.Runtime, tx *sql.Tx) *goja.Object {
	This := &_tx{
		runtime: runtime,
		tx:      tx,
	}
	o := runtime.NewObject()
	o.Set("commit", This.commit)
	return o
}
