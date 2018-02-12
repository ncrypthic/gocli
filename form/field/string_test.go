package field

import (
	"testing"
)

func TestStringValidator(t *testing.T) {
	validator := &stringValidator{}
	if isTrue, _ := validator.Validate([]byte("1")); !isTrue {
		t.Error("`1` should be a valid string value")
	}
	if isTrue, _ := validator.Validate([]byte("true")); !isTrue {
		t.Error("`true` should be a valid string value")
	}
}

func TestStringScanner(t *testing.T) {
	var dst *string
	isString := "1"
	dst = &isString
	scanner := &stringScanner{dst}
	if err := scanner.Scan("1"); err != nil {
		t.Error("`1` should not be scannable by `stringScanner`")
	}
	if err := scanner.Scan(1); err == nil {
		t.Error("`1` should not be scanned successfully by `stringScanner`")
	}
	if err := scanner.Scan(1.0); err == nil {
		t.Error("`1.0` should not be scanned successfully by `stringScanner`")
	}
	if err := scanner.Scan("1"); err != nil {
		t.Error("`\"1\"` should be scanned successfully by `stringScanner`")
	}
	if err := scanner.Scan("1.0"); err != nil {
		t.Error("`\"1.0\"` should be scanned successfully by `stringScanner`")
	}
	if err := scanner.Scan([]byte("1")); err != nil {
		t.Error("`[]byte(\"1\")` should be scanned successfully by `stringScanner`")
	}
}
