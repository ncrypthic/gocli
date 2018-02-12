package field

import (
	"testing"
	"time"
)

var layout string = time.RFC3339Nano
var validDatetimeStr string = "2006-01-02T15:04:05.9999+07:00"

func TestTimeValidator(t *testing.T) {
	validator := &timeValidator{layout}
	if isTrue := validator.Validate([]byte(validDatetimeStr)); !isTrue {
		t.Errorf("`%s` should be a valid time value", validDatetimeStr)
	}
	if isTrue := validator.Validate([]byte("true")); isTrue {
		t.Error("`true` should not be a valid time value")
	}
}

func TestTimeScanner(t *testing.T) {
	var dst *time.Time
	emptyTime := time.Time{}
	isTime := time.Now()
	dst = &isTime
	scanner := &timeScanner{layout, dst}
	if err := scanner.Scan(validDatetimeStr); err != nil {
		t.Errorf("`%s` should not be scannable by `timeScanner`", validDatetimeStr)
	}
	if err := scanner.Scan(emptyTime); err != nil {
		t.Errorf("`%v` should be scanned successfully by `timeScanner`", emptyTime)
	}
	if err := scanner.Scan("n/a"); err == nil {
		t.Error("`n/a` should not be scanned successfully by `stringScanner`")
	}
	if err := scanner.Scan("1.0"); err == nil {
		t.Error("`\"1.0\"` should not be scanned successfully by `stringScanner`")
	}
}
