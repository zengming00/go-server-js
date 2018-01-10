package lib

import (
	"time"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/nodejs/require"
)

type _time struct {
	runtime *goja.Runtime
	t       *time.Time
}

func (This *_time) stringFunc(call goja.FunctionCall) goja.Value {
	str := This.t.String()
	return This.runtime.ToValue(str)
}

func (This *_time) sleep(call goja.FunctionCall) goja.Value {
	d := call.Argument(0).ToInteger()
	<-time.After(time.Duration(d) * time.Millisecond)
	return nil
}

func (This *_time) nowString(call goja.FunctionCall) goja.Value {
	const DATE_TIME_FORMAT = "2006-01-02 15:04:05"
	cst := time.FixedZone("CST", 28800)
	str := time.Now().In(cst).Format(DATE_TIME_FORMAT)
	return This.runtime.ToValue(str)
}

func NewTime(runtime *goja.Runtime, t *time.Time) *goja.Object {
	obj := &_time{
		runtime: runtime,
		t:       t,
	}

	o := runtime.NewObject()
	o.Set("string", obj.stringFunc)
	return o
}

func init() {
	require.RegisterNativeModule("time", func(runtime *goja.Runtime, module *goja.Object) {
		This := &_time{
			runtime: runtime,
		}
		o := module.Get("exports").(*goja.Object)
		o.Set("sleep", This.sleep)
		o.Set("nowString", This.nowString)
	})
}
