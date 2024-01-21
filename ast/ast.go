package ast

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

