package json

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

type JSONValue any

// ParseJSON: JSON 문자열을 파싱하여 JSONValue 반환
// input: JSON 문자열
// output: 파싱된 JSON 객체 또는 오류
func ParseJSON(input string) (JSONValue, error) {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return nil, fmt.Errorf("empty input")
	}

	switch input[0] {
	case '{':
		return parseObject(input)
	case '[':
		return parseArray(input)
	case '"':
		return parseString(input)
	case 't', 'f':
		return parseBoolean(input)
	case 'n':
		return parseNull(input)
	default:
		if unicode.IsDigit(rune(input[0])) || input[0] == '-' {
			return parseNumber(input)
		}
	}
	return nil, fmt.Errorf("invalid input")
}

// parseObject: JSON 객체 문자열을 파싱하여 map[string]JSONValue 반환
// input: JSON 객체 문자열
// output: 파싱된 JSON 객체(map 형태) or 에러
func parseObject(input string) (map[string]JSONValue, error) {
	if input[0] != '{' || input[len(input)-1] != '}' {
		return nil, fmt.Errorf("invalid object")
	}
	input = strings.TrimSpace(input[1 : len(input)-1])
	obj := make(map[string]JSONValue)
	for len(input) > 0 {
		colonIndex := strings.Index(input, ":")
		if colonIndex == -1 {
			return nil, fmt.Errorf("invalid object")
		}
		key := strings.TrimSpace(input[:colonIndex])
		key, err := parseString(key)
		if err != nil {
			return nil, err
		}

		input = strings.TrimSpace(input[colonIndex+1:])
		var value JSONValue
		value, input, err = parseValue(input)
		if err != nil {
			return nil, err
		}
		obj[key] = value
		if len(input) == 0 {
			break
		}
		if input[0] == ',' {
			input = strings.TrimSpace(input[1:])
		} else {
			return nil, fmt.Errorf("invalid object")
		}
	}
	return obj, nil
}

func parseArray(input string) ([]JSONValue, error) {
	if input[0] != '[' || input[len(input)-1] != ']' {
		return nil, fmt.Errorf("invalid array")
	}
	input = strings.TrimSpace(input[1 : len(input)-1])
	arr := []JSONValue{}
	for len(input) > 0 {
		var value JSONValue
		var err error
		value, input, err = parseValue(input)
		if err != nil {
			return nil, err
		}
		arr = append(arr, value)
		if len(input) == 0 {
			break
		}
		if input[0] == ',' {
			input = strings.TrimSpace(input[1:])
		} else {
			return nil, fmt.Errorf("invalid array")
		}
	}
	return arr, nil
}

func parseString(input string) (string, error) {
	if input[0] != '"' {
		return "", fmt.Errorf("invalid string")
	}
	endIndex := strings.Index(input[1:], "\"")
	if endIndex == -1 {
		return "", fmt.Errorf("invalid string")
	}
	return input[1 : endIndex+1], nil
}

func parseBoolean(input string) (bool, error) {
	if strings.HasPrefix(input, "true") {
		return true, nil
	} else if strings.HasPrefix(input, "false") {
		return false, nil
	}
	return false, fmt.Errorf("invalid boolean")
}

func parseNull(input string) (any, error) {
	if strings.HasPrefix(input, "null") {
		return nil, nil
	}
	return nil, fmt.Errorf("invalid null")
}

func parseNumber(input string) (float64, error) {
	var num float64
	n, err := fmt.Sscanf(input, "%f", &num)
	if n != 1 || err != nil {
		return 0, fmt.Errorf("invalid number")
	}
	return num, nil
}

func parseValue(input string) (JSONValue, string, error) {
	input = strings.TrimSpace(input)
	var value JSONValue
	var err error
	switch input[0] {
	case '{':
		value, err = parseObject(input)
	case '[':
		value, err = parseArray(input)
	case '"':
		value, err = parseString(input)
	case 't', 'f':
		value, err = parseBoolean(input)
	case 'n':
		value, err = parseNull(input)
	default:
		if unicode.IsDigit(rune(input[0])) || input[0] == '-' {
			value, err = parseNumber(input)
		} else {
			err = fmt.Errorf("invalid value")
		}
	}
	if err != nil {
		return nil, "", err
	}
	remaining := input[len(fmt.Sprintf("%v", value)):]
	return value, strings.TrimSpace(remaining), nil
}

// ToJSON: 아무거나 JSON 문자열로 변환
// input: 아무거나
// output: JSON 문자열 또는 오류
func ToJSON(value any) (string, error) {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Map:
		return mapToJSON(value)
	case reflect.Struct:
		return structToJSON(value)
	case reflect.Slice, reflect.Array:
		return arrayToJSON(value)
	case reflect.String:
		return fmt.Sprintf(`"%s"`, value.(string)), nil
	case reflect.Bool:
		if value.(bool) {
			return "true", nil
		}
		return "false", nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", value), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", value), nil
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", value), nil
	case reflect.Interface:
		if v.IsNil() {
			return "null", nil
		}
		return ToJSON(v.Elem().Interface())
	default:
		return "", fmt.Errorf("unsupported type: %s", v.Kind())
	}
}

func mapToJSON(input interface{}) (string, error) {
	v := reflect.ValueOf(input)
	if v.Kind() != reflect.Map {
		return "", fmt.Errorf("input is not a map")
	}
	var builder strings.Builder
	builder.WriteString("{")
	for i, key := range v.MapKeys() {
		if i > 0 {
			builder.WriteString(", ")
		}
		keyStr, err := ToJSON(key.Interface())
		if err != nil {
			return "", err
		}
		valueStr, err := ToJSON(v.MapIndex(key).Interface())
		if err != nil {
			return "", err
		}
		builder.WriteString(fmt.Sprintf(`%s: %s`, keyStr, valueStr))
	}
	builder.WriteString("}")
	return builder.String(), nil
}

func structToJSON(input interface{}) (string, error) {
	v := reflect.ValueOf(input)
	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("input is not a struct")
	}
	var builder strings.Builder
	builder.WriteString("{")
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if i > 0 {
			builder.WriteString(", ")
		}
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()
		valueStr, err := ToJSON(fieldValue)
		if err != nil {
			return "", err
		}
		builder.WriteString(fmt.Sprintf(`"%s": %s`, fieldName, valueStr))
	}
	builder.WriteString("}")
	return builder.String(), nil
}

func arrayToJSON(input interface{}) (string, error) {
	v := reflect.ValueOf(input)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return "", fmt.Errorf("input is not a slice or array")
	}
	var builder strings.Builder
	builder.WriteString("[")
	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			builder.WriteString(", ")
		}
		valueStr, err := ToJSON(v.Index(i).Interface())
		if err != nil {
			return "", err
		}
		builder.WriteString(valueStr)
	}
	builder.WriteString("]")
	return builder.String(), nil
}
