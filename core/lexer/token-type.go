package lexer

type TokenType int

const (
	// Statements
	Exit TokenType = iota
	Let
	Puts
	If

	// Identifiers
	Ident
	IntIdent
	
	// Symbols
	OpenParenthesis
	CloseParenthesis
	OpenCurly
	CloseCurly
	Semicolon
	
	//
	Equal
	Plus
	Minus
	Multiply
	Divide
	Modulo
)

func (t TokenType) String() string {
	switch t {
	case Exit:
		return "Exit"
	case Let:
		return "Let"
	case Puts:
		return "Puts"
	case If:
		return "If"

	case Ident:
		return "Ident"
	case IntIdent:
		return "IntIdent"

	case Equal:
		return "Equal"
	case Plus:
		return "Plus"
	case Minus:
		return "Minus"
	case Multiply:
		return "Multiply"
	case Divide:
		return "Divide"
	case Modulo:
		return "Modulo"
	case OpenParenthesis:
		return "OpenParenthesis"
	case CloseParenthesis:
		return "CloseParenthesis"
	case OpenCurly:
		return "OpenCurly"
	case CloseCurly:
		return "CloseCurly"
	case Semicolon:
		return "Semicolon"
	default:
		return "Unknown"
	}
}