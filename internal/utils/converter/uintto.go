package converter

type UintTo struct{}

// 부호 없는 정수(uint)를 문자열로 변환합니다.
func (UintTo) UintToString(u uint) string {
	if u == 0 {
		return "0"
	}

	var result []byte
	for u > 0 {
		digit := u % 10
		result = append([]byte{byte(digit) + '0'}, result...)
		u /= 10
	}

	return string(result)
}

// 부호 없는 정수(uint)를 32비트 부호 없는 정수(uint32)로 변환합니다.
func (UintTo) UintToUint32(u uint) uint32 {
	return uint32(u)
}

// 부호 없는 정수(uint)를 64비트 부호 없는 정수(uint64)로 변환합니다.
func (UintTo) UintToUint64(u uint) uint64 {
	return uint64(u)
}

// 부호 없는 정수(uint)를 정수(int)로 변환합니다.
// int 범위를 넘어설 경우 0을 반환합니다.
func (UintTo) UintToInt(u uint) int {
	if u > uint(maxInt) {
		return 0
	}
	return int(u)
}

// 부호 없는 정수(uint)를 32비트 정수(int32)로 변환합니다.
func (UintTo) UintToFloat(u uint) float64 {
	return float64(u)
}

// 부호 없는 정수(uint)를 32비트 부동 소수점(float32)으로 변환합니다.
func (UintTo) UintToFloat32(u uint) float32 {
	return float32(u)
}

// 부호 없는 정수(uint)를 64비트 부동 소수점(float64)으로 변환합니다.
func (UintTo) UintToFloat64(u uint) float64 {
	return float64(u)
}

// 부호 없는 정수(uint)를 64비트 복소수(complex128)로 변환합니다.
func (UintTo) UintToComplex(u uint) complex128 {
	return complex(float64(u), 0)
}
