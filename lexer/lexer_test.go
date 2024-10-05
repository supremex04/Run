package lexer

import (
	"testing"
	"run/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	// defines array of anonymous structure,
	// each struct has two fields expectedType and expectedLiteral
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{	//expectedType   expectedLiteral
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			// %q is used to print string inside " "
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
