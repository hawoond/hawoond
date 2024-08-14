package checker

import (
	"reflect"
)

type Checker struct{}

// CheckConvertibleTypes는 주어진 값을 받아서 변환 가능한 타입들을 반환합니다.
func (Checker) CheckConvertibleTypes(value interface{}) []string {
	types := []string{}
	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.Int:
		types = append(types, "uint", "int32", "int64", "uint32", "uint64", "float32", "float64", "complex64", "complex128", "bool", "byte", "rune")
	case reflect.Uint:
		types = append(types, "int", "int32", "int64", "uint32", "uint64", "float32", "float64", "complex64", "complex128", "bool", "byte", "rune")
	case reflect.Int32:
		types = append(types, "int", "uint", "int64", "uint32", "uint64", "float32", "float64", "complex64", "complex128", "bool", "byte", "rune")
	case reflect.Int64:
		types = append(types, "int", "uint", "int32", "uint32", "uint64", "float32", "float64", "complex64", "complex128", "bool", "byte", "rune")
	case reflect.Uint32:
		types = append(types, "int", "uint", "int32", "int64", "uint64", "float32", "float64", "complex64", "complex128", "bool", "byte", "rune")
	case reflect.Uint64:
		types = append(types, "int", "uint", "int32", "int64", "uint32", "float32", "float64", "complex64", "complex128", "bool", "byte", "rune")
	case reflect.Float32:
		types = append(types, "int", "uint", "int32", "int64", "uint32", "uint64", "float64", "complex64", "complex128", "bool", "byte", "rune")
	case reflect.Float64:
		types = append(types, "int", "uint", "int32", "int64", "uint32", "uint64", "float32", "complex64", "complex128", "bool", "byte", "rune")
	case reflect.Complex64:
		types = append(types, "int", "uint", "int32", "int64", "uint32", "uint64", "float32", "float64", "complex128", "bool", "byte", "rune")
	case reflect.Complex128:
		types = append(types, "int", "uint", "int32", "int64", "uint32", "uint64", "float32", "float64", "complex64", "bool", "byte", "rune")
	case reflect.Bool:
		types = append(types, "int", "uint", "int32", "int64", "uint32", "uint64", "float32", "float64", "complex64", "complex128", "byte", "rune")
	case reflect.Uint8:
		types = append(types, "int", "uint", "int32", "int64", "uint32", "uint64", "float32", "float64", "complex64", "complex128", "bool", "rune")
	case reflect.Int8:
		types = append(types, "int", "uint", "int32", "int64", "uint32", "uint64", "float32", "float64", "complex64", "complex128", "bool", "byte", "rune")
	case reflect.Int16:
		types = append(types, "int", "uint", "int32", "int64", "uint32", "uint64", "float32", "float64", "complex64", "complex128", "bool", "byte", "rune")
	case reflect.Uint16:
		types = append(types, "int", "uint", "int32", "int64", "uint32", "uint64", "float32", "float64", "complex64", "complex128", "bool", "rune")
	case reflect.String:
		types = append(types, "int", "uint", "int32", "int64", "uint32", "uint64", "float32", "float64", "complex64", "complex128", "bool", "byte", "rune")
	}

	return types
}

func IsEmpty(value any) bool {
	return reflect.ValueOf(value).IsZero()
}
