package conn

import "reflect"

type (
	Converter interface {
		// Convert from src to dstPtr's element.
		// `dstPtr` argument must be pointer.
		Convert(src, dstPtr interface{}) (err error)
	}

	MapLike interface {
		Converter
		ValueWithKey(src interface{}, key string) (value interface{}, exist bool)
		EnumKeys(src interface{}) []string
	}

	Replacer interface {
		MakeConverter(ty reflect.Type) (converter Converter, err error)
		MakeMapLike(ty reflect.Type) (maplike MapLike, err error)
	}
)

func MakeConverter(replacer Replacer, target interface{}) (converter Converter, err error) {
	return replacer.MakeConverter(reflect.TypeOf(target))
}

func MakeMapLike(replacer Replacer, target interface{}) (maplike MapLike, err error) {
	return replacer.MakeMapLike(reflect.TypeOf(target))
}
