package command_handler

import (
	"fmt"
	"main/core"
)

type ReturnHandler struct {
}

func (*ReturnHandler) Run(cmdCtx core.CommandContext) error {
	var retValue interface{}
	//ret         ; 返回空值 “/”
	//ret 1       ; 返回常数
	//ret a       ; 返回变量值
	//ret ~       ; 取出栈顶元素，返回其值。
	if len(cmdCtx.Values) == 1 {
		switch v := cmdCtx.Values[0].(type) {
		case int:
			retValue = v
		case string:
			switch v {
			case "~":
				retValue = cmdCtx.Stack.Pop()
			default:
				searchVar, ok := cmdCtx.VarTable.SearchVar(v)
				if !ok {
					return fmt.Errorf("can't find var: %s", v)
				}

				retValue = searchVar.GetValue()
			}
		default:
			fmt.Errorf("return value error: %v", cmdCtx.Values)
		}
	}

	var retInfo RetInfo
	for {
		i, ok := cmdCtx.Stack.Pop().(RetInfo)
		if ok {
			retInfo = i
			break
		}
	}

	// 回滚参数
	for i := 0; i < retInfo.ArgNumber; i++ {
		cmdCtx.Stack.Pop()
	}

	// 放入返回值
	cmdCtx.Stack.Push(retValue)

	return nil
}
