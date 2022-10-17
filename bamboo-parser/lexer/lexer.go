package lexer

import "bamboo/token"

// 词法分析是解释器要做的第一件事 lexical analysis
// 词法分析器读入组成源程序的字符流 将其组织成有意义的lexeme序列
// 对于每个lexeme 分析器将产生如下形式的token: <token-type,token-Literal>
// 所产生的token被传给下一步骤: 语法分析
// 同时 该步骤要过滤源程序中的注释和空白 将错误消息与源程序的位置联系起来

type Lexer struct {
	input        string
	position     int  // 输入字符串的当前位置
	readPosition int  // 当前字符下一个字符
	ch           byte // 当前字符
}

// New 创建词法分析器
func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

// 读取input下一个字符 并前移在input中的位置
func (lexer *Lexer) readChar() {
	// 检查是否已经到达input末尾
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}
	// 更新位置 readPosition始终指向下一个将读取的字符位置
	// position始终指向刚刚读取的位置
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

// NextToken 检查当前正在查看的字符
// 根据具体的字符来返回对应的词法单元
func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token
	lexer.skipWhiteSpace()

	switch lexer.ch {
	case '=':
		if lexer.peekChar() == '=' {
			ch := lexer.ch
			lexer.readChar()
			literal := string(ch) + string(lexer.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, lexer.ch)
		}
	case '+':
		tok = newToken(token.PLUS, lexer.ch)
	case '-':
		tok = newToken(token.MINUS, lexer.ch)
	case '!':
		if lexer.peekChar() == '=' {
			ch := lexer.ch
			lexer.readChar()
			literal := string(ch) + string(lexer.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, lexer.ch)
		}
	case '/':
		tok = newToken(token.SLASH, lexer.ch)
	case '*':
		tok = newToken(token.ASTERISK, lexer.ch)
	case '<':
		tok = newToken(token.LT, lexer.ch)
	case '>':
		tok = newToken(token.GT, lexer.ch)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.ch)
	case '(':
		tok = newToken(token.LPAREN, lexer.ch)
	case ')':
		tok = newToken(token.RPAREN, lexer.ch)
	case ',':
		tok = newToken(token.COMMA, lexer.ch)
	case '{':
		tok = newToken(token.LBRACE, lexer.ch)
	case '}':
		tok = newToken(token.RBRACE, lexer.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = lexer.readString()
	case '[':
		tok = newToken(token.LBRACKET, lexer.ch)
	case ']':
		tok = newToken(token.RBRACKET, lexer.ch)
	case ':':
		tok = newToken(token.COLON, lexer.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default: // 检查是否为标识符
		if isLetter(lexer.ch) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(lexer.ch) {
			tok.Type = token.INT
			tok.Literal = lexer.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.ch)
		}
	}
	lexer.readChar()
	return tok
}

func newToken(tokenType token.Type, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// 读入一个标识符并前移词法分析器的扫描位置
func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	// 遇见非字母字符时停止
	for isLetter(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

// 跳过空白字符
func (lexer *Lexer) skipWhiteSpace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}

// 读取整数
func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

// 读取字符串
func (lexer *Lexer) readString() string {
	position := lexer.position + 1
	for {
		lexer.readChar()
		// 双引号和0为字符串的末尾
		if lexer.ch == '"' || lexer.ch == 0 {
			break
		}
	}
	return lexer.input[position:lexer.position]
}

// 读取下一个字符但不前移
func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
}

// 判断给定的参数是否为字母
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// 判断给定参数是否为数字
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// 由上述定义的程序 解析一条语句:
// let x = 1  ---> <LET,"let"> <IDENT,"x"> <ASSIGN,"="> <INT,1>
