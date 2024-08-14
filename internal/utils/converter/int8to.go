package converter

type Int8To struct{}

// 8비트 정수를 정수(int)로 변환합니다.
func (Int8To) Int8ToInt(i int8) int {
	return int(i)
}

// 8비트 정수를 부호 없는 정수(uint)로 변환합니다.
func (Int8To) Int8ToUint(i int8) uint {
	return uint(i)
}

// 8비트 정수를 32비트 정수(int32)로 변환합니다.
func (Int8To) Int8ToInt32(i int8) int32 {
	return int32(i)
}

// 8비트 정수를 64비트 정수(int64)로 변환합니다.
func (Int8To) Int8ToInt64(i int8) int64 {
	return int64(i)
}

// 8비트 정수를 32비트 부호 없는 정수(uint32)로 변환합니다.
func (Int8To) Int8ToUint32(i int8) uint32 {
	return uint32(i)
}

// 8비트 정수를 64비트 부호 없는 정수(uint64)로 변환합니다.
func (Int8To) Int8ToUint64(i int8) uint64 {
	return uint64(i)
}

// 8비트 정수를 32비트 부동 소수점으로 변환합니다.
func (Int8To) Int8ToFloat32(i int8) float32 {
	return float32(i)
}

// 8비트 정수를 64비트 부동 소수점으로 변환합니다.
func (Int8To) Int8ToFloat64(i int8) float64 {
	return float64(i)
}

// 8비트 정수를 64비트 복소수로 변환합니다.
func (Int8To) Int8ToComplex64(i int8) complex64 {
	return complex(float32(i), 0)
}

// 8비트 정수를 128비트 복소수로 변환합니다.
func (Int8To) Int8ToComplex128(i int8) complex128 {
	return complex(float64(i), 0)
}

// 8비트 정수를 bool 타입으로 변환합니다.
func (Int8To) Int8ToBool(i int8) bool {
	return i != 0
}

// 8비트 정수를 바이트로 변환합니다.
func (Int8To) Int8ToByte(i int8) byte {
	return byte(i)
}

// 8비트 정수를 rune 타입으로 변환합니다.
func (Int8To) Int8ToRune(i int8) rune {
	return rune(i)
}
