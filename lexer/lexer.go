package lexer

import (
	"github.com/kkirsche/monkey/token"
)

// Lexer is the structure responsible for converting the input text into a
// series of tokens
type Lexer struct {
	input        string
	inputLen     int
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // churrent char under examination
	column       int  // the current number of the column of the line
	line         int  // the current number of line
}

// New is used to create a new lexer instance from the input text
func New(input string) *Lexer {
	l := &Lexer{
		input:    input,
		inputLen: len(input),
		line:     1,
	}
	// this ensures that our position, readPosition, and column number are
	// all initialized before the caller uses the lexer
	l.readChar()
	return l
}

func newToken(tType token.Type, ch byte, column, line int) token.Token {
	return token.Token{Type: tType, Literal: string(ch), Column: column, Line: line}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// readChar advances our position in the input string and provides us with the
// next character. If we have reached the end of the string, we set the
// character to 0, which is the ASCII "NUL" code.
func (l *Lexer) readChar() {
	if l.ch == '\n' {
		l.line++
		l.column = 0
	}

	if l.readPosition >= l.inputLen {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	l.column++
}

// peekChar is similar to readChar, but instead we peek ahead at the next
// character in the input stream rather than actually advancing forward.
// This allows for us to look for two character tokens more easily.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= l.inputLen {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

// skipWhitespace is used to skip over general whitespace characters
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// NextToken is used to read from the input stream and identify what the next
// token is
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			// get the first equal
			ch := l.ch
			// get the second equal and advance our position
			l.readChar()
			// construct the == literal
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal, Column: l.column - 1, Line: l.line}
		} else {
			tok = newToken(token.ASSIGN, l.ch, l.column, l.line)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch, l.column, l.line)
	case '-':
		tok = newToken(token.MINUS, l.ch, l.column, l.line)
	case '!':
		if l.peekChar() == '=' {
			// get the first bang
			ch := l.ch
			// get the second equal and advance our position
			l.readChar()
			// construct the != literal
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal, Column: l.column - 1, Line: l.line}
		} else {
			tok = newToken(token.BANG, l.ch, l.column, l.line)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch, l.column, l.line)
	case '/':
		tok = newToken(token.SLASH, l.ch, l.column, l.line)
	case '<':
		tok = newToken(token.LT, l.ch, l.column, l.line)
	case '>':
		tok = newToken(token.GT, l.ch, l.column, l.line)
	case ',':
		tok = newToken(token.COMMA, l.ch, l.column, l.line)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, l.column, l.line)
	case '(':
		tok = newToken(token.LPAREN, l.ch, l.column, l.line)
	case ')':
		tok = newToken(token.RPAREN, l.ch, l.column, l.line)
	case '{':
		tok = newToken(token.LBRACE, l.ch, l.column, l.line)
	case '}':
		tok = newToken(token.RBRACE, l.ch, l.column, l.line)
	case 0:
		// EOF case
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Column = l.column - 1
		tok.Line = l.line
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			tok.Column = l.column - len(tok.Literal)
			tok.Line = l.line
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			tok.Column = l.column - len(tok.Literal)
			tok.Line = l.line
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ch, l.column-1, l.line)
	}

	l.readChar()
	return tok
}
