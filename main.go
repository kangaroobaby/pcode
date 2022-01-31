package main

import (
	"context"
	"errors"
	"io/ioutil"
	"main/command_handler"
	"main/core"
	"main/runtime"
	"os"
)

func main() {
	var sourceFile string
	if len(os.Args) > 1 {
		sourceFile = os.Args[1]
	} else {
		sourceFile = "demo.asm"
	}

	content, err := ioutil.ReadFile(sourceFile)
	check(err)

	funcTable := runtime.NewFuncTable(string(content))
	err = funcTable.Initialize()
	check(err)

	f, ok := funcTable.SearchFunc("")
	if !ok {
		f, ok = funcTable.SearchFunc("main")
		if !ok {
			panic(errors.New("main func error"))
		}
	}

	cmdCtx := core.CommandContext{}
	cmdCtx.Ctx = context.Background()
	cmdCtx.Command = core.Command{}
	cmdCtx.Stack = runtime.NewAdvanceStack()
	cmdCtx.Func = f
	cmdCtx.FuncTable = funcTable
	cmdCtx.VarTable = runtime.NewVarTable(cmdCtx)

	callHandler := &command_handler.CallHandler{}
	err = callHandler.Call(cmdCtx)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
