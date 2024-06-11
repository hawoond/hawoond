package parser

import (
	"errors"
	"reflect"
)

// StructToMap 함수는 주어진 구조체를 맵으로 변환합니다.
func StructToMap(v interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{})
	val := reflect.ValueOf(v)

	// 입력이 구조체가 아닌 경우 에러
	if val.Kind() != reflect.Struct {
		return nil, errors.New("input is not a struct")
	}

	typ := val.Type()

	// 구조체의 각 필드를 순회하며 맵에 추가
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)
		result[field.Name] = fieldValue.Interface()
	}

	return
}

// MapToStruct 함수는 주어진 맵을 구조체로 변환합니다.
func MapToStruct(m map[string]interface{}, v interface{}) error {
	val := reflect.ValueOf(v)

	// 입력이 포인터가 아닌 경우 오류 반환
	if val.Kind() != reflect.Ptr {
		return errors.New("input is not a pointer to a struct")
	}

	val = val.Elem()

	// 입력이 구조체가 아닌 경우 오류 반환
	if val.Kind() != reflect.Struct {
		return errors.New("input is not a struct")
	}

	typ := val.Type()

	// 맵의 각 키-값 쌍을 순회하며 구조체 필드에 할당
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		if value, ok := m[field.Name]; ok {
			fieldValue := val.Field(i)
			if fieldValue.CanSet() {
				fieldValue.Set(reflect.ValueOf(value))
			}
		}
	}

	return nil
}
