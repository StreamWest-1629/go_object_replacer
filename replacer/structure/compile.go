package structure

import (
	"errors"
	"reflect"
	"regexp"

	"github.com/streamwest-1629/go_object_replacer/replacer/conn"
)

const (
	labelName          = `replacer`
	labelRegexp        = `^(?P<key>[a-zA-Z0-9][a-zA-Z0-9_-]*)(!)?$`
	labelRequireRegexp = `[a-zA-Z0-9][a-zA-Z0-9_-]*(!)$`
	labelEmbedded      = `<-`
)

var (
	labelMatches         = regexp.MustCompile(labelRegexp)
	labelRequiredMatches = regexp.MustCompile(labelRequireRegexp)
)

func (s *Struct) initialize(replacer conn.Replacer, target reflect.Type) error {

	// initialize myown member
	*s = Struct{
		replacer:        replacer,
		reflects:        s.reflects,
		keyMembers:      map[string]Member{},
		embeddedMembers: []Member{},
	}

	// make member converter
	for i, l := 0, target.NumField(); i < l; i++ {

		field := target.Field(i)
		label := field.Tag.Get(labelName)

		if label == labelEmbedded {

			// get maplike instance
			maplike, err := replacer.MakeMapLike(field.Type)

			if err != nil {
				return errors.New("cannot make maplike instance: " + err.Error())
			}

			// make embedded member
			s.embeddedMembers = append(s.embeddedMembers,
				Member{
					reflects:  field.Type,
					memberAt:  i,
					converter: maplike,
					requires:  true,
				})

		} else if matches := labelMatches.FindStringSubmatchIndex(label); matches != nil {

			// check keyname
			keyname := string(labelMatches.ExpandString([]byte{}, `${key}`, label, matches))

			// get converter instance
			converter, err := replacer.MakeConverter(field.Type)

			if err != nil {
				return errors.New("cannot make maplike instance: " + err.Error())
			}

			// make converter member
			s.keyMembers[keyname] = Member{
				reflects:  field.Type,
				memberAt:  i,
				converter: converter,
				requires:  labelRequiredMatches.MatchString(label),
			}
		}
	}

	return nil
}
