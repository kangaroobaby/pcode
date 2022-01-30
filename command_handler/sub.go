package command_handler

import "main/core"

type SubHandler struct {
}

func (*SubHandler) Run(cmdCtx core.CommandContext) error {
	b := cmdCtx.Stack.Pop().(int)
	a := cmdCtx.Stack.Pop().(int)
	cmdCtx.Stack.Push(a - b)
	return nil
}
