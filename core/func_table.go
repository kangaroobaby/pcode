package core

type FuncTable interface {
	SearchFunc(funcName string) (Func, bool)
}
