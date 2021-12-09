package structure

import (
	"errors"
	"reflect"

	"github.com/streamwest-1629/go_object_replacer/replacer/conn"
	"github.com/streamwest-1629/go_object_replacer/util"
)

func (sc structCache) MapLike(target reflect.Type) (maplike conn.MapLike, err error) {

	if target.Kind() == reflect.Struct {
		return nil, errors.New("target type is not struct")
	}

	name := util.TypeFullname(target)

	if s, exists := sc.cache[name]; exists {
		return s, nil
	} else {

		// make new converter instance
		tmp := &Struct{}
		sc.cache[name] = tmp

		if err := tmp.initialize(sc.replacer, target); err != nil {
			delete(sc.cache, name)
			return nil, errors.New("failed to make new converter instance: " + err.Error())
		}

		return tmp, nil
	}
}
