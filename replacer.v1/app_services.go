package replacer

type (
	// The interface to convert type.
	Converter interface {
		Convert(src, dst interface{}, property string) error
	}

	// The interface to get value with identity key (no convert).
	// if object has no value, return error.
	MapLike interface {
		ValueWithKey(src interface{}, key string) (result interface{}, err error)
	}
)
