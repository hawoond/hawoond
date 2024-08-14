package time

import (
	"errors"
	"time"
)

type TimeTo struct{}

// 주어진 시간 문자열을 지정된 포맷으로 변환합니다.
// fromFormat이 제공되지 않은 경우 inputTime의 포맷을 자동으로 분석합니다.
func (TimeTo) ConvertTimeFormat(inputTime string, toFormat string, fromFormat ...string) (result string, err error) {
	var layout string

	// fromFormat이 제공된 경우 해당 포맷을 사용하여 파싱
	if len(fromFormat) > 0 {
		layout = fromFormat[0]
	} else {
		// fromFormat이 제공되지 않은 경우 inputTime의 포맷을 자동으로 분석
		layout, err = detectTimeFormat(inputTime)
		if err != nil {
			return "", errors.New("failed to detect time format, please provide fromFormat")
		}
	}

	// 입력된 시간 문자열을 분석된 또는 제공된 포맷으로 파싱
	t, err := time.Parse(layout, inputTime)
	if err != nil {
		return "", errors.New("invalid time format")
	}

	// 원하는 포맷으로 변환하여 반환
	result = t.Format(toFormat)
	return
}

// 주어진 시간 문자열의 포맷을 자동으로 분석합니다.
func detectTimeFormat(inputTime string) (string, error) {
	formats := []string{
		"2006-01-02T15:04:05",                                // ISO 8601
		"2006-01-02 15:04:05",                                // Common datetime
		"02-Jan-2006 03:04:05 PM",                            // Common datetime with PM
		"2006-01-02",                                         // Date only
		"15:04:05",                                           // Time only
		"02/01/2006",                                         // Common date format with slashes
		"02-01-2006",                                         // Common date format with dashes
		"01/02/2006",                                         // Common date format with slashes
		"01-02-2006",                                         // Common date format with dashes
		"2006/01/02",                                         // Common date format with slashes
		"2006-01-02",                                         // Common date format with dashes
		"01/02/2006",                                         // Common date format with slashes
		"01-02-2006",                                         // Common date format with dashes
		"2006/01/02 15:04:05",                                // Common datetime with slashes
		"2006-01-02 15:04:05",                                // Common datetime with dashes
		"01/02/2006 15:04:05",                                // Common datetime with slashes
		"01-02-2006 15:04:05",                                // Common datetime with dashes
		"2006/01/02 03:04:05 PM",                             // Common datetime with slashes and PM
		"2006-01-02 03:04:05 PM",                             // Common datetime with dashes and PM
		"01/02/2006 03:04:05 PM",                             // Common datetime with slashes and PM
		"01-02-2006 03:04:05 PM",                             // Common datetime with dashes and PM
		"2006/01/02 03:04:05 pm",                             // Common datetime with slashes and pm
		"2006-01-02 03:04:05 pm",                             // Common datetime with dashes and pm
		"01/02/2006 03:04:05 pm",                             // Common datetime with slashes and pm
		"01-02-2006 03:04:05 pm",                             // Common datetime with dashes and pm
		"2006/01/02 03:04:05 PM",                             // Common datetime with slashes and PM
		"2006-01-02 03:04:05 PM",                             // Common datetime with dashes and PM
		"01/02/2006 03:04:05 PM",                             // Common datetime with slashes and PM
		"01-02-2006 03:04:05 PM",                             // Common datetime with dashes and PM
		"2006/01/02 03:04:05 pm",                             // Common datetime with slashes and pm
		"2006-01-02 03:04:05 pm",                             // Common datetime with dashes and pm
		"01/02/2006 03:04:05 pm",                             // Common datetime with slashes and pm
		"01-02-2006 03:04:05 pm",                             // Common datetime with dashes and pm
		"2006/01/02 03:04:05.000",                            // Common datetime with milliseconds and slashes
		"2006-01-02 03:04:05.000",                            // Common datetime with milliseconds and dashes
		"01/02/2006 03:04:05.000",                            // Common datetime with milliseconds and slashes
		"01-02-2006 03:04:05.000",                            // Common datetime with milliseconds and dashes
		"2006/01/02 03:04:05.000 PM",                         // Common datetime with milliseconds, slashes and PM
		"2006-01-02 03:04:05.000 PM",                         // Common datetime with milliseconds, dashes and PM
		"01/02/2006 03:04:05.000 PM",                         // Common datetime with milliseconds, slashes and PM
		"01-02-2006 03:04:05.000 PM",                         // Common datetime with milliseconds, dashes and PM
		"2006/01/02 03:04:05.000 pm",                         // Common datetime with milliseconds, slashes and pm
		"2006-01-02 03:04:05.000 pm",                         // Common datetime with milliseconds, dashes and pm
		"01/02/2006 03:04:05.000 pm",                         // Common datetime with milliseconds, slashes and pm
		"01-02-2006 03:04:05.000 pm",                         // Common datetime with milliseconds, dashes and pm
		"2006/01/02 03:04:05.000 PM",                         // Common datetime with milliseconds, slashes and PM
		"2006-01-02 03:04:05.000 PM",                         // Common datetime with milliseconds, dashes and PM
		"01/02/2006 03:04:05.000 PM",                         // Common datetime with milliseconds, slashes and PM
		"01-02-2006 03:04:05.000 PM",                         // Common datetime with milliseconds, dashes and PM
		"2006/01/02 03:04:05.000 pm",                         // Common datetime with milliseconds, slashes and pm
		"2006-01-02 03:04:05.000 pm",                         // Common datetime with milliseconds, dashes and pm
		"01/02/2006 03:04:05.000 pm",                         // Common datetime with milliseconds, slashes and pm
		"01-02-2006 03:04:05.000 pm",                         // Common datetime with milliseconds, dashes and pm
		"2006/01/02 15:04:05.000",                            // Common time with milliseconds and slashes
		"2006-01-02 15:04:05.000",                            // Common time with milliseconds and dashes
		"01/02/2006 15:04:05.000",                            // Common time with milliseconds and slashes
		"01-02-2006 15:04:05.000",                            // Common time with milliseconds and dashes
		"2006/01/02 15:04:05.000 PM",                         // Common time with milliseconds, slashes and PM
		"2006-01-02 15:04:05.000 PM",                         // Common time with milliseconds, dashes and PM
		"01/02/2006 15:04:05.000 PM",                         // Common time with milliseconds, slashes and PM
		"01-02-2006 15:04:05.000 PM",                         // Common time with milliseconds, dashes and PM
		"2006/01/02 15:04:05.000 pm",                         // Common time with milliseconds, slashes and pm
		"2006-01-02 15:04:05.000 pm",                         // Common time with milliseconds, dashes and pm
		"01/02/2006 15:04:05.000 pm",                         // Common time with milliseconds, slashes and pm
		"01-02-2006 15:04:05.000 pm",                         // Common time with milliseconds, dashes and pm
		"2006-01-02T15:04:05Z",                               // ISO 8601 with Z
		"2006-01-02T15:04:05-07:00",                          // ISO 8601 with timezone
		"2006-01-02T15:04:05.000Z",                           // ISO 8601 with milliseconds and Z
		"2006-01-02T15:04:05.000-07:00",                      // ISO 8601 with milliseconds and timezone
		"2006-01-02T15:04:05.000000Z",                        // ISO 8601 with microseconds and Z
		"2006-01-02T15:04:05.000000-07:00",                   // ISO 8601 with microseconds and timezone
		"2006-01-02T15:04:05.000000000Z",                     // ISO 8601 with nanoseconds and Z
		"2006-01-02T15:04:05.000000000-07:00",                // ISO 8601 with nanoseconds and timezone
		"2006-01-02T15:04:05.000000000Z",                     // ISO 8601 with nanoseconds and Z
		"2006-01-02T15:04:05.000000000-07:00",                // ISO 8601 with nanoseconds and timezone
		"2006-01-02T15:04:05.000000000000Z",                  // ISO 8601 with picoseconds and Z
		"2006-01-02T15:04:05.000000000000-07:00",             // ISO 8601 with picoseconds and timezone
		"2006-01-02T15:04:05.000000000000Z",                  // ISO 8601 with picoseconds and Z
		"2006-01-02T15:04:05.000000000000-07:00",             // ISO 8601 with picoseconds and timezone
		"2006-01-02T15:04:05.000000000000000Z",               // ISO 8601 with femtoseconds and Z
		"2006-01-02T15:04:05.000000000000000-07:00",          // ISO 8601 with femtoseconds and timezone
		"2006-01-02T15:04:05.000000000000000Z",               // ISO 8601 with femtoseconds and Z
		"2006-01-02T15:04:05.000000000000000-07:00",          // ISO 8601 with femtoseconds and timezone
		"2006-01-02T15:04:05.000000000000000000Z",            // ISO 8601 with attoseconds and Z
		"2006-01-02T15:04:05.000000000000000000-07:00",       // ISO 8601 with attoseconds and timezone
		"2006-01-02T15:04:05.000000000000000000Z",            // ISO 8601 with attoseconds and Z
		"2006-01-02T15:04:05.000000000000000000-07:00",       // ISO 8601 with attoseconds and timezone
		"2006-01-02T15:04:05.000000000000000000000Z",         // ISO 8601 with zeptoseconds and Z
		"2006-01-02T15:04:05.000000000000000000000-07:00",    // ISO 8601 with zeptoseconds and timezone
		"2006-01-02T15:04:05.000000000000000000000Z",         // ISO 8601 with zeptoseconds and Z
		"2006-01-02T15:04:05.000000000000000000000-07:00",    // ISO 8601 with zeptoseconds and timezone
		"2006-01-02T15:04:05.000000000000000000000000Z",      // ISO 8601 with yoctoseconds and Z
		"2006-01-02T15:04:05.000000000000000000000000-07:00", // ISO 8601 with yoctoseconds and timezone
	}

	for _, format := range formats {
		if _, err := time.Parse(format, inputTime); err == nil {
			return format, nil
		}
	}

	return "", errors.New("unknown time format")
}
