package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "NAA-MALMUS"
	EOF     = "AAKHAR-E-KHAT"

	IDENT = "IDENT"
	INT  = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "TAABEH"
	LET      = "BESAAZ"
)

var keywords = map[string]TokenType {
	"taab": FUNCTION,
	"besaaz": LET,
}


func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}