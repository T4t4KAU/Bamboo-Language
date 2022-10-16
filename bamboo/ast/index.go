package ast

import (
	"bytes"
	"monkey/token"
)

// 索引运算: <表达式>[<表达式>]

type IndexExpression struct {
	Token token.Token
	Left  Expression // 正在访问的对象
	Index Expression // 产生一个整数的表达式
}

func (ie *IndexExpression) expressionNode() {}
func (ie *IndexExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}
