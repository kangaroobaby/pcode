package command_handler

import (
	"fmt"
	"main/core"
	"main/runtime"
)

type CallHandler struct {
}

func (c *CallHandler) Run(cmdCtx core.CommandContext) error {
	var err error
	for {
		cmdCtx.Command, err = cmdCtx.Func.Next()
		if err == runtime.EOF {
			break
		}
		if err != nil {
			return err
		}

		handle, ok := getHandler(cmdCtx)
		if !ok {
			// call ?
			if cmdCtx.Operate[0] == '$' {
				newCtx := cmdCtx
				newCtx.VarTable = runtime.NewVarTable(newCtx)
				newCtx.Func, ok = newCtx.FuncTable.SearchFunc(cmdCtx.Operate[1:])
				if !ok {
					return fmt.Errorf("can't find func: %v", cmdCtx.Command)
				}

				newCtx.Func.Seek(0)
				defer cmdCtx.Func.Seek(cmdCtx.Eip)

				err = c.Call(newCtx)
				if err != nil {
					return fmt.Errorf("call error: %v", cmdCtx.Command)
				}
				continue
			}
			return fmt.Errorf("can't find handler: %v", cmdCtx.Command)
		}

		err = handle.Run(cmdCtx)
		if err != nil {
			return fmt.Errorf("handle run error: %v", cmdCtx.Command)
		}

	}

	return nil
}

func (c *CallHandler) Call(cmdCtx core.CommandContext) error {
	cmdCtx.Stack.Push(RetInfo{})
	return c.Run(cmdCtx)
}

func getHandler(cmdCtx core.CommandContext) (core.CommandHandler, bool) {
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
		"cmplt":   &LessthanHandler{},
		"cmpeq":   &EqualHandler{},
		"or":      &OrHandler{},
	}

	handler, ok := handlers[cmdCtx.Operate]
	return handler, ok
}
