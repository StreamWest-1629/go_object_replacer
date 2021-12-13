package builtin

import (
	"errors"
	"reflect"

	"github.com/streamwest-1629/go_object_replacer/replacer/conn"
)

type (
	sliceConverter struct {
		elemConv conn.Converter
		elemType reflect.Type
	}
)

// Make converter from any object (slice) to slice object.
// Returned converter can convert to only targetSlice's type.
func NewSliceConverter(replacer conn.Replacer, targetSlice reflect.Type) (converter conn.Converter, err error) {
	if conv, err := replacer.MakeConverter(targetSlice.Elem()); err != nil {
		return nil, errors.New("cannot get element converter: " + err.Error())
	} else {
		elemType := targetSlice.Elem()
		return &sliceConverter{
			elemConv: conv,
			elemType: elemType,
		}, nil
	}
}

// Convert from any object (src argument) to slice object.
// `dstPtr` argument must be slice's pointer type.
func (sc *sliceConverter) Convert(src, dstPtr interface{}) (err error) {

	// make slice instance
	dstVal := reflect.ValueOf(dstPtr)
	if dstVal.Kind() != reflect.Ptr || dstVal.Elem().Kind() != reflect.Slice {
		return errors.New("dstPtr argument is not cannot supported: it is not slice's pointer")
	}

	dstVal = dstVal.Elem()
	if dstVal.IsNil() {
		dstVal.Set(reflect.MakeSlice(reflect.SliceOf(sc.elemType), 0, 0))
	}

	// convert elements
	switch srcVal := reflect.ValueOf(src); srcVal.Kind() {
	case reflect.Array, reflect.Slice:
		for i, l := 0, srcVal.Len(); i < l; i++ {

			elem := reflect.New(sc.elemType)

			if err := sc.elemConv.Convert(srcVal.Index(i).Interface(), elem.Interface()); err != nil {
				return errors.New("cannot convert slice's element: " + err.Error())
			}
			dstVal.Set(reflect.Append(dstVal, elem.Elem()))
		}
	}

	return nil
}
