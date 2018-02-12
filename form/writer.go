package form

type Color int

type Writer interface {
	Write([]byte) (int, error)
}

type ColorWriter interface {
	Writer
	WriteColor(Color, []byte)
}
