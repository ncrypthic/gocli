package form

import (
	"bytes"
	"fmt"
	"github.com/ncrypthic/gocli/form/field"
	"io"
)

type Form interface {
	Prompt(FormWriter, FormReader, []field.Field) error
}

type form struct{}

func (f *form) promptField(w FormWriter, r FormReader, inputField field.Field) (result []byte, err error) {
	msg := []byte(inputField.Message())
	if _, err = w.Write(msg); err != nil {
		return result, err
	}
	b := make([]byte, 128)
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

func (f *form) Prompt(w FormWriter, r FormReader, fields []field.Field) error {
	for _, inputField := range fields {
		var value []byte
		switch t := inputField.(type) {
		case field.RequiredValidatedField:
			for t.Empty(value) || !t.Validator().Validate(value) {
				switch {
				case value != nil && t.Empty(value):
					w.Write([]byte(fmt.Sprintf("`%s` is required\n", inputField.Name())))
				case value != nil && t.Validator().Validate(value):
					w.Write([]byte(inputField.InvalidMessage()))
				}
				if res, err := f.promptField(w, r, inputField); err != nil {
					return err
				} else {
					value = res
				}
			}
		case field.RequiredField:
			for t.Empty(value) {
				if value != nil {
					w.Write([]byte(fmt.Sprintf("`%` is required", inputField.Name())))
				}
				if res, err := f.promptField(w, r, inputField); err != nil {
					return err
				} else {
					value = res
				}
			}
		case field.ValidatedField:
			for !t.Validator().Validate(value) {
				if value != nil {
					w.Write([]byte(inputField.InvalidMessage()))
				}
				if res, err := f.promptField(w, r, inputField); err != nil {
					return err
				} else {
					value = res
				}
			}
		default:
			if res, err := f.promptField(w, r, inputField); err != nil {
				return err
			} else {
				value = res
			}
		}
	}
	return nil
}

func NewForm() Form {
	return &form{}
}
