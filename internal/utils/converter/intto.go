package converter

import (
	"github.com/hawoond/hawoond/internal/generator"
)

type IntTo struct{}

// 정수를 문자열로 변환합니다.
func (IntTo) IntToString(i int) string {
	return generator.IntToString(i)
}

// 정수를 부호 없는 정수로 변환합니다.
func (IntTo) IntToUint(i int) uint {
	return uint(i)
}

// 정수를 32비트 정수로 변환합니다.
func (IntTo) IntToInt32(i int) int32 {
	return int32(i)
}

// 정수를 32비트 부호 없는 정수로 변환합니다.
func (IntTo) IntToUint32(i int) uint32 {
	return uint32(i)
}

// 정수를 64비트 정수로 변환합니다.
func (IntTo) IntToInt64(i int) int64 {
	return int64(i)
}

// 정수를 64비트 부호 없는 정수로 변환합니다.
func (IntTo) IntToUint64(i int) uint64 {
	return uint64(i)
}

// 정수를 32비트 부동 소수점으로 변환합니다.
func (IntTo) IntToFloat32(i int) float32 {
	return float32(i)
}

// 정수를 64비트 부동 소수점으로 변환합니다.
func (IntTo) IntToFloat64(i int) float64 {
	return float64(i)
}

// 정수를 64비트 복소수로 변환합니다.
func (IntTo) IntToComplex64(i int) complex64 {
	return complex(float32(i), 0)
}

// 정수를 128비트 복소수로 변환합니다.
func (IntTo) IntToComplex128(i int) complex128 {
	return complex(float64(i), 0)
}

// 정수를 bool 타입으로 변환합니다.
func (IntTo) IntToBool(i int) bool {
	return i != 0
}

// 정수를 바이트로 변환합니다.
func (IntTo) IntToByte(i int) byte {
	return byte(i)
}

// 정수를 rune 타입으로 변환합니다.
func (IntTo) IntToRune(i int) rune {
	return rune(i)
}
