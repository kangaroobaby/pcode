package handler

import "main/core"

type AddHandler struct {
}

func (*AddHandler) Run(cmdCtx core.CommandContext) error {
	b := cmdCtx.Stack.Pop().(int)
	a := cmdCtx.Stack.Pop().(int)
	cmdCtx.Stack.Push(a + b)
	return nil
}
