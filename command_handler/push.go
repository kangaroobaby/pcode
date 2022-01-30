package command_handler

import (
	"fmt"
	"main/core"
)

type PushHandler struct {
}

func (*PushHandler) Run(cmdCtx core.CommandContext) error {
	switch v := cmdCtx.Values[0].(type) {
	case int:
		cmdCtx.Stack.Push(v)
	case string:
		searchVar, ok := cmdCtx.VarTable.SearchVar(v)
		if !ok {
			return fmt.Errorf("can't find var: %s", v)
		}

		cmdCtx.Stack.Push(searchVar.GetValue())
	default:
		return fmt.Errorf("push value error: %v", cmdCtx.Values[0])
	}

	return nil
}
