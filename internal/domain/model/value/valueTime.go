package value

import(
	"fmt"
	"time"
	"strings"
)

const timeLayout = "2006/01/02 15:04:05"

var NullTime time.Time
type valueTime struct {
	time.Time
}

func NewValueTime(t *time.Time) *valueTime{
	if t != nil {
		return &valueTime{Time: *t}
	}
	return nil
}

func (v *valueTime) UnmarshalJSON(b []byte) (error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		v.Time = time.Time{}
		return nil
	}
	var err error
	v.Time, err = time.Parse(timeLayout, s)
	return err
}

func (v *valueTime) MarshalJSON() ([]byte, error) {
	if v.Time.UnixNano() == NullTime.UnixNano() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", v.Time.Format(timeLayout))), nil
}
