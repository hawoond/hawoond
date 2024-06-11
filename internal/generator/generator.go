package generator

import (
	"hawoond/internal/common/errorkit"
)

const intSize = 32 << (^uint(0) >> 63)
const IntSize = intSize

// Lower는 주어진 바이트를 소문자로 변환합니다.
func Lower(c byte) byte {
	return c | ('x' - 'X')
}

// ParseInt는 주어진 문자열을 정수(int64)로 변환합니다.
func ParseInt(s string, base int, bitSize int) (int64, error) {
	if s == "" {
		return 0, errorkit.SyntaxError("ParseInt", s)
	}

	neg := false
	if s[0] == '-' {
		neg = true
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	un, err := ParseUint(s, base, bitSize)
	if err != nil {
		return 0, &errorkit.NumError{Func: "ParseInt", Num: s, Err: err}
	}

	if bitSize == 0 {
		bitSize = IntSize
	}

	cutoff := uint64(1 << (bitSize - 1))
	if !neg && un >= cutoff {
		return int64(cutoff - 1), errorkit.RangeError("ParseInt", s)
	}
	if neg && un > cutoff {
		return -int64(cutoff), errorkit.RangeError("ParseInt", s)
	}
	n := int64(un)
	if neg {
		n = -n
	}
	return n, nil
}

// ParseUint는 주어진 문자열을 부호 없는 정수(uint64)로 변환합니다.
func ParseUint(s string, base int, bitSize int) (uint64, error) {
	if s == "" {
		return 0, errorkit.SyntaxError("ParseUint", s)
	}

	if base == 0 {
		if s[0] == '0' {
			switch {
			case len(s) > 2 && (s[1] == 'b' || s[1] == 'B'):
				base = 2
				s = s[2:]
			case len(s) > 2 && (s[1] == 'o' || s[1] == 'O'):
				base = 8
				s = s[2:]
			case len(s) > 2 && (s[1] == 'x' || s[1] == 'X'):
				base = 16
				s = s[2:]
			default:
				base = 8
				s = s[1:]
			}
		} else {
			base = 10
		}
	}

	if bitSize == 0 {
		bitSize = IntSize
	}

	var n uint64
	for _, c := range []byte(s) {
		var d byte
		switch {
		case '0' <= c && c <= '9':
			d = c - '0'
		case 'a' <= Lower(c) && Lower(c) <= 'z':
			d = Lower(c) - 'a' + 10
		default:
			return 0, errorkit.SyntaxError("ParseUint", s)
		}

		if d >= byte(base) {
			return 0, errorkit.SyntaxError("ParseUint", s)
		}

		n1 := n*uint64(base) + uint64(d)
		if n1 < n {
			return 0, errorkit.RangeError("ParseUint", s)
		}
		n = n1
	}

	return n, nil
}

// IntToString는 정수를 문자열로 변환합니다.
func IntToString(i int) string {
	if i == 0 {
		return "0"
	}

	var neg bool
	if i < 0 {
		neg = true
		i = -i
	}

	var b [20]byte
	bp := len(b)
	for i > 0 {
		bp--
		b[bp] = byte(i%10) + '0'
		i /= 10
	}

	if neg {
		bp--
		b[bp] = '-'
	}

	return string(b[bp:])
}
