package lib

import (
	"time"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/nodejs/require"
)

func NewTime(runtime *goja.Runtime, t *time.Time) *goja.Object {
	o := runtime.NewObject()
	o.Set("string", func(call goja.FunctionCall) goja.Value {
		str := t.String()
		return runtime.ToValue(str)
	})

	return o
}

func init() {
	require.RegisterNativeModule("time", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("sleep", func(call goja.FunctionCall) goja.Value {
			d := call.Argument(0).ToInteger()
			<-time.After(time.Duration(d) * time.Millisecond)
			return nil
		})

		o.Set("nowString", func(call goja.FunctionCall) goja.Value {
			const DATE_TIME_FORMAT = "2006-01-02 15:04:05"
			cst := time.FixedZone("CST", 28800)
			str := time.Now().In(cst).Format(DATE_TIME_FORMAT)
			return runtime.ToValue(str)
		})
	})
}
