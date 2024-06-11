package converter

// Int8ToInt는 8비트 정수를 정수(int)로 변환합니다.
func Int8ToInt(i int8) int {
	return int(i)
}

// Int8ToUint는 8비트 정수를 부호 없는 정수(uint)로 변환합니다.
func Int8ToUint(i int8) uint {
	return uint(i)
}

// Int8ToInt32는 8비트 정수를 32비트 정수(int32)로 변환합니다.
func Int8ToInt32(i int8) int32 {
	return int32(i)
}

// Int8ToInt64는 8비트 정수를 64비트 정수(int64)로 변환합니다.
func Int8ToInt64(i int8) int64 {
	return int64(i)
}

// Int8ToUint32는 8비트 정수를 32비트 부호 없는 정수(uint32)로 변환합니다.
func Int8ToUint32(i int8) uint32 {
	return uint32(i)
}

// Int8ToUint64는 8비트 정수를 64비트 부호 없는 정수(uint64)로 변환합니다.
func Int8ToUint64(i int8) uint64 {
	return uint64(i)
}

// Int8ToFloat32는 8비트 정수를 32비트 부동 소수점으로 변환합니다.
func Int8ToFloat32(i int8) float32 {
	return float32(i)
}

// Int8ToFloat64는 8비트 정수를 64비트 부동 소수점으로 변환합니다.
func Int8ToFloat64(i int8) float64 {
	return float64(i)
}

// Int8ToComplex64는 8비트 정수를 64비트 복소수로 변환합니다.
func Int8ToComplex64(i int8) complex64 {
	return complex(float32(i), 0)
}

// Int8ToComplex128은 8비트 정수를 128비트 복소수로 변환합니다.
func Int8ToComplex128(i int8) complex128 {
	return complex(float64(i), 0)
}

// Int8ToBool은 8비트 정수를 불리언으로 변환합니다.
func Int8ToBool(i int8) bool {
	return i != 0
}

// Int8ToByte는 8비트 정수를 바이트로 변환합니다.
func Int8ToByte(i int8) byte {
	return byte(i)
}

// Int8ToRune는 8비트 정수를 룬으로 변환합니다.
func Int8ToRune(i int8) rune {
	return rune(i)
}
