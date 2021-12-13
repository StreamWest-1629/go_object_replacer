package builtin

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/streamwest-1629/go_object_replacer/util"
)

type StringConvertible interface {
	String() string
}

// Make converter from any object to string object.
func StringConvert(src, dst interface{}) error {

	dst_str, ok := dst.(*string)
	if !ok {
		return errors.New("dst argument type is not string's pointer")
	}

	switch src := src.(type) {
	case string:
		*dst_str = src
	case int:
		*dst_str = strconv.Itoa(src)
	case int8:
		*dst_str = strconv.FormatInt(int64(src), 10)
	case int16:
		*dst_str = strconv.FormatInt(int64(src), 10)
	case int32:
		*dst_str = strconv.FormatInt(int64(src), 10)
	case int64:
		*dst_str = strconv.FormatInt(src, 10)
	case uint:
		*dst_str = strconv.FormatUint(uint64(src), 10)
	case uint8:
		*dst_str = strconv.FormatUint(uint64(src), 10)
	case uint16:
		*dst_str = strconv.FormatUint(uint64(src), 10)
	case uint32:
		*dst_str = strconv.FormatUint(uint64(src), 10)
	case uint64:
		*dst_str = strconv.FormatUint(src, 10)
	case bool:
		*dst_str = strconv.FormatBool(src)
	case StringConvertible:
		*dst_str = src.String()
	default:
		return errors.New("it is cannot convert to string (type: " + util.TypeFullname(reflect.TypeOf(src)) + ")")
	}
	return nil
}
