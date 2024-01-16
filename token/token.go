package token

import "fmt"

type TokenType string

const (
	TokenTypeIllegal    TokenType = "ILLEGAL"
	TokenTypeEOF        TokenType = "EOF"
	TokenTypeIdentifier TokenType = "IDENT"
	TokenTypeInteger    TokenType = "INT"
	TokenTypeAssign     TokenType = "ASSIGN"
	TokenTypePlus       TokenType = "PLUS"
	TokenTypeMinus      TokenType = "MINUS"
	TokenTypeAsterisk   TokenType = "ASTERISK"
	TokenTypeComma      TokenType = "COMMA"
	TokenTypeSemicolon  TokenType = "SEMICOLON"
	TokenTypeLParen     TokenType = "LPAREN"
	TokenTypeRParen     TokenType = "RPAREN"
	TokenTypeLBrace     TokenType = "LBRACE"
	TokenTypeRBrace     TokenType = "RBRACE"
	TokenTypeFunction   TokenType = "FUNC"
	TokenTypeLet        TokenType = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}

func (token Token) String() string {
	return fmt.Sprintf("{%s:\"%s\"}", token.Type, token.Literal)
}
