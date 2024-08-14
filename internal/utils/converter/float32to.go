package converter

type Float32To struct{}

// 32비트 부동 소수점을 정수(int)로 변환합니다.
func (Float32To) Float32ToInt(f float32) int {
	return int(f)
}

// 32비트 부동 소수점을 부호 없는 정수(uint)로 변환합니다.
func (Float32To) Float32ToUint(f float32) uint {
	return uint(f)
}

// 32비트 부동 소수점을 32비트 정수(int32)로 변환합니다.
func (Float32To) Float32ToInt32(f float32) int32 {
	return int32(f)
}

// 32비트 부동 소수점을 32비트 부호 없는 정수(uint32)로 변환합니다.
func (Float32To) Float32ToUint32(f float32) uint32 {
	return uint32(f)
}

// 32비트 부동 소수점을 64비트 정수(int64)로 변환합니다.
func (Float32To) Float32ToInt64(f float32) int64 {
	return int64(f)
}

// 32비트 부동 소수점을 64비트 부호 없는 정수(uint64)로 변환합니다.
func (Float32To) Float32ToUint64(f float32) uint64 {
	return uint64(f)
}

// 32비트 부동 소수점을 64비트 부동 소수점으로 변환합니다.
func (Float32To) Float32ToFloat64(f float32) float64 {
	return float64(f)
}

// 32비트 부동 소수점을 64비트 복소수로 변환합니다.
func (Float32To) Float32ToComplex64(f float32) complex64 {
	return complex(f, 0)
}

// 32비트 부동 소수점을 128비트 복소수로 변환합니다.
func (Float32To) Float32ToComplex128(f float32) complex128 {
	return complex(float64(f), 0)
}

// 32비트 부동 소수점을 bool 타입으로 변환합니다.
func (Float32To) Float32ToBool(f float32) bool {
	return f != 0
}

// 32비트 부동 소수점을 바이트로 변환합니다.
func (Float32To) Float32ToByte(f float32) byte {
	return byte(f)
}

// 32비트 부동 소수점을 rune 타입으로 변환합니다.
func (Float32To) Float32ToRune(f float32) rune {
	return rune(f)
}
