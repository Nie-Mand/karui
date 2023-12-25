package lexer

type TokenType int

const (
	Exit TokenType = iota
	Let
	Puts

	Ident
	IntIdent
	
	Equal
	LeftParenthesis
	RightParenthesis
	Semicolon
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
	case LeftParenthesis:
		return "LeftParenthesis"
	case RightParenthesis:
		return "RightParenthesis"
	case Semicolon:
		return "Semicolon"
	default:
		return "Unknown"
	}
}