package converter

func UintToString(u uint) string {
	return IntToString(int(u))
}

func UintToUint32(u uint) uint32 {
	return uint32(u)
}

func UintToUint64(u uint) uint64 {
	return uint64(u)
}

func UintToInt(u uint) int {
	return int(u)
}

func UintToFloat(u uint) float64 {
	return float64(u)
}

func UintToFloat32(u uint) float32 {
	return float32(u)
}

func UintToFloat64(u uint) float64 {
	return float64(u)
}

func UintToComplex(u uint) complex128 {
	return complex(float64(u), 0)
}
