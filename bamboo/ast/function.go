package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

// 函数字面量

/* 函数是一个表达式
格式: fn <参数列表> <块语句>
参数列表: (<参数1>,<参数2>,<参数3>,...)
eg. fn(x, y) { return x + y; }
let func = fn(x, y) { return x + y; }
*/

type FunctionLiteral struct {
	Token      token.Token     //'fn'词法单元
	Parameters []*Identifier   // 参数列表
	Body       *BlockStatement // 函数体 语句块
}

func (fl *FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	var params []string
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

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
