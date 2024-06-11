package converter

import (
	"github.com/hawoond/hawoond/internal/generator"
)

// IntToString은 정수를 문자열로 변환합니다.
func IntToString(i int) string {
	return generator.IntToString(i)
}

// IntToUint는 정수를 부호 없는 정수로 변환합니다.
func IntToUint(i int) uint {
	return uint(i)
}

// IntToInt32는 정수를 32비트 정수로 변환합니다.
func IntToInt32(i int) int32 {
	return int32(i)
}

// IntToUint32는 정수를 32비트 부호 없는 정수로 변환합니다.
func IntToUint32(i int) uint32 {
	return uint32(i)
}

// IntToInt64는 정수를 64비트 정수로 변환합니다.
func IntToInt64(i int) int64 {
	return int64(i)
}

// IntToUint64는 정수를 64비트 부호 없는 정수로 변환합니다.
func IntToUint64(i int) uint64 {
	return uint64(i)
}

// IntToFloat32는 정수를 32비트 부동 소수점으로 변환합니다.
func IntToFloat32(i int) float32 {
	return float32(i)
}

// IntToFloat64는 정수를 64비트 부동 소수점으로 변환합니다.
func IntToFloat64(i int) float64 {
	return float64(i)
}

// IntToComplex64는 정수를 64비트 복소수로 변환합니다.
func IntToComplex64(i int) complex64 {
	return complex(float32(i), 0)
}

// IntToComplex128은 정수를 128비트 복소수로 변환합니다.
func IntToComplex128(i int) complex128 {
	return complex(float64(i), 0)
}

// IntToBool은 정수를 불리언으로 변환합니다.
func IntToBool(i int) bool {
	return i != 0
}

// IntToByte는 정수를 바이트로 변환합니다.
func IntToByte(i int) byte {
	return byte(i)
}

// IntToRune는 정수를 룬으로 변환합니다.
func IntToRune(i int) rune {
	return rune(i)
}
