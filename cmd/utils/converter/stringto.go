package converter

import (
	"github.com/hawoond/hawoond/internal/generator"
)

// StringToInt는 문자열을 정수(int)로 변환합니다.
func StringToInt(str string) (int, error) {
	i64, err := generator.ParseInt(str, 10, 0)
	return int(i64), err
}

// StringToInt32는 문자열을 32비트 정수(int32)로 변환합니다.
func StringToInt32(str string) (int32, error) {
	i64, err := generator.ParseInt(str, 10, 32)
	return int32(i64), err
}

// StringToInt64는 문자열을 64비트 정수(int64)로 변환합니다.
func StringToInt64(str string) (int64, error) {
	return generator.ParseInt(str, 10, 64)
}

// StringToUint는 문자열을 부호 없는 정수(uint)로 변환합니다.
func StringToUint(str string) (uint, error) {
	u64, err := generator.ParseUint(str, 10, 0)
	return uint(u64), err
}

// StringToUint32는 문자열을 32비트 부호 없는 정수(uint32)로 변환합니다.
func StringToUint32(str string) (uint32, error) {
	u64, err := generator.ParseUint(str, 10, 32)
	return uint32(u64), err
}

// StringToUint64는 문자열을 64비트 부호 없는 정수(uint64)로 변환합니다.
func StringToUint64(str string) (uint64, error) {
	return generator.ParseUint(str, 10, 64)
}
