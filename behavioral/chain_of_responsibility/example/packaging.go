package main

import (
	"fmt"
)

type packaging struct {
	next section
}

func (p *packaging) execute(t *task) {
	if t.packagingExecuted {
		fmt.Println("Packaging already done")
		p.next.execute(t)
		return
	}
	fmt.Println("Packaging section doing packaging")
}

func (p *packaging) setNext(next section) {
	p.next = next
}