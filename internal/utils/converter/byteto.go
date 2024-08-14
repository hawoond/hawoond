package converter

type ByteTo struct{}

// 바이트를 정수(int)로 변환합니다.
func (ByteTo) ByteToInt(b byte) int {
	return int(b)
}

// 바이트를 부호 없는 정수(uint)로 변환합니다.
func (ByteTo) ByteToUint(b byte) uint {
	return uint(b)
}

// 바이트를 32비트 정수(int32)로 변환합니다.
func (ByteTo) ByteToInt32(b byte) int32 {
	return int32(b)
}

// 바이트를 32비트 부호 없는 정수(uint32)로 변환합니다.
func (ByteTo) ByteToUint32(b byte) uint32 {
	return uint32(b)
}

// 바이트를 64비트 정수(int64)로 변환합니다.
func (ByteTo) ByteToInt64(b byte) int64 {
	return int64(b)
}

// 바이트를 64비트 부호 없는 정수(uint64)로 변환합니다.
func (ByteTo) ByteToUint64(b byte) uint64 {
	return uint64(b)
}

// 바이트를 32비트 부동 소수점으로 변환합니다.
func (ByteTo) ByteToFloat32(b byte) float32 {
	return float32(b)
}

// 바이트를 64비트 부동 소수점으로 변환합니다.
func (ByteTo) ByteToFloat64(b byte) float64 {
	return float64(b)
}

// 바이트를 64비트 복소수로 변환합니다.
func (ByteTo) ByteToComplex64(b byte) complex64 {
	return complex(float32(b), 0)
}

// 바이트를 128비트 복소수로 변환합니다.
func (ByteTo) ByteToComplex128(b byte) complex128 {
	return complex(float64(b), 0)
}

// 바이트를 bool 타입으로 변환합니다.
func (ByteTo) ByteToBool(b byte) bool {
	return b != 0
}

// 바이트를 rune 타입으로 변환합니다.
func (ByteTo) ByteToRune(b byte) rune {
	return rune(b)
}
