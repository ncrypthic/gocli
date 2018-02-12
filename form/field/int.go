package field

import (
	"fmt"
	"strconv"
)

type intValidator struct{}

func (v *intValidator) Validate(val []byte) (isValid bool, msg string) {
	str := string(val)
	_, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		isValid = true
	} else {
		msg = fmt.Sprintf("`%s` is invalid, expected integer value\n")
	}
	return
}

type intScanner struct {
	dst *int64
}

func (s *intScanner) Scan(val interface{}) error {
	switch i := val.(type) {
	case string:
		parsed, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			return err
		}
		*s.dst = parsed
		return nil
	case int:
		*s.dst = int64(i)
		return nil
	case int8:
		*s.dst = int64(i)
		return nil
	case int16:
		*s.dst = int64(i)
		return nil
	case int32:
		*s.dst = int64(i)
		return nil
	case int64:
		*s.dst = i
		return nil
	}
	return fmt.Errorf("Invalid value %v (%T), expect `int64` type", val, val)
}

func NewInt(dst *int64, name, msg string) RequiredValidatedField {
	return &inputField{name, msg, true, &intValidator{}, &intScanner{dst}}
}

func NewIntOpt(dst *int64, name, msg string) ValidatedField {
	return &inputField{name, msg, false, &intValidator{}, &intScanner{dst}}
}
