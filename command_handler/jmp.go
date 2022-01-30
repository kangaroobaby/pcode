package command_handler

import "main/core"

type JmpHandler struct {
}

func (*JmpHandler) Run(cmdCtx core.CommandContext) error {
	label := cmdCtx.Values[0].(string)
	return cmdCtx.Func.Goto(label)
}
