package repl

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/t-hg/monkey-interpreter/lexer"
	. "github.com/t-hg/monkey-interpreter/token"
)


const Prompt = "monkey> "

func StartRepl() {
  scanner := bufio.NewScanner(os.Stdin)

  for {
    fmt.Printf(Prompt)
    if !scanner.Scan() {
      return
    }
    line := scanner.Text()
    lexer := NewLexer(line)
    for token := lexer.NextToken(); token.Type != TokenTypeEOF; token = lexer.NextToken() {
      fmt.Printf("%s\n", token) 
    }
  }
}
