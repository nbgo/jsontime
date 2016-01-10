package jsontime

import (
	"time"
	"fmt"
	"strings"
)

// RFC3339Nano is a time type that (un)marshaled as time in RFC3339Nano format.
type RFC3339Nano time.Time


// MarshalJSON marshals RFC3339Nano into JSON
//noinspection GoMethodOnNonLocalType
func (t RFC3339Nano) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RFC3339Nano))
	return []byte(stamp), nil
}

// UnmarshalJSON unmarshals RFC3339Nano from JSON
//noinspection GoMethodOnNonLocalType
func (t *RFC3339Nano) UnmarshalJSON(data []byte) error {
	parsedTime, err := time.Parse(time.RFC3339Nano, strings.Trim(string(data), "\""))
	if err != nil {
		return err
	}
	jt := RFC3339Nano(parsedTime)
	*t = jt
	return nil
}