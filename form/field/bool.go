package field

import "fmt"

type boolValidator struct{}

func (f *boolValidator) Validate(val []byte) (isValid bool, msg string) {
	switch {
	case string(val) == "true":
		isValid = true
	case string(val) == "false":
		isValid = true
	default:
		msg = fmt.Sprintf("`%s` is invalid, valid values are `true` or `false`\n")
	}
	return
}

type boolScanner struct {
	dst *bool
}

func (s *boolScanner) Scan(val interface{}) error {
	if t, ok := val.(bool); ok {
		*s.dst = t
		return nil
	}
	return fmt.Errorf("Invalid value %v (%T), expect `bool` type", val, val)
}

func NewBool(dst *bool, name, msg string) RequiredValidatedField {
	return &inputField{name, msg, true, &boolValidator{}, &boolScanner{dst}}
}

func NewBoolOpt(dst *bool, name, msg string) ValidatedField {
	return &inputField{name, msg, false, &boolValidator{}, &boolScanner{dst}}
}
