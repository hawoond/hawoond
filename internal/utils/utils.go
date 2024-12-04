package utils

import (
	"encoding/base64"
	"strings"
)

type Utils struct{}

// Base64URL 인코딩
func (u Utils) Base64UrlEncode(data []byte) string {
	return Base64UrlEncode(data)
}

// Base64URL 디코딩
func (u Utils) Base64UrlDecode(s string) ([]byte, error) {
	return Base64UrlDecode(s)
}

func Base64UrlEncode(data []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(data), "=")
}

func Base64UrlDecode(s string) ([]byte, error) {
	padding := 4 - (len(s) % 4)
	if padding != 4 {
		s += strings.Repeat("=", padding)
	}
	return base64.URLEncoding.DecodeString(s)
}
