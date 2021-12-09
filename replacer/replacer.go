package replacer

import (
	"reflect"

	"github.com/streamwest-1629/go_object_replacer/replacer/conn"
	"github.com/streamwest-1629/go_object_replacer/replacer/structure"
)

type controller struct {
	structs structure.StructCache
}

func (c controller) MakeConverter(ty reflect.Type) (converter conn.Converter, err error) {

}

func (c controller) MakeMapLike(ty reflect.Type) (maplike conn.MapLike, err error) {

}
