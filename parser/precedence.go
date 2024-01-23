package parser

type Precedence int

const (
  PrecedenceLowest Precedence = iota
  PrecedenceEquals
  PrecedenceLessGreater
  PrecedenceSum
  PrecedenceProduct
  PrecedencePrefix
  PrecedenceCall
)
