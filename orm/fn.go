package orm

import (
	"fmt"
	"strings"
	"time"
)

func TimeToLocalTime(c time.Time) string {
	return c.Local().Format("2006-01-02 15:04:05")
}

func TimeParse(s string) time.Time {
	var err error
	var ret time.Time
	// 可能遇到多种情况
	if strings.HasSuffix(s, "Z") {
		if s != "0000-00-00T00:00:00Z" {
			ret, err = time.ParseInLocation("2006-01-02T15:04:05Z", s, time.Local)
		}
	} else {
		if s != "0000-00-00 00:00:00" {
			ret, err = time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
		}
	}
	if s != "" && err != nil {
		println("db.TimeParse error:", err.Error(), s)
	}
	return ret
}

func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func TimeParseLocalTime(s string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return t
	}
	localTime := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(),
		t.Second(), t.Nanosecond(), time.Local)
	return localTime
}

func OffsetLimit(offset, limit int) string {
	if limit <= 0 {
		return ""
	}
	if offset <= 0 {
		return fmt.Sprintf(" LIMIT %d", limit)
	}
	return fmt.Sprintf(" LIMIT %d, %d", offset, limit)
}

func NewStringSlice(len int, val string) []string {
	s := make([]string, len)
	for i := 0; i < len; i++ {
		s[i] = val
	}
	return s
}

func ToFloat64(val interface{}) float64 {
	//! TODO
	return val.(float64)
}
