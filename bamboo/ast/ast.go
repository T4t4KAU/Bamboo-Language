package ast

import (
	"bytes"
	"monkey/token"
)

// AST = Abstract Syntax Tree 抽象语法树

// Node 节点类型
// AST中的每一个节点都必须实现Node接口
type Node interface {
	TokenLiteral() string // 返回关联的词法单元的字面量
	String() string
}

// 下述定义了语句和表达式的结构
// 其中的statementNode和expressionNode方法仅仅是占位方法

// Statement 语句结构
type Statement interface {
	Node
	statementNode()
}

// Expression 表达式结构
type Expression interface {
	Node
	expressionNode()
}

// ExpressionStatement 表达式构成的语句
type ExpressionStatement struct {
	Token      token.Token // 表达式中第一个词法单元
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// Program AST根节点
type Program struct {
	Statements []Statement
}

// TokenLiteral 返回与其关联的词法单元的字面量
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// Identifier 标识符号
type Identifier struct {
	Token token.Token // 词法单元
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

// BlockStatement 块语句结构
// 该结构表示一个语句块 包含多条语句
type BlockStatement struct {
	Token      token.Token // '{'词法单元
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
