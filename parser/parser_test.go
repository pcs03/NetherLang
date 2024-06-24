package parser

import (
    "testing"
    "netherlang/ast"
    "netherlang/lexer"
)

func TestLetStatement(t *testing.T) {
    input := `
    maak x = 5;
    maak y = 10;
    maak foebar = 83838;
    `

    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()
    checkParserErrors(t, p)

    if program == nil {
        t.Fatalf("ParseProgram() returned nil")
    }

    if len(program.Statements) != 3 {
        t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
    }

    tests := []struct {
        expectedIdentifier string
    }{
        {"x"},
        {"y"},
        {"foebar"},
    }

    for i, tt := range tests {
        statement := program.Statements[i]

        if !testLetStatement(t, statement, tt.expectedIdentifier) {
            return
        }
    }
}

func checkParserErrors(t *testing.T, p *Parser) {
    errors := p.Errors()

    if len(errors) == 0 {
        return
    }

    t.Errorf("parser has %d errors", len(errors))
    for _, msg := range errors {
        t.Errorf("parser error: %q", msg)
    }
    t.FailNow()
}

func testLetStatement(t *testing.T, statement ast.Statement, name string) bool {
    if statement.TokenLiteral() != "maak" {
        t.Errorf("statement.TokenLiteral not 'maak'. got=%q", statement.TokenLiteral())
        return false
    }

    letStatement, ok := statement.(*ast.LetStatement)

    if !ok {
        t.Errorf("s not *ast.LetStatement, got =%T", statement)
        return false
    }

    if letStatement.Name.Value != name {
        t.Errorf("letStatement.Name.Value not '%s', got=%s", name, letStatement.Name.Value)
        return false
    }

    if letStatement.Name.TokenLiteral() != name {
        t.Errorf("letStatement.Name.TokenLiteral() not '%s'. got=%s", name, letStatement.Name.TokenLiteral())
        return false
    }

    return true
}
