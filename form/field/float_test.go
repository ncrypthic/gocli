package field

import (
	"testing"
)

func TestFloatValidator(t *testing.T) {
	validator := &floatValidator{}
	if isTrue, _ := validator.Validate([]byte("n/a")); isTrue {
		t.Error("`n/a` should be an invalid float64 value")
	}
	if isTrue, _ := validator.Validate([]byte("1")); !isTrue {
		t.Error("`1` should be a valid float64 value")
	}
	if isTrue, _ := validator.Validate([]byte("1.0")); !isTrue {
		t.Error("`1.0` should be a valid float64 value")
	}
}

func TestFloatScanner(t *testing.T) {
	var dst *float64
	isFloat := 1.0
	dst = &isFloat
	scanner := &floatScanner{dst}
	if err := scanner.Scan("n/a"); err == nil {
		t.Error("`n/a` should not be scannable by `floatScanner`")
	}
	if err := scanner.Scan(1); err != nil {
		t.Error("`1` should be scanned successfully by `floatScanner`")
	}
	if err := scanner.Scan(1.0); err != nil {
		t.Error("`1.0` should be scanned successfully by `floatScanner`")
	}
	if err := scanner.Scan("1"); err != nil {
		t.Error("`\"1\"` should be scanned successfully by `floatScanner`")
	}
}
