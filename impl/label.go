package impl

type Label struct {
	f     *Func
	label string
}

func (l *Label) Goto() error {
	return l.f.Goto(l.label)
}
