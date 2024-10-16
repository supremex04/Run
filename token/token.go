package token

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT"
	// Operators
	ASSIGN= "="
	PLUS = "+"
	// Delimiters
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
	)

// defining TokenType as a  new type that is based on string type
type TokenType string

type Token struct {
	// Type variable holds the type of token
	Type TokenType
	// Literal variable holds the literal value of token
	Literal string
}

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		// returns the keyword's TokenType like FUNCTION, LET
		return tok
	}
	// returning IDENT indicates it's just a regular identifier
	return IDENT
}