package ast

type Identifier struct {
	Literal string
}

func (identifier Identifier) String() string {
  return identifier.Literal
}
