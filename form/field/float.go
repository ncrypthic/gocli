package field

import (
	"fmt"
	"strconv"
)

type floatValidator struct{}

func (v *floatValidator) Validate(val []byte) (isValid bool, msg string) {
	str := string(val)
	_, err := strconv.ParseFloat(str, 64)
	if err == nil {
		isValid = true
	} else {
		msg = fmt.Sprintf("`%s` is invalid, expected float value\n")
	}
	return
}

type floatScanner struct {
	dst *float64
}

func (s *floatScanner) Scan(val interface{}) error {
	switch t := val.(type) {
	case string:
		parsed, err := strconv.ParseFloat(t, 64)
		if err != nil {
			return err
		}
		*s.dst = parsed
		return nil
	case float32:
		*s.dst = float64(t)
		return nil
	case float64:
		*s.dst = t
		return nil
	case int:
		*s.dst = float64(t)
		return nil
	case int8:
		*s.dst = float64(t)
		return nil
	case int16:
		*s.dst = float64(t)
		return nil
	case int32:
		*s.dst = float64(t)
		return nil
	case int64:
		*s.dst = float64(t)
		return nil
	}
	return fmt.Errorf("Invalid value %v (%T), expect `float64` type", val, val)
}

func NewFloat(dst *float64, name, msg string) RequiredValidatedField {
	return &inputField{name, msg, true, &floatValidator{}, &floatScanner{dst}}
}

func NewFloatOpt(dst *float64, name, msg string) ValidatedField {
	return &inputField{name, msg, false, &floatValidator{}, &floatScanner{dst}}
}
