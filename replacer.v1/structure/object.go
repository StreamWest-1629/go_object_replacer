package structure

import (
	"reflect"

	"github.com/streamwest-1629/go_object_replacer/replacer.v1"
)

type (

	// Defines to convert bitween structure object.
	//
	// Source map object's key value are defined by member's label: `map-to:"keyname"`.
	// Keyname is valid with regular expression [a-zA-Z0-9][a-zA-Z0-9_-]*.
	// If a member is require value, append '!' to keyname. For example, `map-to:"dirname!"`
	// If key is `<-`, it's embedded object.
	Struct struct {
		Members []Member
		Type    reflect.Type
	}

	// Defines member object
	Member struct {
		// Convert function, use for membervalue set.
		replacer.Converter
		// The source map object side's key name.
		Keyname string
		// Thedestination structure object side's member number.
		MemberAt int
		// Shows whether converter occers error when map.
		Requires bool
		// Embedded value, uses same map as given source value.
		Embed bool
	}
)
