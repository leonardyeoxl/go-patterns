package main

import (
	"fmt"
)

type material struct {
	next section
}

func (m *material) execute(t *task) {
	if t.materialCollected {
		fmt.Println("Material already collected")
		m.next.execute(t)
		return
	}
	fmt.Println("Material section gathering materials")
	t.materialCollected = true
	m.next.execute(t)
}

func (m *material) setNext(next section) {
	m.next = next
}