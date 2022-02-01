package handler

import "main/core"

type MulHandler struct {
}

func (*MulHandler) Run(cmdCtx core.CommandContext) error {
	b := cmdCtx.Stack.Pop().(int)
	a := cmdCtx.Stack.Pop().(int)
	cmdCtx.Stack.Push(a * b)
	return nil
}
