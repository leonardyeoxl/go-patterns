package main

type section interface {
	execute(*task)
	setNext(section)
}