package db

import (
	"fmt"
	"strings"
	"time"
	"database/sql/driver"
	"strconv"
)

const (
	formatTime         = "15:04:05"
	formatDate         = "2006-01-02"
	formatDateTime     = "2006-01-02 15:04:05"
	formatDateTimeFull = "2006-01-02 15:04:05.000Z07:00"
	formatDateTimeFul1 = "2006-01-02T15:04:05.999999999Z07:00"
	formatDateTimeFul2 = "\"2006-01-02T15:04:05.999999999Z07:00\""
	formatDateTimeFul3 = "2006-01-02T15:04:05"

	formatJsonTime     = "\"15:04:05\""
	formatJsonDate     = "\"2006-01-02\""
	formatJsonDateTime = "\"2006-01-02 15:04:05\""
	formatJsonFull     = "\"2006-01-02T15:04:05.000Z07:00\""
	formatJsonFull1    = "\"2018-03-29T13:59:49.478756+00:00\""
)

const (
	TypeTimeField     = 16	
	TypeDateField     = 32
	TypeDateTimeField = 64
)

// CRM date field definition
type TDate struct {
	time.Time
}

func (e *TDate) FieldType() int {
	return TypeDateField
}
func (e *TDate) RawValue() interface{} {
	//return e.ValueTime()
	return e.Format(formatDate)
}
func (e *TDate) Set(d time.Time) {
	*e = TDate{d}
}
func (e *TDate) SetRaw(value interface{}) error {
	switch d := value.(type) {
	case time.Time:
		e.Set(d)
	case string:
		if strings.Compare(string(d), "null") == 0 {
			e.Set(time.Time{})
			return nil
		}

		v, err := time.Parse(formatDate, string(d))
		if err == nil{
			e.Set(v)
		}
		return err
	case nil:
		e.Set(time.Time{})
	default:
		return fmt.Errorf("<DateField.SetRaw> unknown value `%s`", value)
	}
	return nil
}
func (e *TDate) String() string {
	return e.ValueTime().Format(formatDate)// String()
}
func (e TDate) ValueTime() time.Time {
	return time.Time(e.Time)
}

// Value implements the driver Valuer interface.
func (e TDate) Value() (val driver.Value, err error) {
	return []byte(e.String()), nil
}

// MarshalJSON implements the json.Marshaler interface.
func (t TDate) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return []byte(t.Format(formatJsonDate)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *TDate) UnmarshalJSON(data []byte) (err error) {
	if strings.Compare(string(data), "null") == 0 {
		*t = TDate{}
		return
	}
	tt, err := time.Parse(formatJsonDate, string(data))
	*t = TDate{tt}
	return
}

// CRM time field definition
type TTime struct {
	time.Time
}

func (e *TTime) FieldType() int {
	return TypeTimeField
}

func (e *TTime) RawValue() interface{} {
	//return e.ValueTime()
	return e.Format(formatTime)
}
func (e *TTime) Set(d time.Time) {
	*e = TTime{d}
}
func (e *TTime) SetRaw(value interface{}) error {
	switch d := value.(type) {
	case time.Time:
		e.Set(d)
	case string:
		if strings.Compare(string(d), "null") == 0 {
			e.Set(time.Time{})
			return nil
		}

		v, err := time.Parse(formatTime, string(d))
		if err == nil {
			e.Set(v)
		}
		return err
	case nil:
		e.Set(time.Time{})
	default:
		return fmt.Errorf("<TimeField.SetRaw> unknown value `%s`", value)
	}
	return nil
}
func (e *TTime) String() string {
	return e.ValueTime().String()
}
func (e TTime) ValueTime() time.Time {
	return time.Time(e.Time)
}

// Value implements the driver Valuer interface.
func (e TTime) Value() (val driver.Value, err error) {
	return []byte(e.String()), nil
}

// MarshalJSON implements the json.Marshaler interface.
func (t TTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return []byte(t.Format(formatJsonTime)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *TTime) UnmarshalJSON(data []byte) (err error) {
	if strings.Compare(string(data), "null") == 0 {
		*t = TTime{}
		return
	}
	tt, err := time.Parse(formatJsonTime, string(data))
	*t = TTime{tt}
	return
}

// CRM datetime field definition
type TDateTime struct {
	time.Time
	//IsShortFormat bool `json:"-"`
}

func (e *TDateTime) FieldType() int {
	return TypeDateTimeField
}

func (e *TDateTime) RawValue() interface{} {
	//	return e.ValueTime()
	return e.Format(formatDateTime)
}

func (e *TDateTime) Set(d time.Time) {
	*e = TDateTime{d}
	//*e = TDateTime{d, false}
}
func (e *TDateTime) SetRaw(value interface{}) error {
	switch d := value.(type) {
	case int64:
		e.Set(time.Unix(d, 0))
	case time.Time:
		e.Set(d)
	case string:
		//IsShortFormat := false

		if strings.Compare(string(d), "null") == 0 {
			e.Set(time.Time{})
			return nil
		}

		v, err := time.Parse(formatDateTimeFull, string(d))
		if err != nil{
			v, err = time.Parse(formatDateTime, string(d))
			if err != nil{
				v, err = time.Parse(formatDate, string(d))
			}
			//IsShortFormat = true
		}
		if  err == nil{
			e.Set(v)
			//e.IsShortFormat = IsShortFormat
		}
		return err
	case nil:
		e.Set(time.Time{})
	default:
		return fmt.Errorf("<DateTimeField.SetRaw> unknown value `%s`", value)
	}
	return nil
}
func (e *TDateTime) String() string {
	//if e.IsShortFormat {
	//	return e.Format(formatJsonDateTime)
	//
	//} else {
		return e.Format(formatJsonFull)
	//}
}

func (e TDateTime) ValueTime() time.Time {
	return time.Time(e.Time)
}

// Value implements the driver Valuer interface.
func (e TDateTime) Value() (val driver.Value, err error) {
	return []byte(e.String()), nil
}

// MarshalJSON implements the json.Marshaler interface.
func (t TDateTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	//if t.Unix() == -2209075200 {
	//	return []byte("null"), nil
	//}
	//if t.IsShortFormat {
	//	return []byte(t.Format(formatJsonDateTime)), nil
	//
	//} else {
		return []byte(t.Format(formatJsonFull)), nil
	//}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The datetime is expected to be a quoted string in RFC 3339 format.
func (t *TDateTime) UnmarshalJSON(data []byte) (err error) {
	if strings.Compare(string(data), "null") == 0 {
		*t = TDateTime{}
		return
	}

	//log.Debug(string(data))

	//var IsShortFormat bool

	tt, err := time.Parse(time.RFC3339, string(data))

	if err != nil{
		tt, err = time.Parse(time.RFC3339Nano, string(data))
	}

	if err != nil{
		tt, err = time.Parse(formatJsonFull, string(data))
	}

	if err != nil{
		tt, err = time.Parse(formatJsonDateTime, string(data))
		//IsShortFormat = true;
	}

	if err != nil{
		tt, err = time.Parse(formatJsonFull1, string(data))
	}

	if err != nil{
		tt, err = time.Parse(formatDateTimeFull, string(data))
	}

	if err != nil{
		tt, err = time.Parse(formatDateTimeFul1, string(data))
	}

	if err != nil{
		tt, err = time.Parse(formatDateTimeFul2, string(data))
	}

	if err != nil{
		tt, err = time.Parse(formatDateTimeFul3, string(data))
	}


	if err != nil{
		var sec int
		if sec, err = strconv.Atoi(string(data)); err == nil{
			tt = time.Unix(int64(sec), 0)
			//IsShortFormat = true;
		}
	}

	if err != nil{
		var valF float64
		if valF, err = strconv.ParseFloat(string(data), 64); err == nil{
			tt = time.Unix(int64(valF), 0)
		}
	}

	//*t = TDateTime{tt, IsShortFormat}
	*t = TDateTime{tt}
	return
}

