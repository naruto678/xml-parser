package main

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	xml_source := "<A>FirstTest</A>"
	tokens := Lexer(xml_source)
	for _, token := range tokens {
		fmt.Println(token.String())
	}
	xml_source = `
		<A>
			<B>some_value</B>
			<C>some_second_value</C>
			<D>some_third_value</D>
			<E> some_more_value </E>
			<F>
				<G>some_nested_value</G>
				<H>some_more_nesting</H>
			</F>
		</A>
	`
	tokens = Lexer(xml_source)
	for _, token := range tokens {
		fmt.Println(token.String())
	}

}
