package converter

type Uint8To struct{}

// 8비트 부호 없는 정수를 정수(int)로 변환합니다.
func (Uint8To) Uint8ToInt(u uint8) int {
	return int(u)
}

// 8비트 부호 없는 정수를 부호 없는 정수(uint)로 변환합니다.
func (Uint8To) Uint8ToUint(u uint8) uint {
	return uint(u)
}

// 8비트 부호 없는 정수를 32비트 정수(int32)로 변환합니다.
func (Uint8To) Uint8ToInt32(u uint8) int32 {
	return int32(u)
}

// 8비트 부호 없는 정수를 64비트 정수(int64)로 변환합니다.
func (Uint8To) Uint8ToInt64(u uint8) int64 {
	return int64(u)
}

// 8비트 부호 없는 정수를 32비트 부호 없는 정수(uint32)로 변환합니다.
func (Uint8To) Uint8ToUint32(u uint8) uint32 {
	return uint32(u)
}

// 8비트 부호 없는 정수를 64비트 부호 없는 정수(uint64)로 변환합니다.
func (Uint8To) Uint8ToUint64(u uint8) uint64 {
	return uint64(u)
}

// 8비트 부호 없는 정수를 32비트 부동 소수점으로 변환합니다.
func (Uint8To) Uint8ToFloat32(u uint8) float32 {
	return float32(u)
}

// 8비트 부호 없는 정수를 64비트 부동 소수점으로 변환합니다.
func (Uint8To) Uint8ToFloat64(u uint8) float64 {
	return float64(u)
}

// 8비트 부호 없는 정수를 64비트 복소수로 변환합니다.
func (Uint8To) Uint8ToComplex64(u uint8) complex64 {
	return complex(float32(u), 0)
}

// 8비트 부호 없는 정수를 128비트 복소수로 변환합니다.
func (Uint8To) Uint8ToComplex128(u uint8) complex128 {
	return complex(float64(u), 0)
}

// 8비트 부호 없는 정수를 bool 타입으로 변환합니다.
func (Uint8To) Uint8ToBool(u uint8) bool {
	return u != 0
}

// 8비트 부호 없는 정수를 바이트로 변환합니다.
func (Uint8To) Uint8ToByte(u uint8) byte {
	return byte(u)
}

// 8비트 부호 없는 정수를 rune 타입으로 변환합니다.
func (Uint8To) Uint8ToRune(u uint8) rune {
	return rune(u)
}
