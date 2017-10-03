// timeutil provides a set of time utilities including comparisons,
// conversion to "DT8" int32 and "DT14" int64 formats and other
// capabilities.
package timeutil

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	DT14            = "20060102150405"
	DT6             = "200601"
	DT8             = "20060102"
	RFC3339Min      = "0000-01-01T00:00:00Z"
	RFC3339Zero     = "0001-01-01T00:00:00Z"
	RFC3339YMD      = "2006-01-02"
	ISO8601YM       = "2006-01"
	ISO8601Z2       = "2006-01-02T15:04:05-07"
	ISO8601Z4       = "2006-01-02T15:04:05-0700"
	ISO8601ZCompact = "20060102T150405Z"
)

var FormatMap = map[string]string{
	"RFC3339":    time.RFC3339,
	"RFC3339YMD": RFC3339YMD,
	"ISO8601YM":  ISO8601YM,
}

func GetFormat(formatName string) (string, error) {
	format, ok := FormatMap[strings.TrimSpace(formatName)]
	if !ok {
		return "", errors.New(fmt.Sprintf("Format Not Found: %v", format))
	}
	return format, nil
}
