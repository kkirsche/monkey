package token

const (
	// ILLEGAL is an illegal token or character which we don't know about
	ILLEGAL = "ILLEGAL"
	// EOF is the end of the file
	EOF = "EOF"

	// Identifiers + literals

	// IDENT is an identifer such as add, foobar, x, y, ...
	IDENT = "IDENT"
	// INT is an integer, such as 1343456
	INT = "INT"

	// Operators

	// ASSIGN =
	ASSIGN = "="
	// PLUS +
	PLUS = "+"

	// Delimiters

	// COMMA ,
	COMMA = ","
	// SEMICOLON ;
	SEMICOLON = ";"
	// LPAREN (
	LPAREN = "("
	// RPAREN )
	RPAREN = ")"
	// LBRACE {
	LBRACE = "{"
	// RBRACE }
	RBRACE = "}"

	// FUNCTION FUNCTION
	FUNCTION = "FUNCTION"
	// LET LET
	LET = "LET"
)

var keywords = map[string]Type{
	"fn":  FUNCTION,
	"let": LET,
}

// Type is a string which allows us to distinguish between types of tokens
// Strings are easy to debug without a lot of boilerplate, though are not as
// performant as something like an iota enum
type Type string

// Token represents an emitted token from the lexer containing both it's type
// and the literal value of the token
type Token struct {
	Type    Type
	Literal string
	Column  int
	Line    int
}

// LookupIdent is used to check if the identifier we've read is a keyword
// if it is not, we return a generic IDENT
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
