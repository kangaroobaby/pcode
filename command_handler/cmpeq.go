package command_handler

import "main/core"

type EqualHandler struct {
}

func (*EqualHandler) Run(cmdCtx core.CommandContext) error {
	b := cmdCtx.Stack.Pop().(int)
	a := cmdCtx.Stack.Pop().(int)
	if a == b {
		cmdCtx.Stack.Push(1)
	} else {
		cmdCtx.Stack.Push(0)
	}

	return nil
}
