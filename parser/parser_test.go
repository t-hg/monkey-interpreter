package parser

import (
	"testing"

	. "github.com/t-hg/monkey-interpreter/ast"
	. "github.com/t-hg/monkey-interpreter/lexer"
)

func TestLetStatement(t *testing.T) {
	input := `
let x = 1;
let y = 5;
let foobar = 10;
`
	lexer := NewLexer(input)
	parser := NewParser(lexer)
	program, err := parser.Parse()
	if err != nil {
		t.Fatal("parsing failed:", err)
	}

	stmts := []Statement{
		LetStatement{
			Literal: "let",
			Identifier: Identifier{
				Literal: "x",
			},
		},
		LetStatement{
			Literal: "let",
			Identifier: Identifier{
				Literal: "y",
			},
		},
		LetStatement{
			Literal: "let",
			Identifier: Identifier{
				Literal: "foobar",
			},
		},
	}

	for index := 0; index < len(program.Statements); index++ {
		expected := stmts[index]
		actual := program.Statements[index]
		if expected.String() != actual.String() {
			t.Errorf("Expected '%s', got '%s'", expected, actual)
		}
	}
}
