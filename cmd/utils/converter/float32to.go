package converter

// Float32ToInt는 32비트 부동 소수점을 정수(int)로 변환합니다.
func Float32ToInt(f float32) int {
	return int(f)
}

// Float32ToUint는 32비트 부동 소수점을 부호 없는 정수(uint)로 변환합니다.
func Float32ToUint(f float32) uint {
	return uint(f)
}

// Float32ToInt32는 32비트 부동 소수점을 32비트 정수(int32)로 변환합니다.
func Float32ToInt32(f float32) int32 {
	return int32(f)
}

// Float32ToUint32는 32비트 부동 소수점을 32비트 부호 없는 정수(uint32)로 변환합니다.
func Float32ToUint32(f float32) uint32 {
	return uint32(f)
}

// Float32ToInt64는 32비트 부동 소수점을 64비트 정수(int64)로 변환합니다.
func Float32ToInt64(f float32) int64 {
	return int64(f)
}

// Float32ToUint64는 32비트 부동 소수점을 64비트 부호 없는 정수(uint64)로 변환합니다.
func Float32ToUint64(f float32) uint64 {
	return uint64(f)
}

// Float32ToFloat64는 32비트 부동 소수점을 64비트 부동 소수점으로 변환합니다.
func Float32ToFloat64(f float32) float64 {
	return float64(f)
}

// Float32ToComplex64는 32비트 부동 소수점을 64비트 복소수로 변환합니다.
func Float32ToComplex64(f float32) complex64 {
	return complex(f, 0)
}

// Float32ToComplex128은 32비트 부동 소수점을 128비트 복소수로 변환합니다.
func Float32ToComplex128(f float32) complex128 {
	return complex(float64(f), 0)
}

// Float32ToBool은 32비트 부동 소수점을 불리언으로 변환합니다.
func Float32ToBool(f float32) bool {
	return f != 0
}

// Float32ToByte는 32비트 부동 소수점을 바이트로 변환합니다.
func Float32ToByte(f float32) byte {
	return byte(f)
}

// Float32ToRune는 32비트 부동 소수점을 룬으로 변환합니다.
func Float32ToRune(f float32) rune {
	return rune(f)
}
