package ast

type MultiplePointerIncrementExpression struct {
	Count       int
	Expressions []Expression
}

func (e *MultiplePointerIncrementExpression) StartPos() int {
	return e.Expressions[0].StartPos()
}

func (e *MultiplePointerIncrementExpression) EndPos() int {
	return e.Expressions[len(e.Expressions)-1].EndPos()
}

func (e *MultiplePointerIncrementExpression) Bytes() []byte {
	b := []byte{}
	for _, expr := range e.Expressions {
		b = append(b, expr.Bytes()...)
	}
	return b
}

func (e *MultiplePointerIncrementExpression) String() string {
	return string(e.Bytes())
}

type MultiplePointerDecrementExpression struct {
	Count       int
	Expressions []Expression
}

func (e *MultiplePointerDecrementExpression) StartPos() int {
	return e.Expressions[0].StartPos()
}

func (e *MultiplePointerDecrementExpression) EndPos() int {
	return e.Expressions[len(e.Expressions)-1].EndPos()
}

func (e *MultiplePointerDecrementExpression) Bytes() []byte {
	b := []byte{}
	for _, expr := range e.Expressions {
		b = append(b, expr.Bytes()...)
	}
	return b
}

func (e *MultiplePointerDecrementExpression) String() string {
	return string(e.Bytes())
}

type PointerMoveExpression struct {
	Count       int
	Expressions []Expression
}

func (e *PointerMoveExpression) StartPos() int {
	return e.Expressions[0].StartPos()
}

func (e *PointerMoveExpression) EndPos() int {
	return e.Expressions[len(e.Expressions)-1].EndPos()
}

func (e *PointerMoveExpression) Bytes() []byte {
	b := []byte{}
	for _, expr := range e.Expressions {
		b = append(b, expr.Bytes()...)
	}
	return b
}

func (e *PointerMoveExpression) String() string {
	return string(e.Bytes())
}

type MultipleValueIncrementExpression struct {
	Count       int
	Expressions []Expression
}

func (e *MultipleValueIncrementExpression) StartPos() int {
	return e.Expressions[0].StartPos()
}

func (e *MultipleValueIncrementExpression) EndPos() int {
	return e.Expressions[len(e.Expressions)-1].EndPos()
}

func (e *MultipleValueIncrementExpression) Bytes() []byte {
	b := []byte{}
	for _, expr := range e.Expressions {
		b = append(b, expr.Bytes()...)
	}
	return b
}

func (e *MultipleValueIncrementExpression) String() string {
	return string(e.Bytes())
}

type MultipleValueDecrementExpression struct {
	Count       int
	Expressions []Expression
}

func (e *MultipleValueDecrementExpression) StartPos() int {
	return e.Expressions[0].StartPos()
}

func (e *MultipleValueDecrementExpression) EndPos() int {
	return e.Expressions[len(e.Expressions)-1].EndPos()
}

func (e *MultipleValueDecrementExpression) Bytes() []byte {
	b := []byte{}
	for _, expr := range e.Expressions {
		b = append(b, expr.Bytes()...)
	}
	return b
}

func (e *MultipleValueDecrementExpression) String() string {
	return string(e.Bytes())
}

type ValueChangeExpression struct {
	Count       int
	Expressions []Expression
}

func (e *ValueChangeExpression) StartPos() int {
	return e.Expressions[0].StartPos()
}

func (e *ValueChangeExpression) EndPos() int {
	return e.Expressions[len(e.Expressions)-1].EndPos()
}

func (e *ValueChangeExpression) Bytes() []byte {
	b := []byte{}
	for _, expr := range e.Expressions {
		b = append(b, expr.Bytes()...)
	}
	return b
}

func (e *ValueChangeExpression) String() string {
	return string(e.Bytes())
}

type ValueResetExpression struct {
	Pos int
}

func (e *ValueResetExpression) StartPos() int {
	return e.Pos
}

func (e *ValueResetExpression) EndPos() int {
	return e.Pos + 3
}

func (e *ValueResetExpression) Bytes() []byte {
	return []byte{'[', '-', ']'}
}

func (e *ValueResetExpression) String() string {
	return string(e.Bytes())
}

type ZeroSearchExpression struct {
	StartPosition int
	EndPosition   int

	SearchWindow int
}

func (e *ZeroSearchExpression) StartPos() int {
	return e.StartPos()
}

func (e *ZeroSearchExpression) EndPos() int {
	return e.EndPos()
}

func (e *ZeroSearchExpression) Bytes() []byte {
	b := []byte{'['}
	if e.SearchWindow != 0 {
		symbol := byte('>')
		if e.SearchWindow < 0 {
			symbol = '<'
		}
		count := e.SearchWindow
		if count < 0 {
			count = -count
		}
		for i := 0; i < count; i++ {
			b = append(b, symbol)
		}
	}
	b = append(b, ']')
	return b
}

func (e *ZeroSearchExpression) String() string {
	return string(e.Bytes())
}
