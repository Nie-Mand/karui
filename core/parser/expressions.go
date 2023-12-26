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
func (AddExpression) String() string {
	return "AddExpression"
}



// SubtractExpression
type SubtractExpression struct {
	BinaryExpression
}

func (SubtractExpression) IsExpression() {}
func (SubtractExpression) String() string {
	return "SubtractExpression"
}


// MultiplyExpression
type MultiplyExpression struct {
	BinaryExpression
}

func (MultiplyExpression) IsExpression() {}
func (MultiplyExpression) String() string {
	return "MultiplyExpression"
}


// DivideExpression
type DivideExpression struct {
	BinaryExpression
}

func (DivideExpression) IsExpression() {}
func (DivideExpression) String() string {
	return "DivideExpression"
}


// ModuloExpression
type ModuloExpression struct {
	BinaryExpression
}

func (ModuloExpression) IsExpression() {}
func (ModuloExpression) String() string {
	return "ModuloExpression"
}