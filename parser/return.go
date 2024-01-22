package parser

import (
 . "github.com/t-hg/monkey-interpreter/ast"
 . "github.com/t-hg/monkey-interpreter/token"
)

func (parser *Parser) parseReturnStatement() error {
	stmt := ReturnStatement{
		Literal: parser.token.Literal,
	}
	// TODO parse expression
	for parser.token.Type != TokenTypeSemicolon {
		parser.nextToken()
	}
	parser.nextToken()
	parser.appendStatement(stmt)
  return nil
}
