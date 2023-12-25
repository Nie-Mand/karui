package lexer

// - TokenType: enum (exit, number, parenthesis, +-*/^, number, if.. variable itself: ident or int ident)
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