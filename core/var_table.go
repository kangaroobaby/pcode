package core

type VarTable interface {
	NewVar(varName string) Var
	SearchVar(varName string) (Var, bool)
}
