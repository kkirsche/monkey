package token

const (
	// Lexer / Parser Control
	ILLEGAL = "ILLEGAL" // ILLEGAL is an illegal token or character which we don't know about
	EOF     = "EOF"     // EOF is the end of the file

	// Identifiers + literals
	IDENT = "IDENT" // IDENT is an identifer such as add, foobar, x, y, ...
	INT   = "INT"   // INT is an integer, such as 1343456

	// Operators
	ASSIGN   = "="  // ASSIGN: Assignment operation / Equal sign
	PLUS     = "+"  // PLUS: Plus sign
	MINUS    = "-"  // MINUS: Minus sign / Hyphen
	BANG     = "!"  // BANG: Bang / Exclamation Point
	ASTERISK = "*"  // ASTERISK: Asterisk
	SLASH    = "/"  // SLASH: Forward slash
	LT       = "<"  // LT: Less Than
	GT       = ">"  // GT: Greater Than
	EQ       = "==" // EQ: Equal operator
	NOT_EQ   = "!=" // NOT_EQ: Inequality operator

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
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
