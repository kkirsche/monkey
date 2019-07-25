package lexer

import (
	"testing"

	"github.com/kkirsche/monkey/token"
	"github.com/stretchr/testify/require"
)

type expected struct {
	// tType is the token type
	tType token.Type
	// tLiteral is the token literal
	tLiteral string
	// tColumn is the token's column number
	tColumn int
	// tLine is the token's line number
	tLine int
}

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
 x + y;
};

let result = add(five, ten);
`

	tests := []expected{
		expected{token.LET, "let", 1, 1},
		expected{token.IDENT, "five", 5, 1},
		expected{token.ASSIGN, "=", 10, 1},
		expected{token.INT, "5", 12, 1},
		expected{token.SEMICOLON, ";", 13, 1},
		expected{token.LET, "let", 1, 2},
		expected{token.IDENT, "ten", 5, 2},
		expected{token.ASSIGN, "=", 9, 2},
		expected{token.INT, "10", 11, 2},
		expected{token.SEMICOLON, ";", 13, 2},
		expected{token.LET, "let", 1, 4},
		expected{token.IDENT, "add", 5, 4},
		expected{token.ASSIGN, "=", 9, 4},
		expected{token.FUNCTION, "fn", 11, 4},
		expected{token.LPAREN, "(", 13, 4},
		expected{token.IDENT, "x", 14, 4},
		expected{token.COMMA, ",", 15, 4},
		expected{token.IDENT, "y", 17, 4},
		expected{token.RPAREN, ")", 18, 4},
		expected{token.LBRACE, "{", 20, 4},
		expected{token.IDENT, "x", 2, 5},
		expected{token.PLUS, "+", 4, 5},
		expected{token.IDENT, "y", 6, 5},
		expected{token.SEMICOLON, ";", 7, 5},
		expected{token.RBRACE, "}", 1, 6},
		expected{token.SEMICOLON, ";", 2, 6},
		expected{token.LET, "let", 1, 8},
		expected{token.IDENT, "result", 5, 8},
		expected{token.ASSIGN, "=", 12, 8},
		expected{token.IDENT, "add", 14, 8},
		expected{token.LPAREN, "(", 17, 8},
		expected{token.IDENT, "five", 18, 8},
		expected{token.COMMA, ",", 22, 8},
		expected{token.IDENT, "ten", 24, 8},
		expected{token.RPAREN, ")", 27, 8},
		expected{token.SEMICOLON, ";", 28, 8},
		expected{token.EOF, "", 0, 9},
	}

	lex := New(input)

	for _, tt := range tests {
		tok := lex.NextToken()

		require.Equal(t, tt.tType, tok.Type, "Invalid token type '%s', expected '%s'", tok.Type, tt.tType)
		require.Equal(t, tt.tLiteral, tok.Literal, "Invalid token literal '%s', expected '%s'", tok.Literal, tt.tLiteral)
		require.Equal(t, tt.tColumn, tok.Column, "Invalid column number %d for token literal '%s'", tok.Column, tok.Literal)
		require.Equal(t, tt.tLine, tok.Line, "Invalid line number %d for token literal '%s'", tok.Line, tok.Literal)
	}
}
