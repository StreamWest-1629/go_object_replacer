package builtin

import (
	"errors"
	"reflect"

	"github.com/streamwest-1629/go_object_replacer/replacer/conn"
)

type (
	ptrConverter struct {
		elemConv conn.Converter
		elemType reflect.Type
	}
)

// Make converter from any object to allocated object.
// Returned converter can convert to only targetPtr's type.
func NewPtrConverter(replacer conn.Replacer, targetPtr reflect.Type) (converter conn.Converter, err error) {
	if conv, err := replacer.MakeConverter(targetPtr.Elem()); err != nil {
		return nil, errors.New("cannot get element converter: " + err.Error())
	} else {
		return &ptrConverter{
			elemConv: conv,
			elemType: targetPtr.Elem(),
		}, nil
	}
}

// Convert from any object (src argument) to allocated object.
// `dstPtr` argument must be pointer's pointer type.
//
// If dstPtr's element (this is pointer) is nil,
// this function will be allocate object.
func (s *ptrConverter) Convert(src, dstPtr interface{}) (err error) {
	dstVal := reflect.ValueOf(dstPtr)

	// make instance
	if dstVal.Kind() != reflect.Ptr || dstVal.Elem().Kind() != reflect.Ptr {
		return errors.New("dstPtr argument is not cannot supported: it is not pointer's pointer")
	} else if dstVal = dstVal.Elem(); dstVal.IsNil() {
		dstVal.Set(reflect.New(s.elemType))
	}

	// convert
	if err := s.elemConv.Convert(src, dstVal.Interface()); err != nil {
		return errors.New("cannot pointer's element: " + err.Error())
	}

	return nil
}
