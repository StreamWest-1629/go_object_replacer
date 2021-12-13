package builtin

import (
	"errors"
	"reflect"

	"github.com/streamwest-1629/go_object_replacer/replacer/conn"
)

type (
	mapMapLike struct {
		replacer conn.Replacer
	}
)

func NewMapMapLike(replacer conn.Replacer) (maplike conn.MapLike) {
	return &mapMapLike{
		replacer: replacer,
	}
}

func (m *mapMapLike) Convert(src, dstPtr interface{}) (err error) {

	maplike, err := m.replacer.MakeMapLike(reflect.TypeOf(src))
	var inserting func(key string, val interface{})

	// select inserting function
	switch dst := dstPtr.(type) {
	case *map[string]interface{}:
		inserting = func(key string, val interface{}) {
			(*dst)[key] = val
		}
	default:
		return errors.New("dstPtr argument cannot support")
	}

	if err != nil {
		return errors.New("cannot make maplike instance: " + err.Error())
	}
	keys := maplike.EnumKeys(src)

	for _, key := range keys {
		value, _ := maplike.ValueWithKey(src, key)
		inserting(key, value)
	}

	return nil
}

func (m *mapMapLike) ValueWithKey(src interface{}, key string) (value interface{}, exist bool) {
	switch dst := src.(type) {
	case map[string]interface{}:
		value, exist = dst[key]
		return
	default:
		return nil, exist
	}
}

func (m *mapMapLike) EnumKeys(src interface{}) []string {
	switch dst := src.(type) {
	case map[string]interface{}:
		keys := make([]string, 0, len(dst))
		for k := range dst {
			keys = append(keys, k)
		}

		return keys
	default:
		return nil
	}
}
