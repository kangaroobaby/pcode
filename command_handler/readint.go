package command_handler

import (
	"fmt"
	"main/core"
)

type ReadHandler struct {
}

func (*ReadHandler) Run(cmdCtx core.CommandContext) error {
	var i int

	fmt.Print(cmdCtx.Values[0])
	fmt.Scan(&i)

	cmdCtx.Stack.Push(i)
	return nil
}
