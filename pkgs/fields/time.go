package fields

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cyrnicolase/nulls"
)

// NullTime model 时间格式 Y-m-d H:i:s
type NullTime nulls.Time

const nullTimeFormat = "2006-01-02 15:04:05"

//
func (nt *NullTime) String() string {
	return nt.Time.Format(nullTimeFormat)
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

// MarshalJSON marshals the underlying value to a
// proper JSON representation.
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		tune := nt.Time.Format(`"2006-01-02 15:04:05"`)
		return []byte(tune), nil
	}

	return json.Marshal(nil)
}

// UnmarshalJSON will unmarshal a JSON value into
// the property representation of that value.
func (nt *NullTime) UnmarshalJSON(text []byte) error {
	nt.Valid = false
	txt := string(text)
	if txt == "null" || txt == "" {
		return nil
	}
	t, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, txt, time.Local)
	if err == nil {
		nt.Time = t
		nt.Valid = true
	}

	return err
}

// UnmarshalText will unmarshal text value into
// the property representation of that value.
func (nt *NullTime) UnmarshalText(text []byte) error {
	return nt.UnmarshalJSON(text)
}

// StringToNullTime 字符时间转NullTime
func StringToNullTime(s string) NullTime {
	t, err := time.ParseInLocation(nullTimeFormat, s, time.Local)

	d := NullTime{Time: t}

	if nil != err {
		d.Valid = false
	} else {
		d.Valid = true
	}

	return d
}

// NewNullTime 根据传入时间转NullTime
func NewNullTime(t time.Time) NullTime {
	return StringToNullTime(t.Format(`2006-01-02 15:04:05`))
}

// DateTime date time y-m-d
type DateTime nulls.Time

// DateFormat format
const DateFormat = `2006-01-02`

// Scan datetime scan
func (d *DateTime) Scan(src interface{}) error {
	value, ok := src.(time.Time)
	if ok {
		*d = DateTime{
			Time:  value,
			Valid: value.UnixNano() > 0,
		}

		return nil
	}

	return errors.New("日期格式错误")
}

// Value implements the driver Valuer interface.
func (d DateTime) Value() (driver.Value, error) {
	if !d.Valid {
		return d.Time, nil
	}
	return d.Time.Format(DateFormat), nil
}

// String datetime string
func (d *DateTime) String() string {
	if d.Valid {
		return d.Time.Format(DateFormat)
	}

	return ""
}

// StringToDateTime 字符时间转datetime
func StringToDateTime(s string) DateTime {
	t, err := time.ParseInLocation("2006-01-02", s, time.Local)

	d := DateTime{Time: t}

	if nil != err {
		d.Valid = false
	} else {
		d.Valid = true
	}

	return d
}

// UnixToDateTime 时间戳转化为datetime
func UnixToDateTime(t int64) DateTime {
	s := time.Unix(t, 0).Format(`2006-01-02`)

	return StringToDateTime(s)
}

// ToUnix 转化为实践戳
func (d *DateTime) ToUnix() int64 {
	if nil == d || !d.Valid {
		return 0
	}

	return d.Time.Unix()
}

// MarshalJSON ...
func (d DateTime) MarshalJSON() ([]byte, error) {
	if d.Valid {
		t := d.Time.Format("2006-01-02")
		rs := []byte(fmt.Sprintf(`"%s"`, t))
		return rs, nil
	}

	return json.Marshal(nil)
}

// UnmarshalJSON unmarshal
func (d *DateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		d.Time = time.Time{}
		d.Valid = false
		return
	}

	d.Time, err = time.Parse(DateFormat, s)
	d.Valid = true

	return
}

// LocalTime 06:00:00
type LocalTime [8]byte

// Scan local time scan
func (l *LocalTime) Scan(src interface{}) error {
	var t time.Time
	switch src.(type) {
	case string:
		local := StringToLocalTime(src.(string))
		t, _ = local.Time()
	default:
		t = src.(time.Time)
	}

	copy(l[:], t.Format("15:04:05"))
	return nil
}

// MarshalText 序列化
func (l LocalTime) MarshalText() (text []byte, err error) {
	dst := make([]byte, len(l[:]))
	copy(dst, l[:])
	text = dst

	return
}

// Value local time value
func (l LocalTime) Value() (driver.Value, error) {
	t, err := l.Time()
	if err != nil {
		return nil, err
	}

	return t.Format("15:04:05"), nil
}

// String local time string
func (l LocalTime) String() string {
	return string(l[:])
}

// Time LocalTime转换为Time
func (l LocalTime) Time() (time.Time, error) {
	return time.ParseInLocation("15:04:05", l.String(), time.Local)
}

// StringToLocalTime string to localtime
func StringToLocalTime(s string) LocalTime {
	t := LocalTime{}

	copy(t[:], s)

	return t
}
