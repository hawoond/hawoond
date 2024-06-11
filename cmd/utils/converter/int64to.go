package converter

// Int64ToInt는 64비트 정수를 정수(int)로 변환합니다.
func Int64ToInt(i int64) int {
	return int(i)
}

// Int64ToUint는 64비트 정수를 부호 없는 정수(uint)로 변환합니다.
func Int64ToUint(i int64) uint {
	return uint(i)
}

// Int64ToInt32는 64비트 정수를 32비트 정수(int32)로 변환합니다.
func Int64ToInt32(i int64) int32 {
	return int32(i)
}

// Int64ToUint32는 64비트 정수를 32비트 부호 없는 정수(uint32)로 변환합니다.
func Int64ToUint32(i int64) uint32 {
	return uint32(i)
}

// Int64ToUint64는 64비트 정수를 64비트 부호 없는 정수(uint64)로 변환합니다.
func Int64ToUint64(i int64) uint64 {
	return uint64(i)
}

// Int64ToFloat32는 64비트 정수를 32비트 부동 소수점으로 변환합니다.
func Int64ToFloat32(i int64) float32 {
	return float32(i)
}

// Int64ToFloat64는 64비트 정수를 64비트 부동 소수점으로 변환합니다.
func Int64ToFloat64(i int64) float64 {
	return float64(i)
}

// Int64ToComplex64는 64비트 정수를 64비트 복소수로 변환합니다.
func Int64ToComplex64(i int64) complex64 {
	return complex(float32(i), 0)
}

// Int64ToComplex128은 64비트 정수를 128비트 복소수로 변환합니다.
func Int64ToComplex128(i int64) complex128 {
	return complex(float64(i), 0)
}

// Int64ToBool은 64비트 정수를 불리언으로 변환합니다.
func Int64ToBool(i int64) bool {
	return i != 0
}

// Int64ToByte는 64비트 정수를 바이트로 변환합니다.
func Int64ToByte(i int64) byte {
	return byte(i)
}

// Int64ToRune는 64비트 정수를 룬으로 변환합니다.
func Int64ToRune(i int64) rune {
	return rune(i)
}
