package core

import "context"

type CommandContext struct {
	Ctx context.Context
	Command
	Stack     Stack
	Func      Func
	FuncTable FuncTable
	VarTable  VarTable
}
