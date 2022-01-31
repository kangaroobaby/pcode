package command_handler

import (
	"main/core"
	"reflect"
)

type ArgHandler struct {
}

func (*ArgHandler) Run(cmdCtx core.CommandContext) error {
	// (RetInfo)
	retInfo := cmdCtx.Stack.Pop().(RetInfo)
	defer cmdCtx.Stack.Push(retInfo)

	retInfo.ArgNumber = len(cmdCtx.Values)

	args := make([]int, len(cmdCtx.Values))
	for i, _ := range cmdCtx.Values {
		args[i] = cmdCtx.Stack.Pop().(int)
	}

	reverse(args)
	for i, value := range cmdCtx.Values {
		newVar := cmdCtx.VarTable.NewVar(value.(string))
		newVar.SetValue(args[i])
	}

	return nil
}

func reverse(s interface{}) {
	swapFunc := reflect.Swapper(s)
	for i, j := 0, reflect.ValueOf(s).Len()-1; i < j; i, j = i+1, j-1 {
		swapFunc(i, j)
	}
}
