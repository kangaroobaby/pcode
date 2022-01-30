package command_handler

import "main/core"

type JzHandler struct {
}

func (*JzHandler) Run(cmdCtx core.CommandContext) error {
	i := cmdCtx.Stack.Pop().(int)
	if i == 0 {
		label := cmdCtx.Values[0].(string)
		return cmdCtx.Func.Goto(label)
	}

	return nil
}
