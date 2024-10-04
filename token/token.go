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

type TokenType string

type Token struct {
	// Type variable holds the type of token
	Type TokenType
	// Literal variable holds the literal value of token
	Literal string
}