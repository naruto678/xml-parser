package main

type Printer interface {
	Print() string
}

type ASTPrinter struct {
	root_node Node
}

func (p ASTPrinter) Print() {

}
