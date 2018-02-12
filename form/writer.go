package form

type Color int

type FormWriter interface {
	Write([]byte) (int, error)
}

type ColorizeFormWriter interface {
	FormWriter
	WriteColor(Color, []byte)
}
