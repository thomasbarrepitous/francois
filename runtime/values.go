package runtime

type ValueType string

const (
	NumericType ValueType = "Numeric"
	NullType    ValueType = "Null"
	BooleanType ValueType = "Boolean"
	ObjectType  ValueType = "Object"
)

type RuntimeValue interface {
	Type() ValueType
}

type NumericValue struct {
	Value float64
}

func (v NumericValue) Type() ValueType {
	return NumericType
}

func MakeNumericValue(value float64) NumericValue {
	return NumericValue{Value: value}
}

type NullValue struct {
	Value string
}

func (v NullValue) Type() ValueType {
	return NullType
}

func MakeNullValue() NullValue {
	return NullValue{Value: "rien"}
}

type BooleanValue struct {
	Value bool
}

func (v BooleanValue) Type() ValueType {
	return BooleanType
}

func MakeBooleanValue(value bool) BooleanValue {
	return BooleanValue{Value: value}
}

type ObjectValue struct {
	Properties map[string]RuntimeValue
}

func (v ObjectValue) Type() ValueType {
	return ObjectType
}
