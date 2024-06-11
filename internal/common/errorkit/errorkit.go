package errorkit

import "errors"

// ErrRange는 값이 범위를 벗어났음을 나타내는 오류입니다.
var ErrRange = errors.New("value out of range")

// ErrSyntax는 구문이 잘못되었음을 나타내는 오류입니다.
var ErrSyntax = errors.New("invalid syntax")

// NumError는 숫자 변환 중 발생한 오류를 나타내는 구조체입니다.
type NumError struct {
	Func string // 오류가 발생한 함수명
	Num  string // 변환하려던 숫자
	Err  error  // 발생한 오류
}

// Error는 NumError 구조체의 오류 메시지를 반환합니다.
func (e *NumError) Error() string {
	return "strconv." + e.Func + ": " + "parsing " + e.Num + ": " + e.Err.Error()
}

// SyntaxError는 구문 오류를 생성하는 함수입니다.
func SyntaxError(fn, str string) error {
	return &NumError{Func: fn, Num: str, Err: ErrSyntax}
}

// RangeError는 범위 오류를 생성하는 함수입니다.
func RangeError(fn, str string) error {
	return &NumError{Func: fn, Num: str, Err: ErrRange}
}
