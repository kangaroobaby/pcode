package runtime

import "errors"

type AdvanceStack struct {
	mem []interface{}
}

func NewAdvanceStack() *AdvanceStack {
	return &AdvanceStack{mem: make([]interface{}, 0, 1000)}
}

func (a *AdvanceStack) Len() int {
	return len(a.mem)
}

func (a *AdvanceStack) Push(i interface{}) {
	a.mem = append(a.mem, i)
}

func (a *AdvanceStack) Pop() interface{} {
	if len(a.mem) == 0 {
		panic(errors.New("overflow"))
	}

	n := len(a.mem)
	i := a.mem[n-1]
	a.mem = a.mem[:n-1]
	return i
}

func (a *AdvanceStack) Read(pos int) interface{} {
	if pos >= len(a.mem) {
		panic(errors.New("overflow"))
	}

	return a.mem[pos]
}

func (a *AdvanceStack) Write(pos int, i interface{}) {
	if pos >= len(a.mem) {
		panic(errors.New("overflow"))
	}

	a.mem[pos] = i

}
