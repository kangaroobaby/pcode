package impl

import "main/core"

type LabelTable struct {
	f *Func
}

func NewLabelTable(cmdCtx core.CommandContext) *LabelTable {
	return &LabelTable{
		f: cmdCtx.Func.(*Func),
	}
}

func (l *LabelTable) NewLabel(labelName string) core.Label {
	return &Label{l.f, labelName}
}

func (l *LabelTable) SearchLabel(labelName string) (core.Label, bool) {
	return l.NewLabel(labelName), true
}
