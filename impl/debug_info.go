package impl

import (
	"fmt"
	"main/core"
)

type DebugInfo interface {
	DebugInfo() Detail
}

type Detail struct {
	Title string
	List  []string
}

func PrintDebugInfoWithContext(cmdCtx core.CommandContext) {
	debugInfos := []DebugInfo{
		cmdCtx.Func.(*Func),
		cmdCtx.Stack.(*AdvanceStack),
		cmdCtx.VarTable.(*VarTable),
	}

	PrintDebugInfo(debugInfos)
}

func PrintDebugInfo(debugInfos []DebugInfo) {
	details := make([]Detail, len(debugInfos))
	for i, debugInfo := range debugInfos {
		details[i] = debugInfo.DebugInfo()
	}

	// 计算格式、宽度
	format := ""
	for _, detail := range details {
		width := len(detail.Title)
		for _, value := range detail.List {
			if width < len(value) {
				width = len(value)
			}
		}
		format += fmt.Sprintf("|  %%-%ds  ", width)
	}
	if format[0] == '|' {
		format = format[1:]
	}
	format += "\n"

	// 计算列长度
	columnLen := 0

	// 打印title
	titles := make([]interface{}, len(details))
	for i, _ := range details {
		titles[i] = details[i].Title
		if columnLen < len(details[i].List) {
			columnLen = len(details[i].List)
		}
	}
	fmt.Printf(format, titles...)

	// 打印数据区
	for k := 0; k < columnLen+5; k++ {
		columns := make([]interface{}, len(details))
		for i, _ := range details {
			if k < len(details[i].List) {
				columns[i] = details[i].List[k]
			} else {
				columns[i] = ""
			}
		}
		fmt.Printf(format, columns...)
	}

}
