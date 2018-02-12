package field

import (
	"fmt"
	"time"
)

type timeValidator struct {
	layout string
}

func (v *timeValidator) Validate(val []byte) bool {
	str := string(val)
	_, err := time.Parse(v.layout, str)
	return err == nil
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
	invalidMsg := fmt.Sprintf("Field `%s` value must be a valid time", name)
	return &inputField{name, msg, invalidMsg, true, &timeValidator{layout}, &timeScanner{layout, dst}}
}

func NewTimeOpt(dst *time.Time, layout, name, msg string) ValidatedField {
	invalidMsg := fmt.Sprintf("Field `%s` value must be a valid time", name)
	return &inputField{name, msg, invalidMsg, false, &timeValidator{layout}, &timeScanner{layout, dst}}
}
