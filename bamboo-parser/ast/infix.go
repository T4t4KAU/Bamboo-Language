package ast

import (
	"bamboo/token"
	"bytes"
)

// <表达式> <中缀运算符> <表达式>

// InfixExpression 中缀表达式结构
type InfixExpression struct {
	Token    token.Token // 运算符词法单元
	Left     Expression  // 左表达式
	Operator string      // 操作符
	Right    Expression  // 右表达式
}

func (ie *InfixExpression) expressionNode() {}
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}
