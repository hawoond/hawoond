package converter

// RuneToInt는 룬을 정수(int)로 변환합니다.
func RuneToInt(r rune) int {
	return int(r)
}

// RuneToUint는 룬을 부호 없는 정수(uint)로 변환합니다.
func RuneToUint(r rune) uint {
	return uint(r)
}

// RuneToInt32는 룬을 32비트 정수(int32)로 변환합니다.
func RuneToInt32(r rune) int32 {
	return int32(r)
}

// RuneToUint32는 룬을 32비트 부호 없는 정수(uint32)로 변환합니다.
func RuneToUint32(r rune) uint32 {
	return uint32(r)
}

// RuneToInt64는 룬을 64비트 정수(int64)로 변환합니다.
func RuneToInt64(r rune) int64 {
	return int64(r)
}

// RuneToUint64는 룬을 64비트 부호 없는 정수(uint64)로 변환합니다.
func RuneToUint64(r rune) uint64 {
	return uint64(r)
}

// RuneToFloat32는 룬을 32비트 부동 소수점으로 변환합니다.
func RuneToFloat32(r rune) float32 {
	return float32(r)
}

// RuneToFloat64는 룬을 64비트 부동 소수점으로 변환합니다.
func RuneToFloat64(r rune) float64 {
	return float64(r)
}

// RuneToComplex64는 룬을 64비트 복소수로 변환합니다.
func RuneToComplex64(r rune) complex64 {
	return complex(float32(r), 0)
}

// RuneToComplex128은 룬을 128비트 복소수로 변환합니다.
func RuneToComplex128(r rune) complex128 {
	return complex(float64(r), 0)
}

// RuneToBool은 룬을 불리언으로 변환합니다.
func RuneToBool(r rune) bool {
	return r != 0
}

// RuneToByte는 룬을 바이트로 변환합니다.
func RuneToByte(r rune) byte {
	return byte(r)
}
