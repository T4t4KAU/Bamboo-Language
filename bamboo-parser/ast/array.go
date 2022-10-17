package ast

import (
	"bamboo/token"
	"bytes"
	"strings"
)

// 数组是含有多个元素的有序列表
// 其中的元素可以不相同 数组中的每个元素都可以单独访问
// 数组使用字面量构建 以逗号分隔列表中的元素 使用[]包裹

type ArrayLiteral struct {
	Token    token.Token  // '['词法单元
	Elements []Expression // 数组元素
}

func (al *ArrayLiteral) expressionNode() {}
func (al *ArrayLiteral) TokenLiteral() string {
	return al.Token.Literal
}
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	var elements []string
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
