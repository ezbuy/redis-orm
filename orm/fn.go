package orm

import (
	"encoding"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"encoding/base64"
	"reflect"
)

func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

func toTimeE(i interface{}) (time.Time, error) {
	i = indirect(i)
	switch s := i.(type) {
	case time.Time:
		return s, nil
	default:
		return time.Time{}, fmt.Errorf("Unable to Cast %#v to Time\n", i)
	}
}

func MsSQLTimeParse(s string) time.Time {
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		fmt.Println("MsSQLTimeParse failed:", err)
	}
	return t.Local()
}

func MsSQLTimeFormat(t interface{}) string {
	tm, err := toTimeE(t)
	if err != nil {
		panic(err)
	}
	return tm.UTC().Format("2006-01-02 15:04:05")
}

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

func NewStringSlice(len int, val string) []string {
	s := make([]string, len)
	for i := 0; i < len; i++ {
		s[i] = val
	}
	return s
}

func SliceJoin(objs []interface{}, sep string) string {
	s := make([]string, 0, len(objs))
	for _, obj := range objs {
		s = append(s, fmt.Sprint(obj))
	}
	return strings.Join(s, sep)
}

func ToFloat64(value interface{}) (float64, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		return strconv.ParseFloat(v, 64)
	case int:
		v, _ := value.(int)
		return float64(v), nil
	case int32:
		v, _ := value.(int32)
		return float64(v), nil
	case int64:
		v, _ := value.(int64)
		return float64(v), nil
	case float32:
		v, _ := value.(float32)
		return float64(v), nil
	case float64:
		v, _ := value.(float64)
		return v, nil
	}
	return float64(0), errors.New("unsupport type to float64")
}

func SQLWhere(conditions []string) string {
	if len(conditions) > 0 {
		return fmt.Sprintf("WHERE %s", strings.Join(conditions, " AND "))
	}
	return ""
}

func SQLOrderBy(field string, revert bool) string {
	if field != "" {
		if revert {
			return fmt.Sprintf("ORDER BY %s DESC", field)
		}
		return fmt.Sprintf("ORDER BY %s ASC", field)
	}
	return ""
}

func SQLOffsetLimit(offset, limit int) string {
	if limit <= 0 {
		return ""
	}
	if offset <= 0 {
		return fmt.Sprintf("LIMIT %d", limit)
	}
	return fmt.Sprintf("LIMIT %d, %d", offset, limit)
}

func MsSQLOffsetLimit(offset, limit int) string {
	if limit <= 0 {
		return ""
	}
	if offset < 0 {
		offset = 0
	}
	return fmt.Sprintf("OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", offset, limit)
}

func atoi(b []byte) (int, error) {
	return strconv.Atoi(string(b))
}

func parseInt(b []byte, base int, bitSize int) (int64, error) {
	return strconv.ParseInt(string(b), base, bitSize)
}

func parseUint(b []byte, base int, bitSize int) (uint64, error) {
	return strconv.ParseUint(string(b), base, bitSize)
}

func parseFloat(b []byte, bitSize int) (float64, error) {
	return strconv.ParseFloat(string(b), bitSize)
}

func StringScan(str string, v interface{}) error {
	b := []byte(str)
	switch v := v.(type) {
	case nil:
		return fmt.Errorf("StringScan(nil)")
	case *string:
		*v = str
		return nil
	case *[]byte:
		*v = b
		return nil
	case *int:
		var err error
		*v, err = atoi(b)
		return err
	case *int8:
		n, err := parseInt(b, 10, 8)
		if err != nil {
			return err
		}
		*v = int8(n)
		return nil
	case *int16:
		n, err := parseInt(b, 10, 16)
		if err != nil {
			return err
		}
		*v = int16(n)
		return nil
	case *int32:
		n, err := parseInt(b, 10, 32)
		if err != nil {
			return err
		}
		*v = int32(n)
		return nil
	case *int64:
		n, err := parseInt(b, 10, 64)
		if err != nil {
			return err
		}
		*v = n
		return nil
	case *uint:
		n, err := parseUint(b, 10, 64)
		if err != nil {
			return err
		}
		*v = uint(n)
		return nil
	case *uint8:
		n, err := parseUint(b, 10, 8)
		if err != nil {
			return err
		}
		*v = uint8(n)
		return nil
	case *uint16:
		n, err := parseUint(b, 10, 16)
		if err != nil {
			return err
		}
		*v = uint16(n)
		return nil
	case *uint32:
		n, err := parseUint(b, 10, 32)
		if err != nil {
			return err
		}
		*v = uint32(n)
		return nil
	case *uint64:
		n, err := parseUint(b, 10, 64)
		if err != nil {
			return err
		}
		*v = n
		return nil
	case *float32:
		n, err := parseFloat(b, 32)
		if err != nil {
			return err
		}
		*v = float32(n)
		return err
	case *float64:
		var err error
		*v, err = parseFloat(b, 64)
		return err
	case *bool:
		*v = len(b) == 1 && b[0] == '1'
		return nil
	case encoding.BinaryUnmarshaler:
		return v.UnmarshalBinary(b)
	default:
		return fmt.Errorf(
			"can't unmarshal %T (consider implementing BinaryUnmarshaler)", v)
	}

}

func Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

func Decode(src string) string {
	decoded, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		panic(err)
	}
	return string(decoded)
}
