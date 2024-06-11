package converter

// ByteToInt는 바이트를 정수(int)로 변환합니다.
func ByteToInt(b byte) int {
	return int(b)
}

// ByteToUint는 바이트를 부호 없는 정수(uint)로 변환합니다.
func ByteToUint(b byte) uint {
	return uint(b)
}

// ByteToInt32는 바이트를 32비트 정수(int32)로 변환합니다.
func ByteToInt32(b byte) int32 {
	return int32(b)
}

// ByteToUint32는 바이트를 32비트 부호 없는 정수(uint32)로 변환합니다.
func ByteToUint32(b byte) uint32 {
	return uint32(b)
}

// ByteToInt64는 바이트를 64비트 정수(int64)로 변환합니다.
func ByteToInt64(b byte) int64 {
	return int64(b)
}

// ByteToUint64는 바이트를 64비트 부호 없는 정수(uint64)로 변환합니다.
func ByteToUint64(b byte) uint64 {
	return uint64(b)
}

// ByteToFloat32는 바이트를 32비트 부동 소수점으로 변환합니다.
func ByteToFloat32(b byte) float32 {
	return float32(b)
}

// ByteToFloat64는 바이트를 64비트 부동 소수점으로 변환합니다.
func ByteToFloat64(b byte) float64 {
	return float64(b)
}

// ByteToComplex64는 바이트를 64비트 복소수로 변환합니다.
func ByteToComplex64(b byte) complex64 {
	return complex(float32(b), 0)
}

// ByteToComplex128은 바이트를 128비트 복소수로 변환합니다.
func ByteToComplex128(b byte) complex128 {
	return complex(float64(b), 0)
}

// ByteToBool은 바이트를 불리언으로 변환합니다.
func ByteToBool(b byte) bool {
	return b != 0
}

// ByteToRune는 바이트를 룬으로 변환합니다.
func ByteToRune(b byte) rune {
	return rune(b)
}
