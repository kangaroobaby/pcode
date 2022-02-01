package impl

import (
	"errors"
	"fmt"
)

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

func (a *AdvanceStack) DebugInfo() Detail {
	list := make([]string, 0, 100)
	for _, value := range a.mem {
		if value == nil {
			list = append(list, "/")
		} else {
			list = append(list, fmt.Sprintf("%d", value))
		}
	}
	list = list[1:]

	return Detail{"Stack", list}
}
