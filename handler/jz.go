package handler

import (
	"fmt"
	"main/core"
)

type JzHandler struct {
}

func (*JzHandler) Run(cmdCtx core.CommandContext) error {
	i := cmdCtx.Stack.Pop().(int)
	if i == 0 {
		label := cmdCtx.Values[0].(string)
		searchLabel, ok := cmdCtx.LabelTable.SearchLabel(label)
		if !ok {
			return fmt.Errorf("can't find label: %s", label)
		}

		return searchLabel.Goto()
	}

	return nil
}
