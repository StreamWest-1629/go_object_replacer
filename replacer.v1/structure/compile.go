package structure

import (
	"errors"
	"reflect"
	"regexp"

	"github.com/streamwest-1629/go_object_replacer/replacer.v1"
	"github.com/streamwest-1629/go_object_replacer/util"
)

const (
	Label              = `replacer`
	LabelRegexp        = `^(?P<key>[a-zA-Z0-9][a-zA-Z0-9_-]*)(!)?$`
	LabelRequireRegexp = `[a-zA-Z0-9][a-zA-Z0-9_-]*(!)$`
	LabelEmbed         = `<-`
	LabelEmbedRequired = `<-!`
)

var (
	preCompiled        = make(map[string]*Struct)
	labelMatches       = regexp.MustCompile(LabelRegexp)
	labelRequireRegexp = regexp.MustCompile(LabelRequireRegexp)
)

// Get struct converter instance.
func StructConverter(target interface{}) (compiled *Struct, err error) {

	// get type's infomation
	return StructConverterFromReflect(reflect.TypeOf(target))
}

// Get struct converter instance from reflect.
func StructConverterFromReflect(targetTy reflect.Type) (compiled *Struct, err error) {
	ty := getStructType(targetTy)
	tyName := util.TypeFullname(ty)

	// check cache
	if converter, exist := preCompiled[tyName]; exist {
		// when package has cache
		return converter, nil
	} else if converter, err := makeStructConverter(ty); err != nil {
		// when failed to make struct converter
		return nil, err
	} else {
		// successful
		return converter, nil
	}
}

// Make struct converter's instance.
func makeStructConverter(ty reflect.Type) (converter *Struct, err error) {

	// add map cache
	tyName := util.TypeFullname(getStructType(ty))
	converter = &Struct{Members: []Member{}, Type: ty}
	preCompiled[tyName] = converter

	if err := func() error {
		for i, l := 0, ty.NumField(); i < l; i++ {

			// get member infomation
			field := ty.Field(i)
			label := field.Tag.Get(Label)

			// member's converter
			if memberConv, err := replacer.ConverterFromReflect(ty); err != nil {
				return errors.New("cannot get member object's converter: " + err.Error())

			} else if label == LabelEmbed || label == LabelEmbedRequired {
				converter.Members = append(converter.Members,
					Member{
						Converter: memberConv,
						MemberAt:  i,
						Requires:  label == LabelEmbedRequired,
						Embed:     true,
					})

			} else if matches := labelMatches.FindStringSubmatchIndex(label); matches != nil {
				key := string(labelMatches.ExpandString([]byte{}, `${key}`, label, matches))

				converter.Members = append(converter.Members,
					Member{
						Converter: memberConv,
						MemberAt:  i,
						Keyname:   key,
						Requires:  labelRequireRegexp.MatchString(label),
						Embed:     false,
					})

			}
		}
		return nil

	}(); err != nil {

		// when failed to make converter instance
		delete(preCompiled, tyName)
		return nil, err

	} else {
		return converter, nil
	}
}
