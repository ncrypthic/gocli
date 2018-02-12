package field

import "fmt"

type boolValidator struct{}

func (f *boolValidator) Validate(val []byte) bool {
	switch {
	case string(val) == "true":
		return true
	case string(val) == "false":
		return false
	default:
		return false
	}
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
	invalidMsg := fmt.Sprintf("Field `%s` value must be (true | false)", name)
	return &inputField{name, msg, invalidMsg, true, &boolValidator{}, &boolScanner{dst}}
}

func NewBoolOpt(dst *bool, name, msg string) ValidatedField {
	invalidMsg := fmt.Sprintf("Field `%s` value must be (true | false)", name)
	return &inputField{name, msg, invalidMsg, false, &boolValidator{}, &boolScanner{dst}}
}
