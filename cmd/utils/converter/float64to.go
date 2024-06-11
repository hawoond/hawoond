package converter

// Float64ToInt는 64비트 부동 소수점을 정수(int)로 변환합니다.
func Float64ToInt(f float64) int {
	return int(f)
}

// Float64ToUint는 64비트 부동 소수점을 부호 없는 정수(uint)로 변환합니다.
func Float64ToUint(f float64) uint {
	return uint(f)
}

// Float64ToInt32는 64비트 부동 소수점을 32비트 정수(int32)로 변환합니다.
func Float64ToInt32(f float64) int32 {
	return int32(f)
}

// Float64ToUint32는 64비트 부동 소수점을 32비트 부호 없는 정수(uint32)로 변환합니다.
func Float64ToUint32(f float64) uint32 {
	return uint32(f)
}

// Float64ToInt64는 64비트 부동 소수점을 64비트 정수(int64)로 변환합니다.
func Float64ToInt64(f float64) int64 {
	return int64(f)
}

// Float64ToUint64는 64비트 부동 소수점을 64비트 부호 없는 정수(uint64)로 변환합니다.
func Float64ToUint64(f float64) uint64 {
	return uint64(f)
}

// Float64ToFloat32는 64비트 부동 소수점을 32비트 부동 소수점으로 변환합니다.
func Float64ToFloat32(f float64) float32 {
	return float32(f)
}

// Float64ToComplex64는 64비트 부동 소수점을 64비트 복소수로 변환합니다.
func Float64ToComplex64(f float64) complex64 {
	return complex(float32(f), 0)
}

// Float64ToComplex128은 64비트 부동 소수점을 128비트 복소수로 변환합니다.
func Float64ToComplex128(f float64) complex128 {
	return complex(f, 0)
}

// Float64ToBool은 64비트 부동 소수점을 불리언으로 변환합니다.
func Float64ToBool(f float64) bool {
	return f != 0
}

// Float64ToByte는 64비트 부동 소수점을 바이트로 변환합니다.
func Float64ToByte(f float64) byte {
	return byte(f)
}

// Float64ToRune는 64비트 부동 소수점을 룬으로 변환합니다.
func Float64ToRune(f float64) rune {
	return rune(f)
}
