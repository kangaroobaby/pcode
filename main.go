package main

import (
	"context"
	"errors"
	"flag"
	"io/ioutil"
	"main/core"
	"main/handler"
	"main/impl"
)

var sourceFile = flag.String("f", "demo.asm", "pcode source file")

func main() {
	flag.Parse()
	content, err := ioutil.ReadFile(*sourceFile)
	check(err)

	funcTable := impl.NewFuncTable(string(content))
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
	cmdCtx.Stack = impl.NewAdvanceStack()
	cmdCtx.Func = f
	cmdCtx.FuncTable = funcTable
	cmdCtx.VarTable = impl.NewVarTable(cmdCtx)
	cmdCtx.LabelTable = impl.NewLabelTable(cmdCtx)

	callHandler := &handler.CallHandler{}
	err = callHandler.Call(cmdCtx)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
