package converter

// Int16ToInt는 16비트 정수를 정수(int)로 변환합니다.
func Int16ToInt(i int16) int {
	return int(i)
}

// Int16ToUint는 16비트 정수를 부호 없는 정수(uint)로 변환합니다.
func Int16ToUint(i int16) uint {
	return uint(i)
}

// Int16ToInt32는 16비트 정수를 32비트 정수(int32)로 변환합니다.
func Int16ToInt32(i int16) int32 {
	return int32(i)
}

// Int16ToInt64는 16비트 정수를 64비트 정수(int64)로 변환합니다.
func Int16ToInt64(i int16) int64 {
	return int64(i)
}

// Int16ToUint32는 16비트 정수를 32비트 부호 없는 정수(uint32)로 변환합니다.
func Int16ToUint32(i int16) uint32 {
	return uint32(i)
}

// Int16ToUint64는 16비트 정수를 64비트 부호 없는 정수(uint64)로 변환합니다.
func Int16ToUint64(i int16) uint64 {
	return uint64(i)
}

// Int16ToFloat32는 16비트 정수를 32비트 부동 소수점으로 변환합니다.
func Int16ToFloat32(i int16) float32 {
	return float32(i)
}

// Int16ToFloat64는 16비트 정수를 64비트 부동 소수점으로 변환합니다.
func Int16ToFloat64(i int16) float64 {
	return float64(i)
}

// Int16ToComplex64는 16비트 정수를 64비트 복소수로 변환합니다.
func Int16ToComplex64(i int16) complex64 {
	return complex(float32(i), 0)
}

// Int16ToComplex128은 16비트 정수를 128비트 복소수로 변환합니다.
func Int16ToComplex128(i int16) complex128 {
	return complex(float64(i), 0)
}

// Int16ToBool은 16비트 정수를 불리언으로 변환합니다.
func Int16ToBool(i int16) bool {
	return i != 0
}

// Int16ToByte는 16비트 정수를 바이트로 변환합니다.
func Int16ToByte(i int16) byte {
	return byte(i)
}

// Int16ToRune는 16비트 정수를 룬으로 변환합니다.
func Int16ToRune(i int16) rune {
	return rune(i)
}
