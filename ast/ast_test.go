package ast

import (
    "netherlang/token"
    "testing"
)

func TestString(t *testing.T) {
    program := &Program{
        Statements: []Statement{
            &LetStatement{
                Token: token.Token{Type: token.LET, Literal: "maak"},
                Name: &Identifier{
                    Token: token.Token{Type: token.IDENT, Literal: "variabele"},
                    Value: "variabele",
                },
                Value: &Identifier{
                    Token: token.Token{Type: token.IDENT, Literal: "var"},
                    Value: "var",
                },
            },
        },
    }

    if program.String() != "maak variabele = var;" {
        t.Errorf("program.String() wrong, got=%q", program.String())
    }
}
