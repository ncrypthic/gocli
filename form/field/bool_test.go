package field

import (
	"testing"
)

func TestBoolValidator(t *testing.T) {
	validator := &boolValidator{}
	if isTrue := validator.Validate([]byte("n/a")); isTrue {
		t.Error("`n/a` should be an invalid bool value")
	}
	if isTrue := validator.Validate([]byte("false")); isTrue {
		t.Error("`false` should be an invalid bool value")
	}
	if isTrue := validator.Validate([]byte("true")); !isTrue {
		t.Error("`true` should be a valid bool value")
	}
}

func TestBoolScanner(t *testing.T) {
	var dst *bool
	isFalse := false
	dst = &isFalse
	scanner := &boolScanner{dst}
	if err := scanner.Scan(123); err == nil {
		t.Error("`123` should not be scannable by `boolScanner`")
	}
	if err := scanner.Scan(false); err != nil {
		t.Error("`false` should be scanned successfully by `boolScanner`")
	}
}
