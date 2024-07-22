package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"

	GT   = ">"
	GTE  = ">="
	LT   = "<"
	LTE  = "<="
	BANG = "!"

	EQ  = "=="
	NEQ = "!="

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTIE"
	RETURN   = "GEEF"
	LET      = "MAAK"
	TRUE     = "WAAR"
	FALSE    = "ONWAAR"
	IF       = "ALS"
	ELSE     = "ANDERS"
)

var keywords = map[string]TokenType{
	"functie":    FUNCTION,
	"maak":       LET,
	"geef": RETURN,
	"waar":       TRUE,
	"onwaar":     FALSE,
	"als":        IF,
	"anders":     ELSE,
}

func LookupIdent(ident string) TokenType {
    tok, ok := keywords[ident]

    if ok {
        return tok
    } else {
        return IDENT
    }
}
