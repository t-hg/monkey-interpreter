package token

import "fmt"

type TokenType string

const (
	TokenTypeIllegal     TokenType = "ILLEGAL"
	TokenTypeEOF         TokenType = "EOF"
	TokenTypeIdentifier  TokenType = "IDENTIFIER"
	TokenTypeInteger     TokenType = "INTEGER"
	TokenTypeAssign      TokenType = "ASSIGN"
	TokenTypeEquals      TokenType = "EQUALS"
	TokenTypeNotEquals   TokenType = "NOT_EQUALS"
	TokenTypePlus        TokenType = "PLUS"
	TokenTypeMinus       TokenType = "MINUS"
	TokenTypeAsterisk    TokenType = "ASTERISK"
	TokenTypeSlash       TokenType = "SLASH"
	TokenTypeLessThan    TokenType = "LESS_THAN"
	TokenTypeGreaterThan TokenType = "GREATER_THAN"
	TokenTypeBang        TokenType = "BANG"
	TokenTypeComma       TokenType = "COMMA"
	TokenTypeSemicolon   TokenType = "SEMICOLON"
	TokenTypeLeftParen   TokenType = "LEFT_PAREN"
	TokenTypeRightParen  TokenType = "RIGHT_PAREN"
	TokenTypeLeftBrace   TokenType = "LEFT_BRACE"
  TokenTypeRightBrace  TokenType = "RIGHT_BRACE"
	TokenTypeFunction    TokenType = "FUNCTION"
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
