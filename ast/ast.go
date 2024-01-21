package ast

import "fmt"

type Node interface {
  String() string
}

type Statement interface {
	Node
  isStatement() bool
}

type Expression interface {
	Node
  isExpression() bool
}

type Program struct {
	Statements []Statement
}

type Identifier struct {
	Literal string
}

func (identifier Identifier) String() string {
  return identifier.Literal
}

type LetStatement struct {
	Literal    string
	Identifier Identifier
	Expression Expression
}

func (letStatement LetStatement) isStatement() bool {
  return true
}

func (letStatement LetStatement) String() string {
  return fmt.Sprintf("%s %s = %s;", letStatement.Literal, letStatement.Identifier, letStatement.Expression)
}
