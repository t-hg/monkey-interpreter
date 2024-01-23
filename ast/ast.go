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

func (identifier Identifier) isExpression() bool {
  return true
}

func (identifier Identifier) String() string {
  return fmt.Sprintf("%s;", identifier.Literal)
}

type Integer struct {
  Value int
}

func (integer Integer) isExpression() bool {
  return true
}

func (integer Integer) String() string {
  return fmt.Sprintf("%d;", integer.Value)
}

type LetStatement struct {
	Identifier Identifier
	Expression Expression
}

func (stmt LetStatement) isStatement() bool {
  return true
}

func (stmt LetStatement) String() string {
  return fmt.Sprintf("let %s = %s;", stmt.Identifier, stmt.Expression)
}

type ReturnStatement struct {
  Expression Expression
}

func (stmt ReturnStatement) isStatement() bool {
  return true
}

func (stmt ReturnStatement) String() string {
  return fmt.Sprintf("return %s", stmt.Expression)
}

// A statement that contains out of a single expression
// like:
// monkey> 5 + 5;
type ExpressionStatement struct {
  Expression Expression
}

func (stmt ExpressionStatement) isStatement() bool {
  return true
}

func (stmt ExpressionStatement) String() string {
  return fmt.Sprintf("%s", stmt.Expression)
}
