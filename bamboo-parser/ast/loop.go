package ast

import (
	"bamboo/token"
	"bytes"
)

// 循环结构: while (condition) { expression }

type WhileExpression struct {
	Token     token.Token     // ‘while' 词法单元
	Condition Expression      // 条件语句
	Body      *BlockStatement // 循环体
}

func (we *WhileExpression) expressionNode() {}
func (we *WhileExpression) TokenLiteral() string {
	return we.Token.Literal
}
func (we *WhileExpression) String() string {
	var out bytes.Buffer

	out.WriteString("while")
	out.WriteString(we.Condition.String())
	out.WriteString(we.Body.String())

	return out.String()
}
