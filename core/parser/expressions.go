package parser

type ExpressionType interface {
	IsExpression()
	String() string
}


// BinaryExpressions
type BinaryExpression struct {
	Left ExpressionType
	Right ExpressionType
}

func (BinaryExpression) IsExpression() {}
func (b *BinaryExpression) String() string {
	return b.Left.String() + " " + b.Right.String()
}


// AddExpression
type AddExpression struct {
	BinaryExpression
}

func (AddExpression) IsExpression() {}
func (x *AddExpression) String() string {
	return x.Left.String() + " + " + x.Right.String()
}



// SubtractExpression
type SubtractExpression struct {
	BinaryExpression
}

func (SubtractExpression) IsExpression() {}
func (s *SubtractExpression) String() string {
	return s.Left.String() + " - " + s.Right.String()
}


// MultiplyExpression
type MultiplyExpression struct {
	BinaryExpression
}

func (MultiplyExpression) IsExpression() {}
func (s* MultiplyExpression) String() string {
	return s.Left.String() + " * " + s.Right.String()
}


// DivideExpression
type DivideExpression struct {
	BinaryExpression
}

func (DivideExpression) IsExpression() {}
func (s *DivideExpression) String() string {
	return s.Left.String() + " / " + s.Right.String()
}


// ModuloExpression
type ModuloExpression struct {
	BinaryExpression
}

func (ModuloExpression) IsExpression() {}
func (s *ModuloExpression) String() string {
	return s.Left.String() + " % " + s.Right.String()
}