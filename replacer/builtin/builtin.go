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
	Integer = &Converter{conv: IntegerConvert}
	Bool    = &Converter{conv: BoolConvert}
	String  = &Converter{conv: StringConvert}
)
