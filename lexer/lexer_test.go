package lexer

import (
	"testing"

	"naahid/token"
)


func TestNextToken(t *testing.T) {
	input := `=+{},;`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	} {
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, "AAKHAR-E-KHAT"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestToken(t *testing.T) {
	input := `besaaz five = 5;`

		tests := []struct {
			expectedType token.TokenType
			expectedLiteral string
		}{
			{token.LET, "besaaz"},
			{token.IDENT, "five"},
			{token.ASSIGN, "="},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
		}
}
