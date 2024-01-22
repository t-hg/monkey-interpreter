package ast

import (
  "fmt"
)

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

func (stmt LetStatement) isStatement() bool {
  return true
}

func (stmt LetStatement) String() string {
  return fmt.Sprintf("%s %s = %s;", stmt.Literal, stmt.Identifier, stmt.Expression)
}

type ReturnStatement struct {
  Literal string
  Expression Expression
}

func (stmt ReturnStatement) isStatement() bool {
  return true
}

func (stmt ReturnStatement) String() string {
  return fmt.Sprintf("%s %s", stmt.Literal, stmt.Expression)
}
