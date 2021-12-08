package conn

import "reflect"

type (
	Converter interface {
		Convert(src, dst interface{}) (err error)
	}

	MapLike interface {
		Converter
		ValueWithKey(src interface{}, key string) (value interface{}, err error)
	}

	Replacer interface {
		MakeConverterFromReflect(ty reflect.Type) (converter Converter, err error)
		MakeMapLikeFromReflect(ty reflect.Type) (maplike MapLike, err error)
	}
)

func MakeConverter(replacer Replacer, target interface{}) (converter Converter, err error) {
	return replacer.MakeConverterFromReflect(reflect.TypeOf(target))
}

func MakeMapLike(replacer Replacer, target interface{}) (maplike MapLike, err error) {
	return replacer.MakeMapLikeFromReflect(reflect.TypeOf(target))
}
