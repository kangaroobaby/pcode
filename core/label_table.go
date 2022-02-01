package core

type LabelTable interface {
	NewLabel(labelName string) Label
	SearchLabel(labelName string) (Label, bool)
}
