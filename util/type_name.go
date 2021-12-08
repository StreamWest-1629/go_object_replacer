package util

import (
	"fmt"
	"reflect"
)

func TypeFullname(ty reflect.Type) string {
	if pkg, name := Typename(ty); len(pkg) > 0 {
		return pkg + "." + name
	} else {
		return name
	}
}

func Typename(ty reflect.Type) (pkg string, name string) {

	const (
		Ptr   = "(*%s)"
		Arr   = "[%d]%s"
		Slice = "[]%s"
		Map   = "map[%s]%s"
		Chan  = "(chan <- %s)"
		Other = "%s"
	)

	switch ty.Kind() {
	case reflect.Ptr:
		pkg, name := Typename(ty.Elem())
		return pkg, fmt.Sprintf(Ptr, name)
	case reflect.Array:
		name := TypeFullname(ty.Elem())
		return "", fmt.Sprintf(Arr, ty.Len(), name)
	case reflect.Slice:
		name := TypeFullname(ty.Elem())
		return pkg, fmt.Sprintf(Slice, name)
	case reflect.Map:
		key := TypeFullname(ty.Key())
		value := TypeFullname(ty.Elem())
		return "", fmt.Sprintf(Map, key, value)
	case reflect.Chan:
		name := TypeFullname(ty.Elem())
		return "", fmt.Sprintf(Chan, name)
	default:
		return ty.PkgPath(), ty.Name()
	}
}
