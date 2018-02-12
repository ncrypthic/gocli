package field

import (
	"fmt"
	"strconv"
)

type intValidator struct{}

func (v *intValidator) Validate(val []byte) bool {
	str := string(val)
	_, err := strconv.ParseInt(str, 10, 64)
	return err == nil
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
	invalidMsg := fmt.Sprintf("Field `%s` value must be an int64", name)
	return &inputField{name, msg, invalidMsg, true, &intValidator{}, &intScanner{dst}}
}

func NewIntOpt(dst *int64, name, msg string) ValidatedField {
	invalidMsg := fmt.Sprintf("Field `%s` value must be an int64", name)
	return &inputField{name, msg, invalidMsg, false, &intValidator{}, &intScanner{dst}}
}
