package handler

import (
	"fmt"
	"main/core"
)

type JmpHandler struct {
}

func (*JmpHandler) Run(cmdCtx core.CommandContext) error {
	label := cmdCtx.Values[0].(string)
	searchLabel, ok := cmdCtx.LabelTable.SearchLabel(label)
	if !ok {
		return fmt.Errorf("can't find label: %s", label)
	}

	return searchLabel.Goto()
}
