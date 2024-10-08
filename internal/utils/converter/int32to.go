package converter

type Int32To struct{}

// 32비트 정수를 정수(int)로 변환합니다.
func (Int32To) Int32ToInt(i int32) int {
	return int(i)
}

// 32비트 정수를 부호 없는 정수(uint)로 변환합니다.
func (Int32To) Int32ToUint(i int32) uint {
	return uint(i)
}

// 32비트 정수를 64비트 정수(int64)로 변환합니다.
func (Int32To) Int32ToInt64(i int32) int64 {
	return int64(i)
}

// 32비트 정수를 32비트 부호 없는 정수(uint32)로 변환합니다.
func (Int32To) Int32ToUint32(i int32) uint32 {
	return uint32(i)
}

// 32비트 정수를 64비트 부호 없는 정수(uint64)로 변환합니다.
func (Int32To) Int32ToUint64(i int32) uint64 {
	return uint64(i)
}

// 32비트 정수를 32비트 부동 소수점으로 변환합니다.
func (Int32To) Int32ToFloat32(i int32) float32 {
	return float32(i)
}

// 32비트 정수를 64비트 부동 소수점으로 변환합니다.
func (Int32To) Int32ToFloat64(i int32) float64 {
	return float64(i)
}

// 32비트 정수를 64비트 복소수로 변환합니다.
func (Int32To) Int32ToComplex64(i int32) complex64 {
	return complex(float32(i), 0)
}

// 32비트 정수를 128비트 복소수로 변환합니다.
func (Int32To) Int32ToComplex128(i int32) complex128 {
	return complex(float64(i), 0)
}

// 32비트 정수를 bool 타입으로 변환합니다.
func (Int32To) Int32ToBool(i int32) bool {
	return i != 0
}

// 32비트 정수를 바이트로 변환합니다.
func (Int32To) Int32ToByte(i int32) byte {
	return byte(i)
}

// 32비트 정수를 rune 타입으로 변환합니다.
func (Int32To) Int32ToRune(i int32) rune {
	return rune(i)
}
