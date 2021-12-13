package builtin

import (
	"errors"
	"reflect"
)

// Convert from any object to integer.
// `dstPtr` argument must be integer-type's pointer type.
func IntegerConvert(src, dstPtr interface{}) error {

	srcVal, dstVal := reflect.ValueOf(src), reflect.ValueOf(dstPtr)

	if dstVal.Kind() != reflect.Ptr {
		return errors.New("dst argument type is not integer's pointer")
	} else {
		switch dstVal.Elem().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			break
		default:
			return errors.New("dst argument type is not integer's pointer")
		}
	}

	if !srcVal.CanConvert(dstVal.Elem().Type()) {
		return errors.New("src argument cannot convert to dst type")
	}

	dstVal.Elem().Set(srcVal.Convert(dstVal.Elem().Type()))
	return nil
}

// Convert from any object to bool.
// `dstPtr` argument must be bool's pointer type.
func BoolConvert(src, dstPtr interface{}) error {

	srcVal, srcOk := src.(bool)
	dstVal, dstOk := dstPtr.(*bool)

	if srcOk && dstOk {
		*dstVal = srcVal
		return nil
	} else if !dstOk {
		return errors.New("dst argument type is not bool's pointer")
	} else {
		return errors.New("src argument cannot convert to bool")
	}

}
