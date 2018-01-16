package lib

import (
	"os"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/nodejs/require"
)

type _os struct {
	runtime *goja.Runtime
}

func (This *_os) tempDir(call goja.FunctionCall) goja.Value {
	value := os.TempDir()
	return This.runtime.ToValue(value)
}

func (This *_os) hostname(call goja.FunctionCall) goja.Value {
	name, err := os.Hostname()
	if err != nil {
		return MakeErrorValue(This.runtime, err)
	}
	return MakeReturnValue(This.runtime, name)
}

func (This *_os) getEnv(call goja.FunctionCall) goja.Value {
	key := call.Argument(0).String()
	value := os.Getenv(key)
	return This.runtime.ToValue(value)
}

func (This *_os) remove(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	err := os.Remove(name)
	if err != nil {
		return MakeErrorValue(This.runtime, err)
	}
	return nil
}

func (This *_os) removeAll(call goja.FunctionCall) goja.Value {
	path := call.Argument(0).String()
	err := os.RemoveAll(path)
	if err != nil {
		return MakeErrorValue(This.runtime, err)
	}
	return nil
}

func (This *_os) mkdir(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	perm := call.Argument(1).ToInteger()
	err := os.Mkdir(name, os.FileMode(perm))
	if err != nil {
		return MakeErrorValue(This.runtime, err)
	}
	return nil
}

func (This *_os) mkdirAll(call goja.FunctionCall) goja.Value {
	path := call.Argument(0).String()
	perm := call.Argument(1).ToInteger()
	err := os.MkdirAll(path, os.FileMode(perm))
	if err != nil {
		return MakeErrorValue(This.runtime, err)
	}
	return nil
}

func (This *_os) getwd(call goja.FunctionCall) goja.Value {
	dir, err := os.Getwd()
	if err != nil {
		return MakeErrorValue(This.runtime, err)
	}
	return MakeReturnValue(This.runtime, dir)
}

func (This *_os) chdir(call goja.FunctionCall) goja.Value {
	dir := call.Argument(0).String()
	err := os.Chdir(dir)
	if err != nil {
		return MakeErrorValue(This.runtime, err)
	}
	return nil
}

func (This *_os) openFile(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	flag := call.Argument(1).ToInteger()
	perm := call.Argument(2).ToInteger()

	file, err := os.OpenFile(name, int(flag), os.FileMode(perm))
	if err != nil {
		return MakeErrorValue(This.runtime, err)
	}
	return MakeReturnValue(This.runtime, NewFile(This.runtime, file))
}

func (This *_os) create(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	file, err := os.Create(name)
	if err != nil {
		return MakeErrorValue(This.runtime, err)
	}
	return MakeReturnValue(This.runtime, NewFile(This.runtime, file))
}

func (This *_os) open(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	file, err := os.Open(name)
	if err != nil {
		return MakeErrorValue(This.runtime, err)
	}
	return MakeReturnValue(This.runtime, NewFile(This.runtime, file))
}

func (This *_os) stat(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	fileInfo, err := os.Stat(name)
	if err != nil {
		return MakeErrorValue(This.runtime, err)
	}
	// todo
	return MakeReturnValue(This.runtime, fileInfo)
}

func (This *_os) isExist(call goja.FunctionCall) goja.Value {
	// todo error type
	p0 := call.Argument(0).Export()
	if err, ok := p0.(error); ok {
		return This.runtime.ToValue(os.IsExist(err))
	}
	panic(This.runtime.NewTypeError("p0 is not error type:%T", p0))
}

func (This *_os) isNotExist(call goja.FunctionCall) goja.Value {
	// todo err
	p0 := call.Argument(0).Export()
	if err, ok := p0.(error); ok {
		return This.runtime.ToValue(os.IsNotExist(err))
	}
	panic(This.runtime.NewTypeError("p0 is not error type:%T", p0))
}

func init() {
	require.RegisterNativeModule("os", func(runtime *goja.Runtime, module *goja.Object) {
		This := &_os{
			runtime: runtime,
		}
		o := module.Get("exports").(*goja.Object)
		o.Set("O_CREATE", os.O_CREATE)
		o.Set("O_WRONLY", os.O_WRONLY)
		o.Set("O_RDONLY", os.O_RDONLY)
		o.Set("O_RDWR", os.O_RDWR)
		o.Set("O_APPEND", os.O_APPEND)
		o.Set("O_EXCL", os.O_EXCL)
		o.Set("O_SYNC", os.O_SYNC)
		o.Set("O_TRUNC", os.O_TRUNC)

		o.Set("create", This.create)
		o.Set("openFile", This.openFile)
		o.Set("open", This.open)
		o.Set("args", os.Args)
		o.Set("getEnv", This.getEnv)
		o.Set("tempDir", This.tempDir)
		o.Set("chdir", This.chdir)
		o.Set("getwd", This.getwd)
		o.Set("hostname", This.hostname)
		o.Set("mkdir", This.mkdir)
		o.Set("mkdirAll", This.mkdirAll)
		o.Set("remove", This.remove)
		o.Set("removeAll", This.removeAll)
		o.Set("stat", This.stat)
		o.Set("isExist", This.isExist)
		o.Set("isNotExist", This.isNotExist)
	})
}
