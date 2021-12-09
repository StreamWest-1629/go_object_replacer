package structure

import (
	"reflect"

	"github.com/streamwest-1629/go_object_replacer/replacer/conn"
)

type (
	StructCache interface {
		MapLike(target reflect.Type) (maplike conn.MapLike, err error)
	}

	structCache struct {
		replacer conn.Replacer
		cache    map[string]*Struct
	}

	Struct struct {
		replacer        conn.Replacer
		reflects        reflect.Type
		keyMembers      map[string]Member
		embeddedMembers []Member
	}

	Member struct {
		reflects  reflect.Type
		memberAt  int
		converter conn.Converter
		requires  bool
	}
)

func NewStructCache(replacer conn.Replacer) StructCache {
	return &structCache{
		replacer: replacer,
		cache:    map[string]*Struct{},
	}
}
