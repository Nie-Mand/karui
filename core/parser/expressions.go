package parser

type ExpressionType interface {
	Evaluate(*map[string]int) (int, error)
	String() string
}


// BinaryExpressions
type BinaryExpression struct {
	Left ExpressionType
	Right ExpressionType
}

func (BinaryExpression) Evaluate(_ *map[string]int) (int, error) {
	return 0, nil
}

func (BinaryExpression) String() string {
	return ""
}

// AddExpression
type AddExpression struct {
	BinaryExpression
}

func (x *AddExpression) Evaluate(memory *map[string]int) (int, error) {
	left, err := x.Left.Evaluate(memory)
	if err != nil {
		return 0, err
	}

	right, err := x.Right.Evaluate(memory)
	if err != nil {
		return 0, err
	}

	return left + right, nil
}

func (x *AddExpression) String() string {
	return x.Left.String() + " + " + x.Right.String()
}

// SubtractExpression
type SubtractExpression struct {
	BinaryExpression
}

func (x *SubtractExpression) Evaluate(memory *map[string]int) (int, error) {
	left, err := x.Left.Evaluate(memory)
	if err != nil {
		return 0, err
	}

	right, err := x.Right.Evaluate(memory)
	if err != nil {
		return 0, err
	}

	return left - right, nil
}

func (x *SubtractExpression) String() string {
	return x.Left.String() + " - " + x.Right.String()
}

// MultiplyExpression
type MultiplyExpression struct {
	BinaryExpression
}

func (x *MultiplyExpression) Evaluate(memory *map[string]int) (int, error) {
	left, err := x.Left.Evaluate(memory)
	if err != nil {
		return 0, err
	}

	right, err := x.Right.Evaluate(memory)
	if err != nil {
		return 0, err
	}

	return left * right, nil
}

func (x *MultiplyExpression) String() string {
	return x.Left.String() + " * " + x.Right.String()
}


// DivideExpression
type DivideExpression struct {
	BinaryExpression
}

func (x *DivideExpression) Evaluate(memory *map[string]int) (int, error) {
	left, err := x.Left.Evaluate(memory)
	if err != nil {
		return 0, err
	}

	right, err := x.Right.Evaluate(memory)
	if err != nil {
		return 0, err
	}

	return left / right, nil
}

func (x *DivideExpression) String() string {
	return x.Left.String() + " / " + x.Right.String()
}


// ModuloExpression
type ModuloExpression struct {
	BinaryExpression
}

func (x* ModuloExpression) Evaluate(memory *map[string]int) (int, error) {
	left, err := x.Left.Evaluate(memory)
	if err != nil {
		return 0, err
	}

	right, err := x.Right.Evaluate(memory)
	if err != nil {
		return 0, err
	}

	return left % right, nil
}

func (x *ModuloExpression) String() string {
	return x.Left.String() + " % " + x.Right.String()
}
