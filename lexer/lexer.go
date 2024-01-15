package lexer

import (
  "github.com/t-hg/monkey-interpreter/token"
)

type Lexer struct {
  Input string
}

func New(input string) *Lexer {
  return &Lexer{
    Input: input,
  };
}

func (lexer *Lexer) NextToken() token.Token {
  return token.Token{}
}
