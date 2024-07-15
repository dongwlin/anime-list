package util

import (
	"github.com/ncruces/go-sqlite3"
	"reflect"
	"unsafe"
)

func SetErrCode(sqlite3Err *sqlite3.Error, code sqlite3.ErrorCode) {
	codeField := reflect.ValueOf(&sqlite3Err).Elem().FieldByName("code")
	v := reflect.NewAt(codeField.Type(), unsafe.Pointer(codeField.UnsafeAddr())).Elem()
	codeValue := reflect.ValueOf(uint64(code))
	v.Set(codeValue)
}
