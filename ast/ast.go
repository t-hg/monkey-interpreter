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

// Expressions

type IdentifierExpression struct {
	Literal string
}

func (expr IdentifierExpression) isExpression() bool {
  return true
}

func (expr IdentifierExpression) String() string {
  return fmt.Sprintf("%s", expr.Literal)
}

type IntegerExpression struct {
  Value int
}

func (expr IntegerExpression) isExpression() bool {
  return true
}

func (expr IntegerExpression) String() string {
  return fmt.Sprintf("%d", expr.Value)
}

type PrefixExpression struct {
  Operator string
  Right Expression
}

func (expr PrefixExpression) isExpression() bool {
  return true
}

func (expr PrefixExpression) String() string {
  return fmt.Sprintf("(%s%s)", expr.Operator, expr.Right)
}

// Statements

type LetStatement struct {
	Identifier IdentifierExpression
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
  return fmt.Sprintf("return %s;", stmt.Expression)
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
  return fmt.Sprintf("%s;", stmt.Expression)
}

