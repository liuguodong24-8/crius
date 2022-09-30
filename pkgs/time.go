package pkgs

import (
	"database/sql/driver"
	"encoding/json"
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
	return nt.Time.Format(nullTimeFormat), nil
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
// the propert representation of that value.
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
// the propert representation of that value.
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
