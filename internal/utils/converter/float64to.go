package converter

type Float64To struct{}

// 64비트 부동 소수점을 정수(int)로 변환합니다.
func (Float64To) Float64ToInt(f float64) int {
	return int(f)
}

// 64비트 부동 소수점을 부호 없는 정수(uint)로 변환합니다.
func (Float64To) Float64ToUint(f float64) uint {
	return uint(f)
}

// 64비트 부동 소수점을 32비트 정수(int32)로 변환합니다.
func (Float64To) Float64ToInt32(f float64) int32 {
	return int32(f)
}

// 64비트 부동 소수점을 32비트 부호 없는 정수(uint32)로 변환합니다.
func (Float64To) Float64ToUint32(f float64) uint32 {
	return uint32(f)
}

// 64비트 부동 소수점을 64비트 정수(int64)로 변환합니다.
func (Float64To) Float64ToInt64(f float64) int64 {
	return int64(f)
}

// 64비트 부동 소수점을 64비트 부호 없는 정수(uint64)로 변환합니다.
func (Float64To) Float64ToUint64(f float64) uint64 {
	return uint64(f)
}

// 64비트 부동 소수점을 32비트 부동 소수점으로 변환합니다.
func (Float64To) Float64ToFloat32(f float64) float32 {
	return float32(f)
}

// 64비트 부동 소수점을 64비트 복소수로 변환합니다.
func (Float64To) Float64ToComplex64(f float64) complex64 {
	return complex(float32(f), 0)
}

// 64비트 부동 소수점을 128비트 복소수로 변환합니다.
func (Float64To) Float64ToComplex128(f float64) complex128 {
	return complex(f, 0)
}

// 64비트 부동 소수점을 bool 타입으로 변환합니다.
func (Float64To) Float64ToBool(f float64) bool {
	return f != 0
}

// 64비트 부동 소수점을 바이트로 변환합니다.
func (Float64To) Float64ToByte(f float64) byte {
	return byte(f)
}

// 64비트 부동 소수점을 rune 타입으로 변환합니다.
func (Float64To) Float64ToRune(f float64) rune {
	return rune(f)
}
