package converter

// Uint32ToInt는 32비트 부호 없는 정수를 정수(int)로 변환합니다.
func Uint32ToInt(u uint32) int {
	return int(u)
}

// Uint32ToUint는 32비트 부호 없는 정수를 부호 없는 정수(uint)로 변환합니다.
func Uint32ToUint(u uint32) uint {
	return uint(u)
}

// Uint32ToInt32는 32비트 부호 없는 정수를 32비트 정수(int32)로 변환합니다.
func Uint32ToInt32(u uint32) int32 {
	return int32(u)
}

// Uint32ToInt64는 32비트 부호 없는 정수를 64비트 정수(int64)로 변환합니다.
func Uint32ToInt64(u uint32) int64 {
	return int64(u)
}

// Uint32ToUint64는 32비트 부호 없는 정수를 64비트 부호 없는 정수(uint64)로 변환합니다.
func Uint32ToUint64(u uint32) uint64 {
	return uint64(u)
}

// Uint32ToFloat32는 32비트 부호 없는 정수를 32비트 부동 소수점으로 변환합니다.
func Uint32ToFloat32(u uint32) float32 {
	return float32(u)
}

// Uint32ToFloat64는 32비트 부호 없는 정수를 64비트 부동 소수점으로 변환합니다.
func Uint32ToFloat64(u uint32) float64 {
	return float64(u)
}

// Uint32ToComplex64는 32비트 부호 없는 정수를 64비트 복소수로 변환합니다.
func Uint32ToComplex64(u uint32) complex64 {
	return complex(float32(u), 0)
}

// Uint32ToComplex128은 32비트 부호 없는 정수를 128비트 복소수로 변환합니다.
func Uint32ToComplex128(u uint32) complex128 {
	return complex(float64(u), 0)
}

// Uint32ToBool은 32비트 부호 없는 정수를 불리언으로 변환합니다.
func Uint32ToBool(u uint32) bool {
	return u != 0
}

// Uint32ToByte는 32비트 부호 없는 정수를 바이트로 변환합니다.
func Uint32ToByte(u uint32) byte {
	return byte(u)
}

// Uint32ToRune는 32비트 부호 없는 정수를 룬으로 변환합니다.
func Uint32ToRune(u uint32) rune {
	return rune(u)
}
