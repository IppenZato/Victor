package utils

import (
	"reflect"
	"path"
)

// get pointer indirect type
func IndirectType(v reflect.Type) reflect.Type {
	switch v.Kind() {
	case reflect.Ptr:
		return IndirectType(v.Elem())
	default:
		return v
	}
}

func GetPackageName( i interface{}) string {
	return path.Base(IndirectType(reflect.TypeOf(i)).PkgPath())
}