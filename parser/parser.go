package parser

import (
	"fmt"

	"github.com/kkirsche/monkey/ast"
	"github.com/kkirsche/monkey/lexer"
	"github.com/kkirsche/monkey/token"
)

// Parser is the structure responsible for reading tokens from the lexer and
// generating the appropriate abstract syntax tree based on the read tokens.
type Parser struct {
	l *lexer.Lexer

	errors    []string
	curToken  token.Token
	peekToken token.Token
}

// New creates a new Monkey programming language parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

// Errors is a getter method allowing clients to read the parser errors
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram parses the input provided to the lexer into an abstract syntax
// tree
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// parseStatement is used to parse each statement within the input. In the case
// of an unsupported or unknown token, ignore it by returning nil. This allows
// the parser to "eat" unknown or currently unsupported inputs
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// parseLetStatement is called when a let token has been found during the
// parseStatement branch decision
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// enforce that the next token is an identifier
	// let identifier = expression;
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	// enforce that the next token is an assignment operator
	// let identifier = expression;
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// curTokenIs is used to check that the current token is what we require it to
// be to continue parsing, or returns false
func (p *Parser) curTokenIs(t token.Type) bool {
	return p.curToken.Type == t
}

// peekTokenIs is used to check that the peeked token is what we require it to
// be to continue parsing, or returns false
func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func (p *Parser) peekError(t token.Type) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// expectPeek is used to enforce what the expected next token is. This allows
// us to enforce the correctness of the order of the received tokens by checking the next token.
func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}
