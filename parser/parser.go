package parser

import (
	"fmt"

	. "github.com/t-hg/monkey-interpreter/ast"
	. "github.com/t-hg/monkey-interpreter/lexer"
	. "github.com/t-hg/monkey-interpreter/token"
)

type ParseError struct {
	message string
}

func (err ParseError) Error() string {
	return err.message
}

func newParseError(format string, args ...interface{}) ParseError {
	return ParseError{
		message: fmt.Sprintf(format, args...),
	}
}

func unexpectedToken(expected TokenType, actual TokenType) ParseError {
	return newParseError("Expected '%s', got '%s'", expected, actual)
}

type Parser struct {
	lexer   *Lexer
	program *Program
	token   Token
}

func NewParser(lexer *Lexer) *Parser {
	return &Parser{
		lexer: lexer,
		program: &Program{
			Statements: make([]Statement, 0),
		},
		token: lexer.NextToken(),
	}
}

func (parser *Parser) Parse() (*Program, error) {
	var err error

	for parser.token.Type != TokenTypeEOF {
		switch parser.token.Type {
		case TokenTypeLet:
			err = parser.parseLetStatement()
    case TokenTypeReturn:
      err = parser.parseReturnStatement()
		default:
			err = newParseError("Unexpected token '%s'", parser.token)
		}
		if err != nil {
			break
		}
	}

	return parser.program, err
}

func (parser *Parser) nextToken() {
	parser.token = parser.lexer.NextToken()
}

func (parser *Parser) appendStatement(stmt Statement) {
	parser.program.Statements = append(parser.program.Statements, stmt)
}

func (parser *Parser) expectToken(tokenType TokenType) error {
	if tokenType != parser.token.Type {
		return unexpectedToken(tokenType, parser.token.Type)
	}
	return nil
}

