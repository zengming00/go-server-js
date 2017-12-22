package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type httpRuntime struct {
	runtime *goja.Runtime
}

func (This *httpRuntime) request(call goja.FunctionCall) goja.Value {
	method := call.Argument(0).String()
	url := call.Argument(1).String()
	headers := call.Argument(2).Export().(map[string]interface{})
	body := call.Argument(3).String()
	timeout := call.Argument(4).ToInteger()

	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}

	for k, v := range headers {
		if str, ok := v.(string); ok {
			req.Header.Set(k, str)
		} else {
			panic(This.runtime.NewTypeError(fmt.Sprintf("header[" + k + "] is not string")))
		}
	}

	client := &http.Client{}
	client.Timeout = time.Duration(timeout) * time.Millisecond
	res, err := client.Do(req)
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}
	defer res.Body.Close()
	datas, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(This.runtime.NewGoError(err))
	}

	headerObj := This.runtime.NewObject()
	for k, v := range res.Header {
		headerObj.Set(k, v)
	}
	return This.runtime.ToValue(map[string]interface{}{
		"body":   datas,
		"header": headerObj,
	})
}

func init() {
	require.RegisterNativeModule("http", func(runtime *goja.Runtime, module *goja.Object) {
		This := &httpRuntime{
			runtime: runtime,
		}

		o := module.Get("exports").(*goja.Object)
		o.Set("request", This.request)
	})
}
