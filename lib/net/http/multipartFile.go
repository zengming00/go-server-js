package http

import (
	"mime/multipart"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib"
)

type _multipartFile struct {
	runtime *goja.Runtime
	file    multipart.File
}

func (This *_multipartFile) read(call goja.FunctionCall) goja.Value {
	p0 := call.Argument(0).Export()
	if buf, ok := p0.([]byte); ok {
		n, err := This.file.Read(buf)
		return lib.MakeReturnValue(This.runtime, n, err)
	}
	panic(This.runtime.NewTypeError("p0 is not a []byte"))
}

func (This *_multipartFile) readAt(call goja.FunctionCall) goja.Value {
	p0 := call.Argument(0).Export()
	off := call.Argument(1).ToInteger()
	if buf, ok := p0.([]byte); ok {
		n, err := This.file.ReadAt(buf, off)
		return lib.MakeReturnValue(This.runtime, n, err)
	}
	panic(This.runtime.NewTypeError("p0 is not a []byte"))
}

func (This *_multipartFile) seek(call goja.FunctionCall) goja.Value {
	offset := call.Argument(0).ToInteger()
	whence := call.Argument(1).ToInteger()
	v, err := This.file.Seek(offset, int(whence))
	return lib.MakeReturnValue(This.runtime, v, err)
}

func (This *_multipartFile) close(call goja.FunctionCall) goja.Value {
	err := This.file.Close()
	if err != nil {
		return This.runtime.ToValue(err.Error())
	}
	return nil
}

func (This *_multipartFile) getPrototype(call goja.FunctionCall) goja.Value {
	return This.runtime.ToValue(This.file)
}

func NewMultipartFile(runtime *goja.Runtime, file multipart.File) *goja.Object {
	This := &_multipartFile{
		runtime: runtime,
		file:    file,
	}
	o := runtime.NewObject()
	o.Set("close", This.close)
	o.Set("read", This.read)
	o.Set("readAt", This.readAt)
	o.Set("seek", This.seek)
	o.Set("getPrototype", This.getPrototype)
	return o
}
