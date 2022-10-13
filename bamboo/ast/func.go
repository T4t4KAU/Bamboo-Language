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
