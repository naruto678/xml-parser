package main

import "fmt"

type Printer interface {
	Print()
}

type ASTPrinter struct {
	Root Node
}

func (p ASTPrinter) Print() {

}

type BFSPrinter struct {
	Root Node
}

func (p BFSPrinter) Print() {
	queue := []Node{}
	queue = append(queue, p.Root)
	println(p.Root.Name)
	for len(queue) > 0 {
		result := ""
		for size := len(queue); size > 0; size-- {
			top := queue[0]
			queue = queue[1:]
			for _, child := range top.Children {
				result += child.Name + ""
				queue = append(queue, *child)
			}

		}
		fmt.Println(result)
	}

}
