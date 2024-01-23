package parser

import (
	. "github.com/t-hg/monkey-interpreter/ast"
	. "github.com/t-hg/monkey-interpreter/lexer"
	. "github.com/t-hg/monkey-interpreter/token"
)

type parsePrefixFn func() (Expression, error)
type parseInfixFn func(Expression) Expression

type Parser struct {
	lexer   *Lexer
	program *Program
	token   Token
  parsePrefixFns map[TokenType]parsePrefixFn
  parseInfixFns map[TokenType]parseInfixFn
}

func NewParser(lexer *Lexer) *Parser {
  parser := &Parser{
		lexer: lexer,
		program: &Program{
			Statements: make([]Statement, 0),
		},
		token: lexer.NextToken(),
    parsePrefixFns: make(map[TokenType]parsePrefixFn),
    parseInfixFns: make(map[TokenType]parseInfixFn),
	}

  parser.registerPrefixParser(TokenTypeIdentifier, parser.parseIdentifier)
  return parser
}

func (parser *Parser) registerPrefixParser(tokenType TokenType, fn parsePrefixFn) {
  parser.parsePrefixFns[tokenType] = fn
}
func (parser *Parser) registerInfixParser(tokenType TokenType, fn parseInfixFn) {
  parser.parseInfixFns[tokenType] = fn
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
      err = parser.parseExpressionStatement()
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
		return unexpectedToken2(tokenType, parser.token.Type)
	}
	return nil
}

func (parser *Parser) parseLetStatement() error {
	stmt := LetStatement{}
	parser.nextToken()
	if err := parser.expectToken(TokenTypeIdentifier); err != nil {
		return err
	}
	stmt.Identifier = Identifier{
    Literal: parser.token.Literal,
	}
	parser.nextToken()
	if err := parser.expectToken(TokenTypeAssign); err != nil {
		return err
	}
	// TODO parse expression
	for parser.token.Type != TokenTypeSemicolon {
		parser.nextToken()
	}
	parser.nextToken()
	parser.appendStatement(stmt)
	return nil
}

func (parser *Parser) parseReturnStatement() error {
	stmt := ReturnStatement{}
	// TODO parse expression
	for parser.token.Type != TokenTypeSemicolon {
		parser.nextToken()
	}
	parser.nextToken()
	parser.appendStatement(stmt)
  return nil
}

func (parser *Parser) parseExpressionStatement() error {
  stmt := ExpressionStatement{}
  expr, err := parser.parseExpression(PrecedenceLowest)
  if err != nil {
    return err
  }
  stmt.Expression = expr
  parser.appendStatement(stmt)
  return nil
}

func (parser *Parser) parseExpression(precedence Precedence) (Expression, error) {
  tokenType := parser.token.Type
  parsePrefix := parser.parsePrefixFns[tokenType]
  if parsePrefix == nil {
    return nil, unexpectedToken(tokenType)
  }
  return parsePrefix()
}

func (parser *Parser) parseIdentifier() (Expression, error) {
  identifier := Identifier{Literal: parser.token.Literal}
  parser.nextToken()
  if parser.token.Type == TokenTypeSemicolon {
    parser.nextToken()
  }
  return identifier, nil
}
