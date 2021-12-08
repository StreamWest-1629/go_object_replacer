package structure

import "reflect"

func getStructType(ty reflect.Type) reflect.Type {
	switch ty.Kind() {
	case reflect.Struct:
		return ty
	case reflect.Ptr:
		return getStructType(ty.Elem())
	default:
		panic("the type used in compile is invalid, allow struct or struct's pointer")
	}
}

func 
