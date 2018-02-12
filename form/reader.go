package form

type FormReader interface {
	Read([]byte) (int, error)
}
