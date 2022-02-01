package impl

type Var struct {
	advanceStack *AdvanceStack
	position     int
}

func (v *Var) GetValue() int {
	return v.advanceStack.Read(v.position).(int)
}

func (v *Var) SetValue(i int) {
	v.advanceStack.Write(v.position, i)
}
