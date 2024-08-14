package converter

type RuneTo struct{}

// 룬을 정수(int)로 변환합니다.
func (RuneTo) RuneToInt(r rune) int {
	return int(r)
}

// 룬을 부호 없는 정수(uint)로 변환합니다.
func (RuneTo) RuneToUint(r rune) uint {
	return uint(r)
}

// 룬을 32비트 정수(int32)로 변환합니다.
func (RuneTo) RuneToInt32(r rune) int32 {
	return int32(r)
}

// 룬을 32비트 부호 없는 정수(uint32)로 변환합니다.
func (RuneTo) RuneToUint32(r rune) uint32 {
	return uint32(r)
}

// 룬을 64비트 정수(int64)로 변환합니다.
func (RuneTo) RuneToInt64(r rune) int64 {
	return int64(r)
}

// 룬을 64비트 부호 없는 정수(uint64)로 변환합니다.
func (RuneTo) RuneToUint64(r rune) uint64 {
	return uint64(r)
}

// 룬을 32비트 부동 소수점으로 변환합니다.
func (RuneTo) RuneToFloat32(r rune) float32 {
	return float32(r)
}

// 룬을 64비트 부동 소수점으로 변환합니다.
func (RuneTo) RuneToFloat64(r rune) float64 {
	return float64(r)
}

// 룬을 64비트 복소수로 변환합니다.
func (RuneTo) RuneToComplex64(r rune) complex64 {
	return complex(float32(r), 0)
}

// 룬을 128비트 복소수로 변환합니다.
func (RuneTo) RuneToComplex128(r rune) complex128 {
	return complex(float64(r), 0)
}

// 룬을 bool 타입으로 변환합니다.
func (RuneTo) RuneToBool(r rune) bool {
	return r != 0
}

// 룬을 바이트로 변환합니다.
func (RuneTo) RuneToByte(r rune) byte {
	return byte(r)
}
