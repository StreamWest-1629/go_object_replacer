package replacer

import (
	"errors"
	"reflect"

	"github.com/streamwest-1629/go_object_replacer/replacer/builtin"
	"github.com/streamwest-1629/go_object_replacer/replacer/conn"
	"github.com/streamwest-1629/go_object_replacer/replacer/structure"
	"github.com/streamwest-1629/go_object_replacer/util"
)

type (
	controller struct {
		structs structure.StructCache
	}
)

var (
	ctrl = newController()
)

func newController() (ctrl *controller) {
	ctrl = &controller{}
	(*ctrl) = controller{
		structs: structure.NewStructCache(ctrl),
	}
	return ctrl
}

func MakeConverter(target interface{}) (converter conn.Converter, err error) {
	return ctrl.MakeConverter(reflect.TypeOf(target))
}

func (c controller) MakeConverter(ty reflect.Type) (converter conn.Converter, err error) {
	switch ty.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return builtin.Integer, nil
	case reflect.Bool:
		return builtin.Bool, nil
	case reflect.String:
		return builtin.String, nil
	case reflect.Ptr:
		if build, err := builtin.NewPtrConverter(c, ty); err != nil {
			return nil, err
		} else {
			return build, nil
		}
	case reflect.Slice:
		if build, err := builtin.NewSliceConverter(c, ty); err != nil {
			return nil, err
		} else {
			return build, nil
		}
	case reflect.Struct:
		if build, err := c.structs.MapLike(ty); err != nil {
			return nil, err
		} else {
			return build, nil
		}
	case reflect.Map:
		return builtin.NewMapMapLike(c), nil
	}

	return nil, errors.New("this type is not supported(type: " + util.TypeFullname(ty) + ")")
}

func (c controller) MakeMapLike(ty reflect.Type) (maplike conn.MapLike, err error) {
	switch ty.Kind() {
	case reflect.Struct:
		if build, err := c.structs.MapLike(ty); err != nil {
			return nil, err
		} else {
			return build, nil
		}
	case reflect.Map:
		return builtin.NewMapMapLike(c), nil
	}
	return nil, errors.New("this type is not supported(type: " + util.TypeFullname(ty) + ")")
}
