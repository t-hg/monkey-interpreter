package ast

import "fmt"

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
