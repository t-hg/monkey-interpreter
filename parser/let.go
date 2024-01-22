package parser

import (
 . "github.com/t-hg/monkey-interpreter/ast"
 . "github.com/t-hg/monkey-interpreter/token"
)

func (parser *Parser) parseLetStatement() error {
	stmt := LetStatement{
		Literal: parser.token.Literal,
	}
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

