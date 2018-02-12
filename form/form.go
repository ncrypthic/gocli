package form

import (
	"bytes"
	"fmt"
	"github.com/ncrypthic/gocli/form/field"
	"io"
)

type Form interface {
	Prompt(Writer, Reader, []field.Field) error
}

type form struct{}

func (f *form) promptField(w Writer, r Reader, inputField field.Field) (result []byte, err error) {
	msg := []byte(inputField.Message())
	if _, err = w.Write(msg); err != nil {
		return result, err
	}
	b := make([]byte, 1024*32*1000)
	pool := make([][]byte, 0)
	size := 0
	l := 0
	for l, err = r.Read(b); err == nil; {
		if pos := bytes.Index(b, []byte("\n")); pos >= 0 {
			size += pos
			pool = append(pool, b[0:pos])
			err = io.EOF
			break
		} else {
			size += l
			pool = append(pool, b)
			b = make([]byte, 128)
		}
	}
	switch {
	case err == io.EOF:
		result = make([]byte, size)
		for _, buffered := range pool {
			copy(result, buffered[0:len(buffered)])
		}
		inputField.Scanner().Scan(result)
		return result, nil
	default:
		return result, err
	}
}

//Prompt will start command line interaction for specified fields
func (f *form) Prompt(w Writer, r Reader, fields []field.Field) error {
	for _, inputField := range fields {
		switch t := inputField.(type) {
		case field.RequiredValidatedField:
		RequiredValidatedPromptLoop:
			for {
				res, err := f.promptField(w, r, inputField)
				if err != nil {
					return err
				}
				valid, invalidMsg := t.Validator().Validate(res)
				switch {
				case t.Empty(res):
					w.Write([]byte(fmt.Sprintf("`%s` is required\n", inputField.Name())))
				case !valid:
					w.Write([]byte(invalidMsg))
				default:
					break RequiredValidatedPromptLoop
				}
			}
		case field.RequiredField:
		RequiredLoop:
			for {
				res, err := f.promptField(w, r, inputField)
				if err != nil {
					return err
				}
				if t.Empty(res) {
					w.Write([]byte(fmt.Sprintf("`%s` is required\n", inputField.Name())))
				} else {
					break RequiredLoop
				}
			}
		case field.ValidatedField:
		ValidatedLoop:
			for {
				res, err := f.promptField(w, r, inputField)
				if err != nil {
					return err
				}
				if valid, invalidMsg := t.Validator().Validate(res); !valid {
					w.Write([]byte(invalidMsg))
				} else {
					break ValidatedLoop
				}
			}
		default:
			if _, err := f.promptField(w, r, inputField); err != nil {
				return err
			}
		}
	}
	return nil
}

//NewForm creates a new Form
func NewForm() Form {
	return &form{}
}
