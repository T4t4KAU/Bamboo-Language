package ast

import (
	"bytes"
	"monkey/token"
)

// if条件语句
/* if-else条件语句是表达式
格式: if (<条件>) <结果> else <可替代的结果>
eg. if (x < y) { x } else { y }
结构可以可以划分为4个字段:
if词法单元 条件表达式 if执行的语句块 else执行的语句块
*/

type IfExpression struct {
	Token       token.Token // if词法单元
	Condition   Expression  // 条件表达式
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString("ie.Condition.String()")
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}
