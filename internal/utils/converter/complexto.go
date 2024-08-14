package converter

type ComplexTo struct{}

// 64비트 복소수를 정수(int)로 변환합니다.
func (ComplexTo) Complex64ToInt(c complex64) int {
	return int(real(c))
}

// 64비트 복소수를 부호 없는 정수(uint)로 변환합니다.
func (ComplexTo) Complex64ToUint(c complex64) uint {
	return uint(real(c))
}

// 64비트 복소수를 32비트 정수(int32)로 변환합니다.
func (ComplexTo) Complex64ToInt32(c complex64) int32 {
	return int32(real(c))
}

// 64비트 복소수를 32비트 부호 없는 정수(uint32)로 변환합니다.
func (ComplexTo) Complex64ToUint32(c complex64) uint32 {
	return uint32(real(c))
}

// 64비트 복소수를 64비트 정수(int64)로 변환합니다.
func (ComplexTo) Complex64ToInt64(c complex64) int64 {
	return int64(real(c))
}

// 64비트 복소수를 64비트 부호 없는 정수(uint64)로 변환합니다.
func (ComplexTo) Complex64ToUint64(c complex64) uint64 {
	return uint64(real(c))
}

// 64비트 복소수를 32비트 부동 소수점으로 변환합니다.
func (ComplexTo) Complex64ToFloat32(c complex64) float32 {
	return float32(real(c))
}

// 64비트 복소수를 64비트 부동 소수점으로 변환합니다.
func (ComplexTo) Complex64ToFloat64(c complex64) float64 {
	return float64(real(c))
}

// 128비트 복소수를 정수(int)로 변환합니다.
func (ComplexTo) Complex128ToInt(c complex128) int {
	return int(real(c))
}

// 128비트 복소수를 부호 없는 정수(uint)로 변환합니다.
func (ComplexTo) Complex128ToUint(c complex128) uint {
	return uint(real(c))
}

// 128비트 복소수를 32비트 정수(int32)로 변환합니다.
func (ComplexTo) Complex128ToInt32(c complex128) int32 {
	return int32(real(c))
}

// 128비트 복소수를 32비트 부호 없는 정수(uint32)로 변환합니다.
func (ComplexTo) Complex128ToUint32(c complex128) uint32 {
	return uint32(real(c))
}

// 128비트 복소수를 64비트 정수(int64)로 변환합니다.
func (ComplexTo) Complex128ToInt64(c complex128) int64 {
	return int64(real(c))
}

// 128비트 복소수를 64비트 부호 없는 정수(uint64)로 변환합니다.
func (ComplexTo) Complex128ToUint64(c complex128) uint64 {
	return uint64(real(c))
}

// 128비트 복소수를 32비트 부동 소수점으로 변환합니다.
func (ComplexTo) Complex128ToFloat32(c complex128) float32 {
	return float32(real(c))
}

// 128비트 복소수를 64비트 부동 소수점으로 변환합니다.
func (ComplexTo) Complex128ToFloat64(c complex128) float64 {
	return float64(real(c))
}
