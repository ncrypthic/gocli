package form

type Reader interface {
	Read([]byte) (int, error)
}
