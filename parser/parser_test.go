package parser

import (
	"testing"

	"github.com/kkirsche/monkey/ast"
	"github.com/kkirsche/monkey/lexer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type expected struct {
	identifier string
}

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	require.NotNil(t, program, "ParseProgram() returned nil")
	require.Lenf(t, program.Statements, 3, "program.Statements does not contain 3 statements. got=%d", len(program.Statements))

	tests := []expected{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.identifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if !assert.Equalf(t, "let", s.TokenLiteral(), "s.TokenLiteral not 'let'. got=%q", s.TokenLiteral()) {
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if !assert.Equalf(t, name, letStmt.Name.Value, "letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value) {
		return false
	}

	if !assert.Equalf(t, name, letStmt.Name.TokenLiteral(), "letStmt.Name.TokenLiteral() not '%s'. got=%s", name, letStmt.Name.TokenLiteral()) {
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if assert.Len(t, errors, 0) {
		return
	}

	t.Errorf("parser had %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 993322;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.NotNil(t, program, "ParseProgram() returned nil")
	require.Lenf(t, program.Statements, 3, "program.Statements does not contain 3 statements. got=%d", len(program.Statements))

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}

		assert.Equalf(t, "return", returnStmt.TokenLiteral(), "returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
	}
}
