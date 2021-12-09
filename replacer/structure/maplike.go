package structure

import (
	"errors"
	"reflect"

	"github.com/streamwest-1629/go_object_replacer/replacer/conn"
	"github.com/streamwest-1629/go_object_replacer/util"
)

func (s *Struct) getValueOf(dst interface{}) (val reflect.Value) {

	for val = reflect.ValueOf(dst); val.Kind() != reflect.Struct; {
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		} else {
			panic("this is not struct or its pointer")
		}
	}

	return val
}

func (s *Struct) Convert(src, dstPtr interface{}) (err error) {

	val := s.getValueOf(dstPtr)

	mapLike, err := conn.MakeMapLike(s.replacer, src)
	if err != nil {
		return errors.New("cannot convertible source object to maplike: " + err.Error())
	}

	// check key member
	for key, member := range s.keyMembers {
		if value, exist := mapLike.ValueWithKey(src, key); exist {
			dst := val.Field(member.memberAt).Addr().Interface()

			// convert
			if err := member.converter.Convert(value, dst); err != nil {
				return errors.New("failed to convert member (type: " + util.TypeFullname(s.reflects) + "): " + err.Error())
			}

		} else if member.requires {
			return errors.New("failed to convert in required member (check the pair, key: " + key + ", type: " + util.TypeFullname(member.reflects) + ")")
		}
	}

	// check embedded member
	for i := range s.embeddedMembers {
		dst := val.Field(s.embeddedMembers[i].memberAt).Addr().Interface()

		// convert
		if err := s.embeddedMembers[i].converter.Convert(src, dst); err != nil {
			return errors.New("failed to convert embedded member (type: " + util.TypeFullname(s.reflects) + ")" + err.Error())
		}
	}

	return nil
}

func (s *Struct) ValueWithKey(src interface{}, key string) (value interface{}, exist bool) {

	val := s.getValueOf(src)

	if member, exists := s.keyMembers[key]; exists {
		return val.Field(member.memberAt).Interface(), true
	}

	for i := range s.embeddedMembers {
		if val, exists := s.embeddedMembers[i].converter.(conn.MapLike).ValueWithKey(src, key); exists {
			return val, true
		}
	}

	return nil, false
}
