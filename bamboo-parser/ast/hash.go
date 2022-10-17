package ast

import (
	"bamboo/token"
	"bytes"
	"strings"
)

// 哈希表是一个数据类型 使用花括号括起来的键-值对列表
// 列表中的键-值对用逗号分割
// eg. let h = {"name":"Jimmy","age":72}  h["name"]
// 字符串 整数 布尔值等任何表达式都可以用作索引运算符表达式索引

type HashLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression // 使用Go内置的map作为基础数据结构
}

func (hl *HashLiteral) expressionNode() {}
func (hl *HashLiteral) TokenLiteral() string {
	return hl.Token.Literal
}
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	var pairs []string
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
