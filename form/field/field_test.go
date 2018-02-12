package field

import (
	"testing"
	"time"
)

func TestBoolField(t *testing.T) {
	var dst *bool
	tmp := false
	dst = &tmp
	msg := "Please enter bool value"
	name := "bool"
	boolField := NewBool(dst, "bool", "Please enter bool value")
	if boolField.Message() != msg {
		t.Errorf("BoolField message should be %s", msg)
	}
	if boolField.Name() != name {
		t.Errorf("BoolField name should be %s", name)
	}
	if boolField.Validator() == nil {
		t.Errorf("BoolField validator should not nil")
	}
	if isValid, _ := boolField.Validator().Validate([]byte("true")); !isValid {
		t.Errorf("Value `true` should be valid with BoolField type validator")
	}
	if boolField.Scanner() == nil {
		t.Errorf("BoolField scanner should not nil")
	}
	if err := boolField.Scanner().Scan(true); err != nil {
		t.Errorf("Value `true` should be scanned successfully by `BoolField` scanner")
	}
	if *dst != true {
		t.Errorf("Failed to scan `true` to destination variable")
	}
}

func TestFloatField(t *testing.T) {
	var dst *float64
	tmp := float64(0)
	dst = &tmp
	msg := "Please enter float value"
	name := "float"
	floatField := NewFloat(dst, "float", "Please enter float value")
	if floatField.Message() != msg {
		t.Errorf("FloatField message should be %s", msg)
	}
	if floatField.Name() != name {
		t.Errorf("FloatField name should be %s", name)
	}
	if floatField.Validator() == nil {
		t.Errorf("FloatField validator should not nil")
	}
	if isValid, _ := floatField.Validator().Validate([]byte("1.2")); !isValid {
		t.Errorf("Value `1.2` should be valid with FloatField type validator")
	}
	if floatField.Scanner() == nil {
		t.Errorf("FloatField scanner should not nil")
	}
	if err := floatField.Scanner().Scan(1.2); err != nil {
		t.Errorf("Value `1.2` should be scanned successfully by `FloatField` scanner")
	}
	if *dst != 1.2 {
		t.Errorf("Failed to scan `1.2` to destination variable")
	}
}

func TestIntField(t *testing.T) {
	var dst *int64
	tmp := int64(0)
	dst = &tmp
	msg := "Please enter int value"
	name := "int"
	intField := NewInt(dst, "int", "Please enter int value")
	if intField.Message() != msg {
		t.Errorf("IntField message should be %s", msg)
	}
	if intField.Name() != name {
		t.Errorf("IntField name should be %s", name)
	}
	if intField.Validator() == nil {
		t.Errorf("IntField validator should not nil")
	}
	if isValid, _ := intField.Validator().Validate([]byte("1")); !isValid {
		t.Errorf("Value `1` should be valid with IntField type validator")
	}
	if intField.Scanner() == nil {
		t.Errorf("IntField scanner should not nil")
	}
	if err := intField.Scanner().Scan(1); err != nil {
		t.Errorf("Value `1` should be scanned successfully by `IntField` scanner")
	}
	if *dst != 1 {
		t.Errorf("Failed to scan `1` to destination variable")
	}
}

func TestStringField(t *testing.T) {
	var dst *string
	tmp := ""
	dst = &tmp
	msg := "Please enter string value"
	name := "string"
	stringField := NewString(dst, "string", "Please enter string value")
	if stringField.Message() != msg {
		t.Errorf("StringField message should be %s", msg)
	}
	if stringField.Name() != name {
		t.Errorf("StringField name should be %s", name)
	}
	if stringField.Validator() == nil {
		t.Errorf("StringField validator should not nil")
	}
	if isValid, _ := stringField.Validator().Validate([]byte("n/a")); !isValid {
		t.Errorf("Value `n/a` should be valid with StringField type validator")
	}
	if stringField.Scanner() == nil {
		t.Errorf("StringField scanner should not nil")
	}
	if err := stringField.Scanner().Scan("n/a"); err != nil {
		t.Errorf("Value `1` should be scanned successfully by `StringField` scanner")
	}
	if *dst != "n/a" {
		t.Errorf("Failed to scan `n/a` to destination variable")
	}
}

func TestTimeField(t *testing.T) {
	var layout string = time.RFC3339Nano
	var validDatetimeStr string = "2006-01-02T15:04:05.9999+07:00"
	var dst *time.Time
	tmp := time.Now()
	dst = &tmp
	msg := "Please enter time value"
	name := "time"
	stringField := NewTime(dst, layout, "time", "Please enter time value")
	if stringField.Message() != msg {
		t.Errorf("TimeField message should be %s", msg)
	}
	if stringField.Name() != name {
		t.Errorf("TimeField name should be %s", name)
	}
	if stringField.Validator() == nil {
		t.Errorf("TimeField validator should not nil")
	}
	if isValid, _ := stringField.Validator().Validate([]byte(validDatetimeStr)); !isValid {
		t.Errorf("Value `n/a` should be valid with TimeField type validator")
	}
	if stringField.Scanner() == nil {
		t.Errorf("TimeField scanner should not nil")
	}
	if err := stringField.Scanner().Scan(validDatetimeStr); err != nil {
		t.Errorf("Value `%s` should be scanned successfully by `TimeField` scanner", validDatetimeStr)
	}
	if dst.Format(layout) != validDatetimeStr {
		t.Errorf("Failed to scan `%s` to destination variable", validDatetimeStr)
	}
	current := time.Now()
	if err := stringField.Scanner().Scan(current); err != nil {
		t.Errorf("Value `%s` should be scanned successfully by `TimeField` scanner", validDatetimeStr)
	}
	if *dst != current {
		t.Errorf("Failed to scan `%v` (%T) to destination variable", current, current)
	}
}
