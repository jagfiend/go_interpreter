package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	// special
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	
	// identifiers and literals
	IDENT = "IDENT"
	INT = "INT"
	
	// operators
	ASSIGN = "="
	PLUS = "+"
	
	// delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET = "FUNCTION"
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
