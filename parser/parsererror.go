package parser

import (
	"fmt"

  . "github.com/t-hg/monkey-interpreter/token"
)

type ParseError struct {
	message string
}

func (err ParseError) Error() string {
	return err.message
}

func newParseError(format string, args ...interface{}) ParseError {
	return ParseError{
		message: fmt.Sprintf(format, args...),
	}
}

func unexpectedToken(tokenType TokenType) ParseError {
	return newParseError("Unexprected token '%s'", tokenType)
}

func unexpectedToken2(expected TokenType, actual TokenType) ParseError {
	return newParseError("Expected '%s', got '%s'", expected, actual)
}
