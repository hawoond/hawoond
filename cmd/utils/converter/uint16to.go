package converter

// Uint16ToInt는 16비트 부호 없는 정수를 정수(int)로 변환합니다.
func Uint16ToInt(u uint16) int {
	return int(u)
}

// Uint16ToUint는 16비트 부호 없는 정수를 부호 없는 정수(uint)로 변환합니다.
func Uint16ToUint(u uint16) uint {
	return uint(u)
}

// Uint16ToInt32는 16비트 부호 없는 정수를 32비트 정수(int32)로 변환합니다.
func Uint16ToInt32(u uint16) int32 {
	return int32(u)
}

// Uint16ToInt64는 16비트 부호 없는 정수를 64비트 정수(int64)로 변환합니다.
func Uint16ToInt64(u uint16) int64 {
	return int64(u)
}

// Uint16ToUint32는 16비트 부호 없는 정수를 32비트 부호 없는 정수(uint32)로 변환합니다.
func Uint16ToUint32(u uint16) uint32 {
	return uint32(u)
}

// Uint16ToUint64는 16비트 부호 없는 정수를 64비트 부호 없는 정수(uint64)로 변환합니다.
func Uint16ToUint64(u uint16) uint64 {
	return uint64(u)
}

// Uint16ToFloat32는 16비트 부호 없는 정수를 32비트 부동 소수점으로 변환합니다.
func Uint16ToFloat32(u uint16) float32 {
	return float32(u)
}

// Uint16ToFloat64는 16비트 부호 없는 정수를 64비트 부동 소수점으로 변환합니다.
func Uint16ToFloat64(u uint16) float64 {
	return float64(u)
}

// Uint16ToComplex64는 16비트 부호 없는 정수를 64비트 복소수로 변환합니다.
func Uint16ToComplex64(u uint16) complex64 {
	return complex(float32(u), 0)
}

// Uint16ToComplex128은 16비트 부호 없는 정수를 128비트 복소수로 변환합니다.
func Uint16ToComplex128(u uint16) complex128 {
	return complex(float64(u), 0)
}

// Uint16ToBool은 16비트 부호 없는 정수를 불리언으로 변환합니다.
func Uint16ToBool(u uint16) bool {
	return u != 0
}

// Uint16ToByte는 16비트 부호 없는 정수를 바이트로 변환합니다.
func Uint16ToByte(u uint16) byte {
	return byte(u)
}

// Uint16ToRune는 16비트 부호 없는 정수를 룬으로 변환합니다.
func Uint16ToRune(u uint16) rune {
	return rune(u)
}
