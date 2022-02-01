package handler

import (
	"bytes"
	"flag"
	"fmt"
	"main/core"
	"main/impl"
)

var debug = flag.Bool("d", false, "open debug mode")

var stringBuffer bytes.Buffer // print for debug

type CallHandler struct {
}

func (c *CallHandler) Run(cmdCtx core.CommandContext) error {
	var err error
	var immediate bool

	for {
		if *debug {
			impl.PrintDebugInfoWithContext(cmdCtx)
			fmt.Printf("\n***Terminal***\n%s\n", stringBuffer.String())
			if !immediate {
				var input string
				fmt.Printf("\n\npress enter to step, -r to run.")
				fmt.Scanf("%s\n", &input)
				if input == "-r" {
					immediate = true
				}
			}
		}

		cmdCtx.Command, err = cmdCtx.Func.Next()
		if err == core.EOF {
			break
		}
		check(err)

		handler, ok := getHandler(cmdCtx)
		if !ok {
			// call ?
			if cmdCtx.Operate[0] == '$' {
				newCtx := cmdCtx
				newCtx.Func, ok = newCtx.FuncTable.SearchFunc(cmdCtx.Operate[1:])
				if !ok {
					return fmt.Errorf("can't find func: %v", cmdCtx.Command)
				}
				newCtx.VarTable = impl.NewVarTable(newCtx)
				newCtx.LabelTable = impl.NewLabelTable(newCtx)

				func() {
					newCtx.Func.Seek(core.NullPointer)
					defer cmdCtx.Func.Seek(cmdCtx.Pointer)
					check(c.Call(newCtx))
				}()

				continue
			}
			return fmt.Errorf("can't find handler: %v", cmdCtx.Command)
		}

		check(handler.Run(cmdCtx))

		// return ?
		if _, ok = handler.(*ReturnHandler); ok {
			break
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
