package converter

type Uint64To struct{}

// 64비트 부호 없는 정수를 정수(int)로 변환합니다.
func (Uint64To) Uint64ToInt(u uint64) int {
	return int(u)
}

// 64비트 부호 없는 정수를 부호 없는 정수(uint)로 변환합니다.
func (Uint64To) Uint64ToUint(u uint64) uint {
	return uint(u)
}

// 64비트 부호 없는 정수를 32비트 정수(int32)로 변환합니다.
func (Uint64To) Uint64ToInt32(u uint64) int32 {
	return int32(u)
}

// 64비트 부호 없는 정수를 64비트 정수(int64)로 변환합니다.
func (Uint64To) Uint64ToInt64(u uint64) int64 {
	return int64(u)
}

// 64비트 부호 없는 정수를 32비트 부호 없는 정수(uint32)로 변환합니다.
func (Uint64To) Uint64ToUint32(u uint64) uint32 {
	return uint32(u)
}

// 64비트 부호 없는 정수를 32비트 부동 소수점으로 변환합니다.
func (Uint64To) Uint64ToFloat32(u uint64) float32 {
	return float32(u)
}

// 64비트 부호 없는 정수를 64비트 부동 소수점으로 변환합니다.
func (Uint64To) Uint64ToFloat64(u uint64) float64 {
	return float64(u)
}

// 64비트 부호 없는 정수를 64비트 복소수로 변환합니다.
func (Uint64To) Uint64ToComplex64(u uint64) complex64 {
	return complex(float32(u), 0)
}

// 64비트 부호 없는 정수를 128비트 복소수로 변환합니다.
func (Uint64To) Uint64ToComplex128(u uint64) complex128 {
	return complex(float64(u), 0)
}

// 64비트 부호 없는 정수를 bool 타입으로 변환합니다.
func (Uint64To) Uint64ToBool(u uint64) bool {
	return u != 0
}

// 64비트 부호 없는 정수를 바이트로 변환합니다.
func (Uint64To) Uint64ToByte(u uint64) byte {
	return byte(u)
}

// 64비트 부호 없는 정수를 rune 타입으로 변환합니다.
func (Uint64To) Uint64ToRune(u uint64) rune {
	return rune(u)
}
