package optimizer

import (
	"github.com/rosylilly/brainfxxk/ast"
)

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

		if _, ok := optExpr.(*ast.Comment); ok {
			continue
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
				if last, ok := optimized[len(optimized)-1].(*ast.ValueChangeExpression); ok {
					last.Count += 1
					last.Expressions = append(last.Expressions, optExpr)
					continue
				}
			}

			optExpr = &ast.ValueChangeExpression{
				Count:       1,
				Expressions: []ast.Expression{optExpr},
			}
		case *ast.ValueDecrementExpression:
			if len(optimized) > 0 {
				if last, ok := optimized[len(optimized)-1].(*ast.ValueChangeExpression); ok {
					last.Count -= 1
					last.Expressions = append(last.Expressions, optExpr)
					continue
				}
			}

			optExpr = &ast.ValueChangeExpression{
				Count:       -1,
				Expressions: []ast.Expression{optExpr},
			}
		case *ast.WhileExpression:
			if len(optExpr.(*ast.WhileExpression).Body) == 1 {
				switch optExpr.(*ast.WhileExpression).Body[0].(type) {
				case *ast.ValueDecrementExpression:
					optExpr = &ast.ValueResetExpression{Pos: optExpr.StartPos()}
				}
			} else {
				opBody, err := o.optimizeExpressions(optExpr.(*ast.WhileExpression).Body)
				if err != nil {
					return nil, err
				}

				// Filter out comments from the optimized body
				nonCommentBody := []ast.Expression{}
				for _, expr := range opBody {
					if _, ok := expr.(*ast.Comment); !ok {
						nonCommentBody = append(nonCommentBody, expr)
					}
				}

				if len(nonCommentBody) == 1 {
					switch e := nonCommentBody[0].(type) {
					case *ast.PointerMoveExpression:
						optExpr = &ast.ZeroSearchExpression{
							StartPosition: optExpr.StartPos(),
							EndPosition:   optExpr.EndPos(),
							SearchWindow:  e.Count,
						}
					default:
						optExpr.(*ast.WhileExpression).Body = nonCommentBody
					}
				} else {
					optExpr.(*ast.WhileExpression).Body = nonCommentBody
				}
			}
		}

		if _, ok := optExpr.(*ast.Comment); !ok {
			optimized = append(optimized, optExpr)
		}
	}

	return optimized, nil
}

func (o *Optimizer) optimizeExpression(expr ast.Expression) (ast.Expression, error) {

	return expr, nil
}
