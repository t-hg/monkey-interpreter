package lexer

import (
	. "github.com/t-hg/monkey-interpreter/token"
)

type Lexer struct {
	input    string
	position int
	char     string
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:    input,
		position: 0,
		char:     string(input[0]),
	}
}

func (lexer *Lexer) NextToken() Token {
	lexer.skipWhitespace()

	switch lexer.char {
	case "=":
    lexer.nextChar()
		return Token{Type: TokenTypeAssign, Literal: "="}
	case "+":
    lexer.nextChar()
		return Token{Type: TokenTypePlus, Literal: "+"}
	case "-":
    lexer.nextChar()
		return Token{Type: TokenTypeMinus, Literal: "-"}
	case "*":
    lexer.nextChar()
		return Token{Type: TokenTypeAsterisk, Literal: "*"}
	case "/":
    lexer.nextChar()
		return Token{Type: TokenTypeSlash, Literal: "/"}
	case "<":
    lexer.nextChar()
		return Token{Type: TokenTypeLessThan, Literal: "<"}
	case ">":
    lexer.nextChar()
		return Token{Type: TokenTypeGreaterThan, Literal: ">"}
	case "!":
    lexer.nextChar()
		return Token{Type: TokenTypeBang, Literal: "!"}
	case ",":
    lexer.nextChar()
		return Token{Type: TokenTypeComma, Literal: ","}
	case ";":
    lexer.nextChar()
		return Token{Type: TokenTypeSemicolon, Literal: ";"}
	case "(":
    lexer.nextChar()
		return Token{Type: TokenTypeLParen, Literal: "("}
	case ")":
    lexer.nextChar()
		return Token{Type: TokenTypeRParen, Literal: ")"}
	case "{":
    lexer.nextChar()
		return Token{Type: TokenTypeLBrace, Literal: "{"}
	case "}":
    lexer.nextChar()
		return Token{Type: TokenTypeRBrace, Literal: "}"}
	case "\x00":
		return Token{Type: TokenTypeEOF, Literal: "\x00"}
	}

	if isLetter(lexer.char) {
		literal := ""
		for isLetter(lexer.char) || isDigit(lexer.char) {
			literal += lexer.char
			lexer.nextChar()
		}
		switch literal {
		case "fn":
			return Token{Type: TokenTypeFunction, Literal: literal}
		case "let":
			return Token{Type: TokenTypeLet, Literal: literal}
		case "true":
			return Token{Type: TokenTypeTrue, Literal: literal}
		case "false":
			return Token{Type: TokenTypeFalse, Literal: literal}
		case "if":
			return Token{Type: TokenTypeIf, Literal: literal}
		case "else":
			return Token{Type: TokenTypeElse, Literal: literal}
		default:
			return Token{Type: TokenTypeIdentifier, Literal: literal}
		}
	}

	if isDigit(lexer.char) {
		literal := ""
		for isDigit(lexer.char) {
			literal += lexer.char
			lexer.nextChar()
		}
		return Token{Type: TokenTypeInteger, Literal: literal}
	}

	return Token{Type: TokenTypeIllegal, Literal: lexer.char}
}

func (lexer *Lexer) nextChar() {
	lexer.position += 1
	if lexer.position < len(lexer.input) {
		lexer.char = string(lexer.input[lexer.position])
	} else {
		lexer.char = "\x00"
	}
}

func (lexer *Lexer) skipWhitespace() {
	for isWhitespace(lexer.char) {
		lexer.nextChar()
	}
}

func isWhitespace(char string) bool {
	return char == " " || char == "\t" || char == "\n" || char == "\r"
}

func isLetter(char string) bool {
	return char[0] >= 65 && char[0] <= 90 || char[0] >= 97 && char[0] <= 122
}

func isDigit(char string) bool {
	return char[0] >= 48 && char[0] <= 57
}
