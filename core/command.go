package core

type Command struct {
	Operate string
	Values  []interface{}
	Pointer Pointer
}
