package main

import (
	"fmt"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	content, err := os.ReadFile("./resources/test/simple.xml")
	if err != nil {
		t.Fatal("Cannot read file ", err.Error())
	}
	source := string(content)

	tokens := Lexer(source)

	if len(tokens) == 0 {
		t.Fatal("Lexer failed")
	}

	parser := NewParser(tokens)
	root_node, err := parser.Parse()

	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println("This is my name", root_node.Name)
	fmt.Println("These are my children", root_node.Children)
	for _, child := range root_node.Children {
		fmt.Println(child)
		fmt.Println(child.TextValue)
	}

	bfs_printer := BFSPrinter{
		Root: root_node,
	}
	parser.Print(bfs_printer)
}
