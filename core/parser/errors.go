package parser

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrInvalidStatement = errors.New("invalid statement")

	ErrInvalidScopeStart = errors.New("expected {")
	ErrInvalidScopeEnd = errors.New("expected }")

	ErrMissingOpenParenthesis = errors.New("expected (")
	ErrMissingCloseParenthesis = errors.New("expected )")
	ErrMissingExitCode = errors.New("expected exit code")
	ErrUnexpectedToken = errors.New("unexpected token")

	ErrMissingIdentifier = errors.New("expected identifier")

	ErrProgramExits = func(exitCode int) error { 
		return errors.New("program exits with exit code " + fmt.Sprintf("%d", exitCode))
	 }


	ErrMemoryNotInitialized = errors.New("memory not initialized")
	ErrVariableNotInitialized = errors.New("variable not initialized")
	ErrRedeclearationOfVariable = errors.New("redeclaration of variable")
)