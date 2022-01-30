package command_handler

import (
	"fmt"
	"main/core"
)

type PopHandler struct {
}

func (*PopHandler) Run(cmdCtx core.CommandContext) error {
	var i = cmdCtx.Stack.Pop().(int)

	if len(cmdCtx.Values) > 0 {
		name := cmdCtx.Values[0].(string)
		searchVar, ok := cmdCtx.VarTable.SearchVar(name)
		if !ok {
			return fmt.Errorf("can't find var: %s", name)
		}

		searchVar.SetValue(i)
	}

	return nil
}
