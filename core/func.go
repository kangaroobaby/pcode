package core

type Func interface {
	Next() (Command, error)
	Goto(labelName string) error
}
