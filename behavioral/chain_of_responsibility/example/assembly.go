package main

import (
	"fmt"
)

type assembly struct {
	next section
}

func (a *assembly) execute(t *task) {
	if t.assemblyExecuted {
		fmt.Println("Assembly already done")
		a.next.execute(t)
		return
	}
	fmt.Println("Assembly section assembling...")
	t.assemblyExecuted = true
	a.next.execute(t)
}

func (a *assembly) setNext(next section) {
	a.next = next
}