package parser

import (
	"bamboo/ast"
	"bamboo/token"
)

// 判断当前token是否为指定的词法单元
func (p *Parser) curTokenIs(t token.Type) bool {
	return p.curToken.Type == t
}

// 判断下一个token是否为期望的词法单元
func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

// 判断下一个token是否为指定类型 并作移动
func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// 查看下一token优先级
func (p *Parser) peekPrecedence() int {
	if precedence, ok := precedences[p.peekToken.Type]; ok {
		return precedence
	}
	return LOWEST
}

// 查看当前token优先级
func (p *Parser) curPrecedence() int {
	if precedence, ok := precedences[p.curToken.Type]; ok {
		return precedence
	}
	return LOWEST
}

type (
	prefixParseFn func() ast.Expression               // 前缀解析函数
	infixParseFn  func(ast.Expression) ast.Expression // 中缀解析函数
)

// 注册前缀解析函数
func (p *Parser) registerPrefix(tokenType token.Type, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// 注册中缀解析函数
func (p *Parser) registerInfix(tokenType token.Type, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
