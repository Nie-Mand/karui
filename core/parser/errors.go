package parser

import (
	"Nie-Mand/karui/core/lexer"
	"errors"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrInvalidStatement = errors.New("invalid statement")

	ErrInvalidScopeStart = errors.New("expected " + lexer.SCOPE_START.String())
	ErrInvalidScopeEnd = errors.New("expected " + lexer.SCOPE_END.String())

	ErrMissingOpenParenthesis = errors.New("expected (")
	ErrMissingCloseParenthesis = errors.New("expected )")
	ErrMissingExitCode = errors.New("expected exit code")
	ErrUnexpectedToken = errors.New("unexpected token")

	ErrMissingIdentifier = errors.New("expected identifier")



)