package field

import (
	"fmt"
	"time"
)

type timeValidator struct {
	layout string
}

func (v *timeValidator) Validate(val []byte) (isValid bool, msg string) {
	str := string(val)
	_, err := time.Parse(v.layout, str)
	if err == nil {
		isValid = true
	} else {
		msg = fmt.Sprintf("`%s` is invalid, expected date time value.")
	}
	return
}

type timeScanner struct {
	layout string
	dst    *time.Time
}

func (s *timeScanner) Scan(val interface{}) error {
	switch t := val.(type) {
	case string:
		parsed, err := time.Parse(s.layout, t)
		if err != nil {
			return err
		}
		*s.dst = parsed
		return nil
	case time.Time:
		*s.dst = t
		return nil
	}
	return fmt.Errorf("Invalid value %v (%T). `%v` is not a valid time value, expect `time.Time`", val, val, val)
}

func NewTime(dst *time.Time, layout, name, msg string) RequiredValidatedField {
	return &inputField{name, msg, true, &timeValidator{layout}, &timeScanner{layout, dst}}
}

func NewTimeOpt(dst *time.Time, layout, name, msg string) ValidatedField {
	return &inputField{name, msg, false, &timeValidator{layout}, &timeScanner{layout, dst}}
}
