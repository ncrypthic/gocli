package field

type Value interface{}

//Validator is a function to validate user input which returns boolean
type Validator interface {
	Validate([]byte) bool
}

type Scanner interface {
	Scan(interface{}) error
}

type Field interface {
	Name() string
	Message() string
	InvalidMessage() string
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
	name           string
	message        string
	invalidMessage string
	required       bool
	validator      Validator
	scanner        Scanner
}

func (f *inputField) Name() string {
	return f.name
}

func (f *inputField) Message() string {
	return f.message
}

func (f *inputField) InvalidMessage() string {
	return f.invalidMessage
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

func NewField(name, msg, invalidMsg string, required bool, validator Validator, scanner Scanner) Field {
	return &inputField{name, msg, invalidMsg, required, validator, scanner}
}
