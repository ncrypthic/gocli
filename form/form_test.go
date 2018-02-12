package form

import (
	"github.com/ncrypthic/gocli/form/field"
	"testing"
)

type dummyReader struct {
	data []byte
}

func (r *dummyReader) Read(target []byte) (int, error) {
	if len(target) == 0 {
		return 0, nil
	}
	maxLen := len(target)
	if len(target) > len(r.data) {
		maxLen = len(r.data)
	}
	copy(target, r.data[0:maxLen])
	return maxLen, nil
}

type dummyWriter struct{}

func (w *dummyWriter) Write(d []byte) (int, error) {
	return 0, nil
}

type dummyStruct struct {
	data string
}

func TestForm(t *testing.T) {
	f := NewForm()
	d := dummyStruct{"n/a"}
	dummyData := []byte("test\n")
	fields := []field.Field{
		field.NewString(&d.data, "string", "Please input string data"),
	}
	f.Prompt(&dummyWriter{}, &dummyReader{dummyData}, fields)
	if d.data != "test" {
		t.Errorf(`Expected form.data == "test", got %s instead`, d.data)
	}
}
