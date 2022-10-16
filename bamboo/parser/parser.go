package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"strconv"
)

// 语法分析是解释器的第二个步骤
// 语法分析器使用由词法分析器生成的各个词法单元的第一个分量来创建树形表示
// 该中间表示给出了词法分析产生的词法单元流的语法结构

// 优先级顺序
const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
	INDEX
)

type Parser struct {
	lex    *lexer.Lexer // 词法分析器
	errors []string

	curToken  token.Token // 当前token 已经读到的token
	peekToken token.Token // 下一token 也就是将要读取的token

	prefixParseFns map[token.Type]prefixParseFn // 前缀解析函数关联表
	infixParseFns  map[token.Type]infixParseFn  // 后缀解析函数关联表
}

// 优先级表 优先级依次升高
var precedences = map[token.Type]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
	token.LBRACKET: INDEX,
}

// New 初始化一个语法分析器
func New(lexer *lexer.Lexer) *Parser {
	p := &Parser{
		lex:    lexer,
		errors: []string{},
	}

	// 注册即构建解析函数与对应tokenType的映射
	// 当遇到这一类token时 就去调用关联的解析函数

	p.prefixParseFns = make(map[token.Type]prefixParseFn)

	// 注册前缀表达式解析函数
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)

	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	p.infixParseFns = make(map[token.Type]infixParseFn)

	// 注册中缀表达式解析函数
	// 将每一个中缀运算符与parseInfixExpression函数相关联
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)

	// 注册布尔解析函数
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)

	// 注册分组表达式解析函数
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)

	// 注册if表达式解析函数
	p.registerPrefix(token.IF, p.parseIfExpression)

	// 注册函数解析函数
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)

	// 注册调用表达式解析函数
	p.registerInfix(token.LPAREN, p.parseCallExpression)

	// 注册字符串解析函数
	p.registerPrefix(token.STRING, p.parseStringLiteral)

	// 注册数组解析函数
	p.registerPrefix(token.LBRACKET, p.parseArrayLiteral)

	// 注册索引解析函数
	p.registerInfix(token.LBRACKET, p.parseIndexExpression)

	// 注册哈系表解析函数
	p.registerPrefix(token.LBRACE, p.parseHashLiteral)

	// 读取两个词法单元
	// 初始化curToken和peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

// 更新token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

/* 下面实现的是对源代码的语法分析
这一过程实际上是 扫描源代码 不断调用nextToken
反复前移词法单元指针并检查当前词法单元 以决定下一步操作
在扫描的过程中 读取到由词法分析器返回的token
将其交由parseStatement进行解析 得到子节点放入AST中
*/

// ParseProgram 构造AST根节点
// 随后调用其他函数来构建子节点
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// 遍历到EOF则终止
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken() // 更新token 推进语法分析
	}
	return program
}

// 根据当前token的type 交由相关函数进行解析
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// 为各种类型的token定义对应的解析函数
// 所有解析函数都将返回一个statement结构 作为AST上的节点

// 解析let语句
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// let后 跟随一个标识符
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// 标志符后是一个等于号
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// 等于号后是一个表达式 对表达式再做解析
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)

	// 遍历到分号之后
	for p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// 解析return语句
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()
	stmt.ReturnValue = p.parseExpression(LOWEST)

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// 解析整数
func (p *Parser) parseIntegerLiteral() ast.Expression {
	i := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	i.Value = value

	return i
}

// 解析布尔值
func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}

// 解析函数参数
func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	var identifiers []*ast.Identifier

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return identifiers
}

// 解析函数表达式
func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.curToken}

	// 缺失左括号 返回错误
	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	lit.Parameters = p.parseFunctionParameters() // 解析函数参数

	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	lit.Body = p.parseBlockStatement() // 解析函数体

	return lit
}

// 解析语句块
func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()
	// 循环 直到遇到右花括号或EOF
	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}

/* 在该语言中 除了let和return语句 皆是表达式
表达式可分为前缀表达式和后缀表达式
使用前缀运算符的表达式: -5 !true !false
使用中缀运算符的表达式: 1 + 1 1 - 1 1 == 1
括号能对表达式进行分组并影响求值顺序
下面实现对表达式的解析

采用的方法是自上而下的运算符优先级分析(普拉特解析法)
论文地址: https://eli.thegreenplace.net/2010/01/02/top-down-operator-precedence-parsing
普拉特解析法没有将解析函数与语法规则相关联
每种词法单元类型都可以具有两个与之相关联的解析函数
具体取决于词法单元的位置 比如是中缀或前缀

主要思想: 将解析函数(语义代码)与词法单元类型相关联
每当遇到某个词法单元类型时 调用相关联的解析函数解析对应表达式
最后返回AST节点 每个此法单元类型最多可以关联两个解析函数
取决于中缀位置还是前缀位置
*/

// 解析表达式 分为前缀表达式解析和后缀表达式解析
func (p *Parser) parseExpression(precedence int) ast.Expression {
	// 检查前缀位置是否有与Type相关联的解析函数
	// 有则调用相关解析函数 没有则抛出error
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix() // 调用关联解析函数

	// 重复执行 直到遇到优先级较低的词法单元
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp) // 调用关联解析函数
	}
	return leftExp
}

// 解析表达式语句
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST) // 解析表达式

	// 遇到分号则移动
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// 解析前缀表达式
func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken, // 前缀运算符
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX) // 填充右表达式

	return expression
}

// 解析中缀表达式
// 该函数要传入左表达式 填充Left字段
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}
	precedence := p.curPrecedence() // 当前词法单元优先级
	p.nextToken()
	expression.Right = p.parseExpression(precedence) // 填充右表达式

	return expression
}

// 解析调用表达式
func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseExpressionList(token.RPAREN)
	return exp
}

// 解析调用参数
func (p *Parser) parseCallArguments() []ast.Expression {
	var args []ast.Expression

	// 遇到右括号停止
	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpression(LOWEST)) // 获取第一个参数

	// 如果有逗号 说明列表中含有多个参数
	// 直至遇到最后一个逗号后 停止获取参数
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST)) // 将参数添加进列表
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return args
}

// 解析分组表达式
func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()
	exp := p.parseExpression(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

// 解析if表达式
func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.curToken}

	// 缺失左括号 返回错误
	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST) // 解析条件表达式

	// 右括号缺失 返回错误
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	// 左花括号缺失 返回错误
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	expression.Consequence = p.parseBlockStatement() // 填充if语句块

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		//左花括号缺失 返回错误
		if !p.expectPeek(token.LBRACE) {
			return nil
		}
		expression.Alternative = p.parseBlockStatement() // 填充else语句块
	}

	return expression
}

// 将当前词法单元及其字面量分别提供给Token和Value字段
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: p.curToken}
	array.Elements = p.parseExpressionList(token.RBRACKET)
	return array
}

// 解析参数列表
func (p *Parser) parseExpressionList(end token.Type) []ast.Expression {
	var list []ast.Expression

	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}

// 解析索引表达式
func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	exp := &ast.IndexExpression{Token: p.curToken, Left: left}

	p.nextToken()
	exp.Index = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}
	return exp
}

// 解析哈希字面量
func (p *Parser) parseHashLiteral() ast.Expression {
	hash := &ast.HashLiteral{Token: p.curToken}
	hash.Pairs = make(map[ast.Expression]ast.Expression) // 创建map结构

	// 遍历到右花括号为止
	for !p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		key := p.parseExpression(LOWEST)

		if !p.expectPeek(token.COLON) {
			return nil
		}
		p.nextToken()
		value := p.parseExpression(LOWEST)
		hash.Pairs[key] = value // 存储键值

		if !p.peekTokenIs(token.RBRACE) && !p.expectPeek(token.COMMA) {
			return nil
		}
	}
	if !p.expectPeek(token.RBRACE) {
		return nil
	}
	return hash
}
