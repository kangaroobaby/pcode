package handler

import "main/core"

type OrHandler struct {
}

func (*OrHandler) Run(cmdCtx core.CommandContext) error {
	b := cmdCtx.Stack.Pop().(int)
	a := cmdCtx.Stack.Pop().(int)
	if a == 1 || b == 1 {
		cmdCtx.Stack.Push(1)
	} else {
		cmdCtx.Stack.Push(0)
	}

	return nil
}
