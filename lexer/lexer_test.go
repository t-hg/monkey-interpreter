package lexer

import (
	"testing"

	. "github.com/t-hg/monkey-interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;

let ten = 10;
let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
`

	expectedTokens := []Token{
		{Type: TokenTypeLet, Literal: "let"},
		{Type: TokenTypeIdentifier, Literal: "five"},
		{Type: TokenTypeAssign, Literal: "="},
		{Type: TokenTypeInteger, Literal: "5"},
		{Type: TokenTypeSemicolon, Literal: ";"},
		{Type: TokenTypeLet, Literal: "let"},
		{Type: TokenTypeIdentifier, Literal: "ten"},
		{Type: TokenTypeAssign, Literal: "="},
		{Type: TokenTypeInteger, Literal: "10"},
		{Type: TokenTypeSemicolon, Literal: ";"},
		{Type: TokenTypeLet, Literal: "let"},
		{Type: TokenTypeIdentifier, Literal: "add"},
		{Type: TokenTypeAssign, Literal: "="},
		{Type: TokenTypeFunction, Literal: "fn"},
		{Type: TokenTypeLParen, Literal: "("},
		{Type: TokenTypeIdentifier, Literal: "x"},
		{Type: TokenTypeComma, Literal: ","},
		{Type: TokenTypeIdentifier, Literal: "y"},
		{Type: TokenTypeRParen, Literal: ")"},
		{Type: TokenTypeLBrace, Literal: "{"},
		{Type: TokenTypeIdentifier, Literal: "x"},
		{Type: TokenTypePlus, Literal: "+"},
		{Type: TokenTypeIdentifier, Literal: "y"},
		{Type: TokenTypeSemicolon, Literal: ";"},
		{Type: TokenTypeRBrace, Literal: "}"},
		{Type: TokenTypeSemicolon, Literal: ";"},
		{Type: TokenTypeLet, Literal: "let"},
		{Type: TokenTypeIdentifier, Literal: "result"},
		{Type: TokenTypeAssign, Literal: "="},
		{Type: TokenTypeIdentifier, Literal: "add"},
		{Type: TokenTypeLParen, Literal: "("},
		{Type: TokenTypeIdentifier, Literal: "five"},
		{Type: TokenTypeComma, Literal: ","},
		{Type: TokenTypeIdentifier, Literal: "ten"},
		{Type: TokenTypeRParen, Literal: ")"},
		{Type: TokenTypeSemicolon, Literal: ";"},
		{Type: TokenTypeEOF, Literal: "\x00"},
	}

	lexer := NewLexer(input)

	for _, expectedToken := range expectedTokens {
		actualToken := lexer.NextToken()
		if actualToken != expectedToken {
			t.Errorf("Expected '%s', got '%s'", expectedToken, actualToken)
		}
	}
}

func TestIsWhitespace(t *testing.T) {
	chars := []string{" ", "\t", "\r", "\n"}

	for _, char := range chars {
		if !isWhitespace(char) {
			t.Errorf("Expected '%s' is whitespace", char)
		}
	}
}

func TestIsLetter(t *testing.T) {
	chars := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}

	for _, char := range chars {
		if !isLetter(char) {
			t.Errorf("Expected '%s' is letter", char)
		}
	}
}

func TestIsDigit(t *testing.T) {
	chars := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for _, char := range chars {
		if !isDigit(char) {
			t.Errorf("Expected '%s' is digit", char)
		}
	}
}
