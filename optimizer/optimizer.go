package optimizer

import "github.com/rosylilly/brainfxxk/ast"

type Optimizer struct {
}

func NewOptimizer() *Optimizer {
	return &Optimizer{}
}

func (o *Optimizer) Optimize(p *ast.Program) (*ast.Program, error) {
	exprs, err := o.optimizeExpressions(p.Expressions)
	if err != nil {
		return nil, err
	}

	prog := &ast.Program{
		Expressions: exprs,
	}

	return prog, nil
}

func (o *Optimizer) optimizeExpressions(exprs []ast.Expression) ([]ast.Expression, error) {
	optimized := []ast.Expression{}
	for _, expr := range exprs {
		optExpr, err := o.optimizeExpression(expr)
		if err != nil {
			return nil, err
		}

		switch optExpr.(type) {
		case *ast.PointerIncrementExpression:
			if len(optimized) > 0 {
				if last, ok := optimized[len(optimized)-1].(*ast.PointerMoveExpression); ok {
					last.Count += 1
					last.Expressions = append(last.Expressions, optExpr)
					continue
				}
			}

			optExpr = &ast.PointerMoveExpression{
				Count:       1,
				Expressions: []ast.Expression{optExpr},
			}
		
		case *ast.PointerDecrementExpression:
			if len(optimized) > 0 {
				if last, ok := optimized[len(optimized)-1].(*ast.PointerMoveExpression); ok {
					last.Count -= 1
					last.Expressions = append(last.Expressions, optExpr)
					continue
				}
			}

			optExpr = &ast.PointerMoveExpression{
				Count:       -1,
				Expressions: []ast.Expression{optExpr},
			}

		case *ast.ValueIncrementExpression:
			if len(optimized) > 0 {
				if last, ok := optimized[len(optimized)-1].(*ast.MultipleValueIncrementExpression); ok {
					last.Count += 1
					last.Expressions = append(last.Expressions, optExpr)
					continue
				}
			}

			optExpr = &ast.MultipleValueIncrementExpression{
				Count:       1,
				Expressions: []ast.Expression{optExpr},
			}
		case *ast.ValueDecrementExpression:
			if len(optimized) > 0 {
				if last, ok := optimized[len(optimized)-1].(*ast.MultipleValueDecrementExpression); ok {
					last.Count += 1
					last.Expressions = append(last.Expressions, optExpr)
					continue
				}
			}

			optExpr = &ast.MultipleValueDecrementExpression{
				Count:       1,
				Expressions: []ast.Expression{optExpr},
			}
		}

		optimized = append(optimized, optExpr)
	}

	return optimized, nil
}

func (o *Optimizer) optimizeExpression(expr ast.Expression) (ast.Expression, error) {

	return expr, nil
}
