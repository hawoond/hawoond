package quote

import (
	"unicode/utf8"
)

const lowerhex = "0123456789abcdef"

// Quote는 주어진 문자열을 인용부호로 묶어 반환합니다.
func Quote(s string) string {
	return quoteWith(s, '"', false)
}

// quoteWith는 주어진 문자열을 지정된 인용부호로 묶어 반환합니다.
func quoteWith(s string, quote byte, ASCIIonly bool) string {
	buf := make([]byte, 0, 3*len(s)/2)
	buf = append(buf, quote)
	for width := 0; len(s) > 0; s = s[width:] {
		r, w := utf8.DecodeRuneInString(s)
		width = w
		if r == utf8.RuneError && width == 1 {
			buf = append(buf, `\x`...)
			buf = append(buf, lowerhex[s[0]>>4])
			buf = append(buf, lowerhex[s[0]&0xF])
			continue
		}
		buf = appendEscapedRune(buf, r, quote, ASCIIonly)
	}
	buf = append(buf, quote)
	return string(buf)
}

// appendEscapedRune는 주어진 문자를 이스케이프하여 버퍼에 추가합니다.
func appendEscapedRune(buf []byte, r rune, quote byte, ASCIIonly bool) []byte {
	if r == rune(quote) || r == '\\' {
		buf = append(buf, '\\')
		buf = append(buf, byte(r))
		return buf
	}
	if ASCIIonly {
		if r < utf8.RuneSelf && isPrint(r) {
			buf = append(buf, byte(r))
			return buf
		}
	}
	if isPrint(r) {
		return utf8.AppendRune(buf, r)
	}
	switch r {
	case '\a':
		buf = append(buf, `\a`...)
	case '\b':
		buf = append(buf, `\b`...)
	case '\f':
		buf = append(buf, `\f`...)
	case '\n':
		buf = append(buf, `\n`...)
	case '\r':
		buf = append(buf, `\r`...)
	case '\t':
		buf = append(buf, `\t`...)
	case '\v':
		buf = append(buf, `\v`...)
	default:
		if r < ' ' || r == 0x7f {
			buf = append(buf, `\x`...)
			buf = append(buf, lowerhex[byte(r)>>4])
			buf = append(buf, lowerhex[byte(r)&0xF])
		} else if r < 0x10000 {
			buf = append(buf, `\u`...)
			for s := 12; s >= 0; s -= 4 {
				buf = append(buf, lowerhex[r>>uint(s)&0xF])
			}
		} else {
			buf = append(buf, `\U`...)
			for s := 28; s >= 0; s -= 4 {
				buf = append(buf, lowerhex[r>>uint(s)&0xF])
			}
		}
	}
	return buf
}

// isPrint는 주어진 문자가 출력 가능한지 여부를 확인합니다.
func isPrint(r rune) bool {
	return r >= 0x20 && r <= 0x7E
}
