package handler

import "main/core"

type VarHandler struct {
}

func (*VarHandler) Run(cmdCtx core.CommandContext) error {
	for _, value := range cmdCtx.Values {
		cmdCtx.VarTable.NewVar(value.(string))
	}
	return nil
}
