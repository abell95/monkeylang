package token

// Type is the token type, represented as a string.
type Type string

// Token is the data structure representing a token in the lexer
type Token struct {
	Type    Type
	Literal string
}

const (
	// ILLEGAL : Illegal character
	ILLEGAL = "ILLEGAL"
	// EOF : End of file
	EOF = "EOF"

	// IDENT : Identifier, i.e. variable or func name
	IDENT = "IDENT"
	// INT : A whole number
	INT = "INT" // whole number

	// ASSIGN : An equals sign
	ASSIGN = "="
	// PLUS : A plus sign
	PLUS = "+"

	// COMMA : Argument separator
	COMMA = ","
	// SEMICOLON : End statement
	SEMICOLON = ";"

	// LPAREN : Start paren
	LPAREN = "("
	// RPAREN : End paren
	RPAREN = ")"
	// LBRACE : Start brace
	LBRACE = "{"
	// RBRACE : End brace
	RBRACE = "}"

	// FUNCTION : Function keyword
	FUNCTION = "FUNCTION"
	// LET : Instantiation keyword
	LET = "LET"
)
