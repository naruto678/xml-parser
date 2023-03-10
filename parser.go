package main

import "fmt"

type Node struct {
	Name      string
	Parent    *Node
	Children  []*Node
	TextValue string
}

type Parser struct {
	Idx    int
	Tokens []Token
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		Idx:    0,
		Tokens: tokens,
	}

}

func NewNode(Name string, Parent Node) Node {
	return Node{
		Name:     Name,
		Parent:   &Parent,
		Children: []*Node{},
	}
}

func (p *Parser) Unmarshall() *Node {
	return nil
}

func (p *Parser) Marshall() (string, error) {
	return "", fmt.Errorf("not implemented yet")
}

func (p *Parser) Increment() {
	p.Idx++
}

func (p *Parser) Peek() Token {
	return p.Tokens[p.Idx]
}

func (p *Parser) GetCurrentToken() Token {
	token_val := p.GetToken(p.Idx)
	p.Increment()
	return token_val
}

func (p *Parser) PeekNext() Token {
	if p.Idx+1 >= len(p.Tokens) {
		return EOF()
	} else {
		return p.Tokens[p.Idx+1]
	}
}

func (p *Parser) GetToken(idx int) Token {
	if p.Idx > len(p.Tokens) {
		return EOF()
	}
	return p.Tokens[p.Idx]
}

func (p *Parser) Parse() (Node, error) {
	parent_node := Node{}
	return p.Node(parent_node)
}
func (p *Parser) Node(parent Node) (Node, error) {
	var curr_node Node
	if p.GetCurrentToken().TokenType == LPAREN {
		if p.Peek().TokenType == STRING {
			curr_node = NewNode(p.GetCurrentToken().Literal, parent)
			if curr_token := p.GetCurrentToken(); curr_token.TokenType != RPAREN {
				return curr_node, fmt.Errorf("expected '>' Got %s", p.Peek().Name)
			}

			next_token := p.Peek()

			for next_token.TokenType == LPAREN && p.PeekNext().TokenType != SLASH {
				child_node, err := p.Node(curr_node)
				if err != nil {
					return curr_node, err
				}
				curr_node.Children = append(curr_node.Children, &child_node)
				next_token = p.Peek()

			}

			if next_token.TokenType == STRING {
				curr_node.TextValue = next_token.Literal
				p.GetCurrentToken()
			}

			if p.GetCurrentToken().TokenType == LPAREN && p.GetCurrentToken().TokenType == SLASH {
				closing_symbol_name := p.GetCurrentToken()
				if curr_node.Name != closing_symbol_name.Literal {
					return curr_node, fmt.Errorf("closing tag does not match . expected [%s] got [%s]", curr_node.Name, closing_symbol_name.Literal)
				}
				if p.GetCurrentToken().TokenType != RPAREN {
					return curr_node, fmt.Errorf("expected > got [%s]", p.Peek().Literal)
				}
			}

			return curr_node, nil
		}
	}
	return curr_node, fmt.Errorf("expected < . got %s", p.Peek().Name)
}

func (p *Parser) Print(printer Printer) {
	printer.Print()
}
