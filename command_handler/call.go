package command_handler

import (
	"fmt"
	"main/core"
	"main/runtime"
)

type CallHandler struct {
}

func (c *CallHandler) Run(cmdCtx core.CommandContext) error {
	for {
		cmd, err := cmdCtx.Func.Next()
		if err == runtime.EOF {
			break
		}
		if err != nil {
			return err
		}
		cmdCtx.Command = cmd

		handle, ok := getHandler(cmd)
		if !ok {
			// call ?
			if cmd.Operate[0] == '$' {
				newCtx := cmdCtx
				newCtx.VarTable = runtime.NewVarTable(newCtx)
				newCtx.Func, ok = newCtx.FuncTable.SearchFunc(cmd.Operate[1:])
				if !ok {
					return fmt.Errorf("can't find func: %v", cmd)
				}

				err = c.Call(newCtx)
				if err != nil {
					return fmt.Errorf("call error: %v", cmd)
				}
				continue
			}
			return fmt.Errorf("can't find handler: %v", cmd)
		}

		err = handle.Run(cmdCtx)
		if err != nil {
			return fmt.Errorf("handle run error: %v", cmd)
		}

	}

	return nil
}

func (c *CallHandler) Call(cmdCtx core.CommandContext) error {
	cmdCtx.Stack.Push(RetInfo{})
	return c.Run(cmdCtx)
}

func getHandler(cmd core.Command) (core.CommandHandler, bool) {
	var handlers = map[string]core.CommandHandler{
		"push":    &PushHandler{},
		"pop":     &PopHandler{},
		"ret":     &ReturnHandler{},
		"arg":     &ArgHandler{},
		"var":     &VarHandler{},
		"mul":     &MulHandler{},
		"add":     &AddHandler{},
		"sub":     &SubHandler{},
		"print":   &PrintHandler{},
		"readint": &ReadHandler{},
		"exit":    &ExitHandler{},
		"jmp":     &JmpHandler{},
		"jz":      &JzHandler{},
	}

	handler, ok := handlers[cmd.Operate]
	return handler, ok
}
