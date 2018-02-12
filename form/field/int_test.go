package field

import (
	"testing"
)

func TestIntValidator(t *testing.T) {
	validator := &intValidator{}
	if isTrue := validator.Validate([]byte("n/a")); isTrue {
		t.Error("`n/a` should be an invalid int64 value")
	}
	if isTrue := validator.Validate([]byte("1")); !isTrue {
		t.Error("`1` should be a valid int64 value")
	}
	if isTrue := validator.Validate([]byte("1.0")); isTrue {
		t.Error("`1.0` should be a valid int64 value")
	}
}

func TestIntScanner(t *testing.T) {
	var dst *int64
	isInt64 := int64(1)
	dst = &isInt64
	scanner := &intScanner{dst}
	if err := scanner.Scan("n/a"); err == nil {
		t.Error("`n/a` should not be scannable by `intScanner`")
	}
	if err := scanner.Scan(1); err != nil {
		t.Error("`1.0` should be scanned successfully by `intScanner`")
	}
	if err := scanner.Scan(1.0); err == nil {
		t.Error("`1.0` should be scanned successfully by `intScanner`")
	}
	if err := scanner.Scan("1"); err != nil {
		t.Error("`\"1\"` should be scanned successfully by `intScanner`")
	}
	if err := scanner.Scan("1.0"); err == nil {
		t.Error("`\"1.0\"` should be scanned successfully by `intScanner`")
	}
}
