package ast

import (
	"github.com/kkirsche/monkey/token"
)

// Node is an individual node or chunk of the abstract syntax tree. If the tree
// is the entire program, the node is a single expression. All nodes are
// connected to each other, constucting a tree structure
type Node interface {
	TokenLiteral() string
}

// Statement is a specific type of node, and represents each "statement" or
// syntactic unit of the programming language. It expresses some action to
// be carried out.
type Statement interface {
	Node
	statementNode()
}

// Expression is a specific type of node, and represents the internal
// components of a statement
type Expression interface {
	Node
	expressionNode()
}

// Program is the root of our AST
type Program struct {
	Statements []Statement
}

// TokenLiteral implements the Node interface for the Program structure
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// LetStatement is a node type which handles the `let x = 5 *5` type of
// binding expression. It's composed of three pieces, the Token, Name and Value
// which allows us to keep track of the token literal (via Token), the
// expression which produces the value (Value), and the name of what the
// expression is being bound to
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // the identifier of the binding
	Value *Expression // The expression that produces the value
}

// statementNode implements the Statement interface
func (ls *LetStatement) statementNode() {}

// TokenLiteral implements the Node interface
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// ReturnStatement is how a function can return a value to it's caller. This
// follows structure `return <expression>`. Thus consist solely of a keyword,
// 'return', and the expression
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue *Expression
}

// statementNode implements the Statement interface
func (rs *ReturnStatement) statementNode() {}

// TokenLiteral implements the Node interface
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// Identifier is the individual identifier which represents the expression
// While not all statements have a value for their identifier, some do, and as
// such this structure allows us to reuse the identifier for different
// statements which do include one. This helps key the total number of node
// types small(er)
type Identifier struct {
	Token token.Token // the token.IDENT
	Value string
}

// expressionNode implements the Expression interface
func (i *Identifier) expressionNode() {}

// TokenLiteral implements the Node interface
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
