package handler

import (
	"main/core"
	"reflect"
)

type ArgHandler struct {
}

func (*ArgHandler) Run(cmdCtx core.CommandContext) error {
	var retInfo RetInfo
	// (RetInfo)
	retInfo = cmdCtx.Stack.Pop().(RetInfo)
	retInfo = RetInfo{len(cmdCtx.Values)}
	defer cmdCtx.Stack.Push(retInfo)

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
