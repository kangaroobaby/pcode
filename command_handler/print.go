package command_handler

import (
	"fmt"
	"main/core"
)

type PrintHandler struct {
}

func (*PrintHandler) Run(cmdCtx core.CommandContext) error {
	var format = cmdCtx.Values[0].(string)
	var values = make([]interface{}, 0, 5)

	line := format + "\n"
	for {
		for {
			if line[0] == '%' || line[0] == '\n' {
				break
			}
			line = line[1:]
		}

		if line[0] == '\n' {
			break
		}

		line = line[1:]

		if line[0] == 'd' {
			value := cmdCtx.Stack.Pop()
			values = append(values, value)
		}
	}

	reverse(values)

	_, err := fmt.Printf(format, values...)
	return err
}
