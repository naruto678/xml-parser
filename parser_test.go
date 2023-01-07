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

	for _, token := range tokens {
		fmt.Println(token.String())
	}

}
