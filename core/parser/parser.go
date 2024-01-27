package parser

import (
	"github.com/Nie-Mand/karui/core/common"
	"github.com/Nie-Mand/karui/core/internal"
	"github.com/Nie-Mand/karui/core/internal/tokens"
)

type Parser struct {
	iterator common.Iterator[tokens.Token]
}

func NewParser(_tokens []tokens.Token) *Parser {
	return &Parser{
		iterator: tokens.NewTokenIterator(_tokens),
	}
}

func (p *Parser) ParseTerm() (TermType, error) {

	if p._hasNextAs(tokens.IntLiteral) {
		integerLiteral, _ := p.iterator.Next()
		intLiteralTerm := &IntLiteralTerm{
			Token: integerLiteral,
		}

		return intLiteralTerm, nil
	} else if p._hasNextAs(tokens.Identifier) {
		identifier, _ := p.iterator.Next()
		identifierTerm := &IdentifierTerm{
			Token: identifier,
		}

		return identifierTerm, nil
	} else if p._hasNextAs(tokens.OpenParenthesis) {
		p.iterator.Next()
		parenthesizedTerm := &ParenthesizedTerm{}
		expression, err := p.ParseExpression(0)

		if err != nil {
			return nil, err
		}

		parenthesizedTerm.Expression = expression

		if !p.iterator.HasNext() || !p._hasNextAs(tokens.CloseParenthesis) {
			return nil, ErrMissingCloseParenthesis
		}
		p.iterator.Next()

		return parenthesizedTerm, nil
	} else {
		return nil, ErrInvalidToken
	}
}

func (p *Parser) ParseExpression(minimumPrecision int) (ExpressionType, error) {
	leftTerm, err := p.ParseTerm()

	if err != nil {
		return nil, err
	}

	// fmt.Println("leftTerm", leftTerm)
	for {
		token, err := p.iterator.Peek()
		precision := token.BinaryPrec()

		if err == nil {
			if precision < minimumPrecision {
				break
			}
		} else {
			break
		}

		operation, err := p.iterator.Next()
		if err != nil {
			return nil, err
		}

		rightTerm, err := p.ParseExpression(precision + 1)

		if err != nil {
			return nil, err
		}

		switch operation.Type {
		case tokens.Plus:
			expresion := &AddExpression{
				BinaryExpression: BinaryExpression{
					Left:  leftTerm,
					Right: rightTerm,
				},
			}

			return expresion, nil

		case tokens.Minus:
			expression := &SubtractExpression{
				BinaryExpression: BinaryExpression{
					Left:  leftTerm,
					Right: rightTerm,
				},
			}

			return expression, nil

		case tokens.Multiply:
			expression := &MultiplyExpression{
				BinaryExpression: BinaryExpression{
					Left:  leftTerm,
					Right: rightTerm,
				},
			}

			return expression, nil
		case tokens.Divide:
			expression := &DivideExpression{
				BinaryExpression: BinaryExpression{
					Left:  leftTerm,
					Right: rightTerm,
				},
			}

			return expression, nil

		case tokens.Modulo:
			expression := &ModuloExpression{
				BinaryExpression: BinaryExpression{
					Left:  leftTerm,
					Right: rightTerm,
				},
			}

			return expression, nil
		default:
			return nil, ErrInvalidToken
		}

	}

	return leftTerm, nil
}

func (p *Parser) ParseScope() (*Scope, error) {
	if !p._hasNextAs(internal.SCOPE_START) {
		return nil, ErrInvalidScopeStart
	}

	p.iterator.Next()
	scope := NewScope()

	for {
		statement, err := p.ParseStatement()
		if err != nil {
			break
		}

		scope.Statements = append(scope.Statements, statement)
	}
	if !p._hasNextAs(internal.SCOPE_END) {
		return nil, ErrInvalidScopeEnd
	}

	p.iterator.Next()

	return scope, nil
}

func (p *Parser) ParseStatement() (StatementType, error) {
	// Handle Exit Statement
	if p._hasNextAs(tokens.Exit) {
		return p._parseExitStatement()
	}

	// Handle Let Statement
	if p._hasNextAs(tokens.Let) {
		return p._parseLetStatement()
	}

	// Handle Puts Statement
	if p._hasNextAs(tokens.Puts) {
		return p._parsePutsStatement()
	}

	// Handle If Statement
	if p._hasNextAs(tokens.If) {
		return p._parseIfStatement()
	}

	// Handle Scope
	if p._hasNextAs(internal.SCOPE_START) {
		return p.ParseScope()
	}

	return nil, ErrInvalidStatement
}

func (p *Parser) ParseProgram() (*Program, error) {
	program := NewProgram()

	for p.iterator.HasNext() {
		statement, err := p.ParseStatement()
		if err != nil {
			return nil, err
		}

		program.Statements = append(program.Statements, statement)
	}

	return program, nil
}

// [*] HELPER FUNCTIONS
func (p *Parser) _hasNextAs(t tokens.TokenType) bool {
	if token, err := p.iterator.Peek(); err == nil {
		return token.Type == t
	}
	return false
}

func (p *Parser) _parseExitStatement() (*ExitStatement, error) {
	p.iterator.Next()
	exitStatement := ExitStatement{}

	if !p.iterator.HasNext() || !p._hasNextAs(tokens.OpenParenthesis) {
		return nil, ErrMissingOpenParenthesis
	}
	p.iterator.Next()

	if !p.iterator.HasNext() || !(p._hasNextAs(tokens.IntLiteral) || p._hasNextAs(tokens.Identifier)) {
		return nil, ErrMissingExitCode
	}
	expression, err := p.ParseExpression(0)
	if err != nil {
		return nil, err
	}

	if !p.iterator.HasNext() || !p._hasNextAs(tokens.CloseParenthesis) {
		return nil, ErrMissingCloseParenthesis
	}
	p.iterator.Next()

	if !p.iterator.HasNext() || !p._hasNextAs(tokens.Semicolon) {
		return nil, ErrUnexpectedToken
	}
	p.iterator.Next()

	exitStatement.ExitCode = expression
	return &exitStatement, nil
}

func (p *Parser) _parseLetStatement() (*LetStatement, error) {
	p.iterator.Next()
	letStatement := &LetStatement{}
	if !p.iterator.HasNext() || !p._hasNextAs(tokens.Identifier) {
		return nil, ErrMissingIdentifier
	}
	identfier, _ := p.iterator.Next()
	letStatement.Token = identfier

	if !p.iterator.HasNext() || !p._hasNextAs(tokens.Equal) {
		return nil, ErrUnexpectedToken
	}
	p.iterator.Next()

	expression, err := p.ParseExpression(0)
	if err != nil {
		return nil, err
	}

	letStatement.Expression = expression

	if !p.iterator.HasNext() || !p._hasNextAs(tokens.Semicolon) {
		return nil, ErrUnexpectedToken
	}
	p.iterator.Next()

	return letStatement, nil
}

func (p *Parser) _parsePutsStatement() (*PutStatement, error) {
	p.iterator.Next()
	putsStatement := &PutStatement{}
	if !p.iterator.HasNext() || !p._hasNextAs(tokens.OpenParenthesis) {
		return nil, ErrMissingOpenParenthesis
	}
	p.iterator.Next()

	expression, err := p.ParseExpression(0)

	if err != nil {
		return nil, err
	}

	putsStatement.Expression = expression

	if !p.iterator.HasNext() || !p._hasNextAs(tokens.CloseParenthesis) {
		return nil, ErrMissingCloseParenthesis
	}
	p.iterator.Next()

	if !p.iterator.HasNext() || !p._hasNextAs(tokens.Semicolon) {
		return nil, ErrUnexpectedToken
	}
	p.iterator.Next()

	return putsStatement, nil
}

func (p *Parser) _parseIfStatement() (*IfStatement, error) {
	p.iterator.Next()

	ifStatement := &IfStatement{}

	if !p.iterator.HasNext() || !p._hasNextAs(tokens.OpenParenthesis) {
		return nil, ErrMissingOpenParenthesis
	}
	p.iterator.Next()

	expression, err := p.ParseExpression(0)

	if err != nil {
		return nil, err
	}

	ifStatement.Condition = expression

	if !p.iterator.HasNext() || !p._hasNextAs(tokens.CloseParenthesis) {
		return nil, ErrMissingCloseParenthesis
	}

	p.iterator.Next()

	scope, err := p.ParseScope()

	if err != nil {
		return nil, err
	}

	ifStatement.Scope = *scope

	return ifStatement, nil
}
