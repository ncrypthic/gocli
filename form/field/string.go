package field

import "fmt"

type stringValidator struct{}

func (v *stringValidator) Validate(val []byte) (isValid bool, msg string) {
	return true, ""
}

type stringScanner struct {
	dst *string
}

func (s *stringScanner) Scan(val interface{}) error {
	switch t := val.(type) {
	case string:
		*s.dst = t
		return nil
	case []byte:
		*s.dst = string(t)
		return nil
	}
	return fmt.Errorf("Invalid value %v (%T), expect `string` type", val, val)
}

func NewString(dst *string, name, msg string) RequiredValidatedField {
	return &inputField{name, msg, true, &stringValidator{}, &stringScanner{dst}}
}

func NewStringOpt(dst *string, name, msg string) ValidatedField {
	return &inputField{name, msg, false, &stringValidator{}, &stringScanner{dst}}
}
