package impl

import "main/core"

type VarTable struct {
	advanceStack *AdvanceStack
	vars         map[string]*Var
}

func NewVarTable(cmdCtx core.CommandContext) *VarTable {
	return &VarTable{
		advanceStack: cmdCtx.Stack.(*AdvanceStack),
		vars:         make(map[string]*Var),
	}
}

func (v *VarTable) NewVar(varName string) core.Var {
	newVar := &Var{v.advanceStack, v.advanceStack.Len()}
	v.vars[varName] = newVar
	v.advanceStack.Push(nil)
	return newVar
}

func (v *VarTable) SearchVar(varName string) (core.Var, bool) {
	_var, ok := v.vars[varName]
	return _var, ok
}

func (v *VarTable) DebugInfo() Detail {
	list := make([]string, v.advanceStack.Len())
	for name, value := range v.vars {
		list[value.position] = name
	}

	// 栈顶
	n := len(list)
	if n > 0 {
		list[n-1] += "<-"
	}

	return Detail{"Bind var", list}
}
