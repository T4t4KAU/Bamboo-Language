package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

// 调用表达式

/* 解析函数的调用
格式: <表达式>(<以逗号分割的表达式列表>)
eg. add(2,3) add(2+2,3*3*3)

*/

type CallExpression struct {
	Token     token.Token // '('词法单元
	Function  Expression  // 标识符或函数字面量
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) TokenLiteral() string {
	return ce.TokenLiteral()
}
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	var args []string
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
