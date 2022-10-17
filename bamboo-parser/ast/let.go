package ast

import (
	"bamboo/token"
	"bytes"
)

// 如下定义了let语句的相关结构
// 形式: let <标识符> = <表达式>
// eg. let x = 5
// let x = 5 * 5
// let y = add(2,2)*5/10
// 定义3个字段: 关联词法单元 指向标志符号 指向等号右侧表达式

// LetStatement let语句结构
type LetStatement struct {
	Token token.Token // LET词法单元
	Name  *Identifier // 变量名
	Value Expression  // 等号右侧表达式
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")

	return out.String()
}
