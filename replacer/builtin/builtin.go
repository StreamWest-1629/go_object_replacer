package builtin

type (
	Converter struct {
		conv func(src, dstPtr interface{}) (err error)
	}

	MapLike struct {
		Converter
		valueKey func(src interface{}, key string) (value interface{}, exist bool)
	}
)

func (c *Converter) Convert(src, dstPtr interface{}) (err error) {
	return c.conv(src, dstPtr)
}

func (m *MapLike) ValueWithKey(src interface{}, key string) (value interface{}, exist bool) {
	return m.valueKey(src, key)
}

var (
	// Converter object for Integer.
	Integer = &Converter{conv: IntegerConvert}
	// Converter object for Bool.
	Bool = &Converter{conv: BoolConvert}
	// Converter object for String.
	String = &Converter{conv: StringConvert}
)
