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
			Identifier: IdentifierExpression{
				Literal: "x",
			},
      // TODO Expression
		},
		LetStatement{
			Identifier: IdentifierExpression{
				Literal: "y",
			},
      // TODO Expression
		},
		LetStatement{
			Identifier: IdentifierExpression{
				Literal: "foobar",
			},
      // TODO Expression
		},
	}

	for index := 0; index < len(stmts); index++ {
		expected := stmts[index]
		actual := program.Statements[index]
		if expected.String() != actual.String() {
			t.Errorf("Expected '%s', got '%s'", expected, actual)
		}
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
return 1;
return 5;
return add(10, 5);
`
	lexer := NewLexer(input)
	parser := NewParser(lexer)
	program, err := parser.Parse()
	if err != nil {
		t.Fatal("parsing failed:", err)
	}
  
	stmts := []Statement{
		ReturnStatement{
      // TODO Expression
		},
		ReturnStatement{
      // TODO Expression
		},
		ReturnStatement{
      // TODO Expression
		},
	}

	for index := 0; index < len(stmts); index++ {
		expected := stmts[index]
		actual := program.Statements[index]
		if expected.String() != actual.String() {
			t.Errorf("Expected '%s', got '%s'", expected, actual)
		}
	}
}

func TestExpressionStatement(t * testing.T) {
  input := `
foo;
5;
-42;
!bar;
`
	lexer := NewLexer(input)
	parser := NewParser(lexer)
	program, err := parser.Parse()
	if err != nil {
		t.Fatal("parsing failed:", err)
	}

	stmts := []Statement{
		ExpressionStatement{
      Expression: IdentifierExpression{
        Literal: "foo",
      },
		},
    ExpressionStatement{
      Expression: IntegerExpression{
        Value: 5,
      },
    },
    ExpressionStatement{
      Expression: PrefixExpression{
        Operator: "-",
        Right: IntegerExpression{
          Value: 42,
        },
      },
    },
    ExpressionStatement{
      Expression: PrefixExpression{
        Operator: "!",
        Right: IdentifierExpression{
          Literal: "bar",
        },
      },
    },
	}

	for index := 0; index < len(stmts); index++ {
		expected := stmts[index]
		actual := program.Statements[index]
		if expected.String() != actual.String() {
			t.Errorf("Expected '%s', got '%s'", expected, actual)
		}
	}
}

