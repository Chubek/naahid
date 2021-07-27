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
	MINUS  = "-"
	BANG   = "!"
	ASTERISK = "*"
	SLASH	= "/"
	LT		= "<"
	GT		= ">"

	EQ 		= "=="
	NOT_EQ	= "!="

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "TAABEH"
	LET      = "BESAAZ"
	TRUE	= "RAAST"
	FALSE	= "QALAT"
	IF		= "AGAR"
	ELSE	= "VAGAR"
	RETURN	= "BARGARDAAN"
)

var keywords = map[string]TokenType {
	"taab": FUNCTION,
	"besaaz": LET,
	"raast": TRUE,
	"qalat": FALSE,
	"agar": IF,
	"vagar": ELSE,
	"bargardaan": RETURN,  
}


func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}