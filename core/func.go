package core

type Func interface {
	Next() (Command, error)
	Seek(Pointer) error
}
