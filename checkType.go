package main

import (
	"reflect"
)

var validType = []reflect.Kind{
	reflect.Bool,
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Float32,
	reflect.Float64,
	reflect.Complex64,
	reflect.Complex128,
	reflect.String,
}

func IsValidType(val interface{}) bool {
	k := reflect.TypeOf(val).Kind()
	for _, v := range validType {
		if v == k {
			return true
		}
	}
	return false
}
