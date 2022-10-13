package parser

import (
	"fmt"
	"monkey/token"
)

func (p *Parser) peekError(t token.Type) {
	msg := fmt.Sprintf("expected next token "+
		"to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) noPrefixParseFnError(t token.Type) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}
