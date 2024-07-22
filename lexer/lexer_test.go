package lexer

import (
	"netherlang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `maak vijf = 5;
    maak tien = 10;

    maak telop = functie(x, y) {
        geef x + y;
    };

    maak resultaat = telop(vijf, tien);
    !-/*5;
    5 < 10 > 5;

    als anders waar onwaar;
    !=;

    10 == 10;
    10 != 5;
    10 <= 5;
    10 >= 5;
    `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
        {token.LET, "maak"},
        {token.IDENT, "vijf"},
        {token.ASSIGN, "="},
        {token.INT, "5"},
        {token.SEMICOLON, ";"},
        {token.LET, "maak"},
        {token.IDENT, "tien"},
        {token.ASSIGN, "="},
        {token.INT, "10"},
        {token.SEMICOLON, ";"},
        {token.LET, "maak"},
        {token.IDENT, "telop"},
        {token.ASSIGN, "="},
        {token.FUNCTION, "functie"},
        {token.LPAREN, "("},
        {token.IDENT, "x"},
        {token.COMMA, ","},
        {token.IDENT, "y"},
        {token.RPAREN, ")"},
        {token.LBRACE, "{"},
        {token.RETURN, "geef"},
        {token.IDENT, "x"},
        {token.PLUS, "+"},
        {token.IDENT, "y"},
        {token.SEMICOLON, ";"},
        {token.RBRACE, "}"},
        {token.SEMICOLON, ";"},
        {token.LET, "maak"},
        {token.IDENT, "resultaat"},
        {token.ASSIGN, "="},
        {token.IDENT, "telop"},
        {token.LPAREN, "("},
        {token.IDENT, "vijf"},
        {token.COMMA, ","},
        {token.IDENT, "tien"},
        {token.RPAREN, ")"},
        {token.SEMICOLON, ";"},
        {token.BANG, "!"},
        {token.MINUS, "-"},
        {token.SLASH, "/"},
        {token.ASTERISK, "*"},
        {token.INT, "5"},
        {token.SEMICOLON, ";"},
        {token.INT, "5"},
        {token.LT, "<"},
        {token.INT, "10"},
        {token.GT, ">"},
        {token.INT, "5"},
        {token.SEMICOLON, ";"},
        {token.IF, "als"},
        {token.ELSE, "anders"},
        {token.TRUE, "waar"},
        {token.FALSE, "onwaar"},
        {token.SEMICOLON, ";"},
        {token.NEQ, "!="},
        {token.SEMICOLON, ";"},
        {token.INT, "10"},
        {token.EQ, "=="},
        {token.INT, "10"},
        {token.SEMICOLON, ";"},
        {token.INT, "10"},
        {token.NEQ, "!="},
        {token.INT, "5"},
        {token.SEMICOLON, ";"},
        {token.INT, "10"},
        {token.LTE, "<="},
        {token.INT, "5"},
        {token.SEMICOLON, ";"},
        {token.INT, "10"},
        {token.GTE, ">="},
        {token.INT, "5"},
        {token.SEMICOLON, ";"},

	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
                i, tt.expectedLiteral, tok.Literal)
        }
	}

}
