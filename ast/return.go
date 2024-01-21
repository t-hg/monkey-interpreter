package ast

import "fmt"

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
