package entity

import(
	"fmt"
	"time"
	"strings"
	"database/sql"
)

const timeLayout = "2006/01/02 15:04:05"
var nullTime time.Time

type sqlNullTime struct{
	sql.NullTime
}

func NewSqlNullTime(t *time.Time) *sqlNullTime{
	if t == nil {
		return &sqlNullTime{NullTime: sql.NullTime{}}	
	}
	return &sqlNullTime{NullTime: sql.NullTime{Time: *t, Valid: true}}
}

func (v *sqlNullTime) UnmarshalJSON(b []byte) (error) {
	fmt.Printf("UnmarshalJSON in: %s\n", string(b))
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		v.NullTime = sql.NullTime{}
		return nil
	}
	t, err := time.Parse(timeLayout, s)
	if err != nil {
		return err
	}
	v.NullTime = sql.NullTime{Time: t, Valid: true}
	return err
}

func (v *sqlNullTime) MarshalJSON() ([]byte, error) {
	if v.NullTime.Time.UnixNano() == nullTime.UnixNano() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", v.NullTime.Time.Format(timeLayout))), nil
}