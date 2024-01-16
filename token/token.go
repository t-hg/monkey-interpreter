package token

import "fmt"

type TokenType string

const (
	TokenTypeIllegal     TokenType = "ILLEGAL"
	TokenTypeEOF         TokenType = "EOF"
	TokenTypeIdentifier  TokenType = "IDENT"
	TokenTypeInteger     TokenType = "INT"
	TokenTypeAssign      TokenType = "ASSIGN"
	TokenTypePlus        TokenType = "PLUS"
	TokenTypeMinus       TokenType = "MINUS"
	TokenTypeAsterisk    TokenType = "ASTERISK"
	TokenTypeSlash       TokenType = "SLASH"
	TokenTypeLessThan    TokenType = "LT"
	TokenTypeGreaterThan TokenType = "GT"
	TokenTypeBang        TokenType = "BANG"
	TokenTypeComma       TokenType = "COMMA"
	TokenTypeSemicolon   TokenType = "SEMICOLON"
	TokenTypeLParen      TokenType = "LPAREN"
	TokenTypeRParen      TokenType = "RPAREN"
	TokenTypeLBrace      TokenType = "LBRACE"
	TokenTypeRBrace      TokenType = "RBRACE"
	TokenTypeFunction    TokenType = "FUNC"
	TokenTypeLet         TokenType = "LET"
	TokenTypeTrue        TokenType = "TRUE"
	TokenTypeFalse       TokenType = "FALSE"
	TokenTypeIf          TokenType = "IF"
	TokenTypeElse        TokenType = "ELSE"
	TokenTypeReturn      TokenType = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
}

func (token Token) String() string {
	return fmt.Sprintf("{%s:\"%s\"}", token.Type, token.Literal)
}
