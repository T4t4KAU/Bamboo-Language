package ast

import (
	"bamboo/token"
	"bytes"
)

// 前缀表达式

// 前缀运算符是位于操作数前面的运算符 eg. --5

type PrefixExpression struct {
	Token    token.Token // 前缀词法单元
	Operator string      // 操作符
	Right    Expression  // 右侧表达式
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
