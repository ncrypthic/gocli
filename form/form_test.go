package form

import (
	"github.com/ncrypthic/gocli/form/field"
	"testing"
)

type dummyReader struct {
	data []byte
}

func (r *dummyReader) Read() ([]byte, error) {
	return r.data, nil
}

type dummyWriter struct{}

func (w *dummyWriter) Write(d []byte) error {
	return nil
}

type dummyStruct struct {
	data string
}

func TestForm(t *testing.T) {
	f := NewForm()
	d := dummyStruct{"n/a"}
	dummyData := []byte("test")
	fields := []field.Field{
		field.NewString(&d.data, "string", "Please input string data"),
	}
	f.Prompt(&dummyWriter{}, &dummyReader{dummyData}, fields)
	if d.data != "test" {
		t.Errorf(`Expected form.data == "test", got %s instead`, d.data)
	}
}
