package field

type Value interface{}

//Validator is a function to validate user input which returns boolean
type Validator interface {
	Validate([]byte) (bool, string)
}

type Scanner interface {
	Scan(interface{}) error
}

type Field interface {
	Name() string
	Message() string
	Scanner() Scanner
}

type ValidatedField interface {
	Field
	Validator() Validator
}

type RequiredField interface {
	Field
	Empty([]byte) bool
}

type RequiredValidatedField interface {
	Field
	Empty([]byte) bool
	Validator() Validator
}

type inputField struct {
	name      string
	message   string
	required  bool
	validator Validator
	scanner   Scanner
}

func (f *inputField) Name() string {
	return f.name
}

func (f *inputField) Message() string {
	return f.message
}

func (f *inputField) Empty(val []byte) bool {
	return len(val) == 0 || string(val) == ""
}

func (f *inputField) Validator() Validator {
	return f.validator
}

func (f *inputField) Scanner() Scanner {
	return f.scanner
}

func NewField(name, msg string, required bool, validator Validator, scanner Scanner) Field {
	return &inputField{name, msg, required, validator, scanner}
}

func WithValidator(f Field, validator Validator) Field {
	_, required := f.(RequiredField)
	return &inputField{f.Name(), f.Message(), required, validator, f.Scanner()}
}

func WithScanner(f Field, scanner Scanner) Field {
	_, required := f.(RequiredField)
	var validator Validator
	if validatedField, hasValidator := f.(ValidatedField); hasValidator {
		validator = validatedField.Validator()
	}
	return &inputField{f.Name(), f.Message(), required, validator, scanner}
}
