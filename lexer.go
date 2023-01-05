package main

import (
	"fmt"
)

type TokenType byte

const (
	LPAREN TokenType = '<'
	RPAREN TokenType = '>'
	SLASH  TokenType = '/'
	STRING TokenType = '*'
)

//Takes a tokentype and a value
func NewToken(tokenType TokenType, value string, name string) Token {
	return Token{
		TokenType: tokenType,
		Literal:   value,
		Name:      name,
	}
}

type Token struct {
	TokenType
	Literal string
	Name    string
}

func (t *Token) String() string {
	return fmt.Sprintf("Token(%s, %s)", t.Name, t.Literal)
}

func Lexer(source string) []Token {
	var tokens []Token
	current_pos := 0
	for current_pos < len(source) {
		switch source[current_pos] {
		case byte(LPAREN):
			tokens = append(tokens, NewToken(LPAREN, "<", "LPAREN"))
			current_pos += 1
		case byte(RPAREN):
			tokens = append(tokens, NewToken(RPAREN, ">", "RPAREN"))
			current_pos += 1
		case byte(SLASH):
			tokens = append(tokens, NewToken(SLASH, "/", "SLASH"))
			current_pos += 1
		case byte('\n'), byte(' '), byte('\t'): // ignore whitespace, tabs and new lines
			current_pos++
		default:
			literal := ""
			for current_pos < len(source) && isLetter(source[current_pos]) {
				literal += string(source[current_pos])
				current_pos++
			}
			tokens = append(tokens, NewToken(STRING, literal, "STRING"))
		}
	}

	return tokens
}

func isLetter(letter byte) bool {
	result := letter == '<' || letter == '>'
	return !result
}
