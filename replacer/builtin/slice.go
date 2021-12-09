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

func NewSliceConverter(replacer conn.Replacer, targetElem reflect.Type) (converter conn.Converter, err error) {
	if conv, err := replacer.MakeConverter(targetElem); err != nil {
		return nil, errors.New("cannot get element converter: " + err.Error())
	} else {
		return &sliceConverter{
			elemConv: conv,
			elemType: targetElem,
		}, nil
	}
}

func (sc *sliceConverter) Convert(src, dstPtr interface{}) (err error) {

	// make slice instance
	dstVal := reflect.ValueOf(dstPtr)
	if dstVal.Kind() != reflect.Ptr || dstVal.Elem().Kind() != reflect.Slice {
		return errors.New("dstPtr argument is not cannot supported: it is not slice's pointer")
	}

	dstVal = reflect.Indirect(dstVal)
	if dstVal.IsNil() {
		dstVal = reflect.MakeSlice(sc.elemType, 0, 0)
	}

	// convert elements
	switch srcVal := reflect.ValueOf(src); srcVal.Kind() {
	case reflect.Array, reflect.Slice:
		for i, l := 0, srcVal.Len(); i < l; i++ {

			elem := reflect.New(dstVal.Type())
			if err := sc.elemConv.Convert(srcVal.Index(i).Interface(), elem.Addr().Interface()); err != nil {
				return errors.New("cannot convert slice's element: " + err.Error())
			}
			reflect.Append(dstVal, elem)
		}
	}

	return nil
}
