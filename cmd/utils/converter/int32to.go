package converter

// Int32ToInt는 32비트 정수를 정수(int)로 변환합니다.
func Int32ToInt(i int32) int {
	return int(i)
}

// Int32ToUint는 32비트 정수를 부호 없는 정수(uint)로 변환합니다.
func Int32ToUint(i int32) uint {
	return uint(i)
}

// Int32ToInt64는 32비트 정수를 64비트 정수(int64)로 변환합니다.
func Int32ToInt64(i int32) int64 {
	return int64(i)
}

// Int32ToUint32는 32비트 정수를 32비트 부호 없는 정수(uint32)로 변환합니다.
func Int32ToUint32(i int32) uint32 {
	return uint32(i)
}

// Int32ToUint64는 32비트 정수를 64비트 부호 없는 정수(uint64)로 변환합니다.
func Int32ToUint64(i int32) uint64 {
	return uint64(i)
}

// Int32ToFloat32는 32비트 정수를 32비트 부동 소수점으로 변환합니다.
func Int32ToFloat32(i int32) float32 {
	return float32(i)
}

// Int32ToFloat64는 32비트 정수를 64비트 부동 소수점으로 변환합니다.
func Int32ToFloat64(i int32) float64 {
	return float64(i)
}

// Int32ToComplex64는 32비트 정수를 64비트 복소수로 변환합니다.
func Int32ToComplex64(i int32) complex64 {
	return complex(float32(i), 0)
}

// Int32ToComplex128은 32비트 정수를 128비트 복소수로 변환합니다.
func Int32ToComplex128(i int32) complex128 {
	return complex(float64(i), 0)
}

// Int32ToBool은 32비트 정수를 불리언으로 변환합니다.
func Int32ToBool(i int32) bool {
	return i != 0
}

// Int32ToByte는 32비트 정수를 바이트로 변환합니다.
func Int32ToByte(i int32) byte {
	return byte(i)
}

// Int32ToRune는 32비트 정수를 룬으로 변환합니다.
func Int32ToRune(i int32) rune {
	return rune(i)
}
