package command_handler

import (
	"fmt"
	"main/core"
	"os"
)

type ExitHandler struct {
}

func (*ExitHandler) Run(cmdCtx core.CommandContext) error {
	switch v := cmdCtx.Values[0].(type) {
	case int:
		os.Exit(v)
	case string:
		switch v {
		case "~":
			i := cmdCtx.Stack.Pop().(int)
			os.Exit(i)
		default:
			searchVar, ok := cmdCtx.VarTable.SearchVar(v)
			if !ok {
				return fmt.Errorf("can't find var: %s", v)
			}

			os.Exit(searchVar.GetValue())
		}
	default:
		return fmt.Errorf("exit value error: %v", cmdCtx.Values)
	}

	return nil
}
