package converter

const maxInt = int(^uint(0) >> 1)

type Converter struct {
	ByteTo
	ComplexTo
	Float32To
	Float64To
	IntTo
	Int8To
	Int16To
	Int32To
	Int64To
	UintTo
	Uint8To
	Uint16To
	Uint32To
	Uint64To
	RuneTo
	StringTo
}

func init() {

}
