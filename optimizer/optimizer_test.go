package optimizer_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/rosylilly/brainfxxk/ast"
	"github.com/rosylilly/brainfxxk/optimizer"
	"github.com/rosylilly/brainfxxk/parser"
)

func TestOptimizer(t *testing.T) {
	testCases := []struct {
		source   string
		expected *ast.Program
	}{
		{
			source: "+++++----->>>>><<<<<",
			expected: &ast.Program{
				Expressions: []ast.Expression{
					&ast.ValueChangeExpression{
						Count: 0,
						Expressions: []ast.Expression{
							&ast.ValueIncrementExpression{Pos: 0},
							&ast.ValueIncrementExpression{Pos: 1},
							&ast.ValueIncrementExpression{Pos: 2},
							&ast.ValueIncrementExpression{Pos: 3},
							&ast.ValueIncrementExpression{Pos: 4},
							&ast.ValueDecrementExpression{Pos: 5},
							&ast.ValueDecrementExpression{Pos: 6},
							&ast.ValueDecrementExpression{Pos: 7},
							&ast.ValueDecrementExpression{Pos: 8},
							&ast.ValueDecrementExpression{Pos: 9},
						},
					},
					&ast.PointerMoveExpression{
						Count: 0,
						Expressions: []ast.Expression{
							&ast.PointerIncrementExpression{Pos: 10},
							&ast.PointerIncrementExpression{Pos: 11},
							&ast.PointerIncrementExpression{Pos: 12},
							&ast.PointerIncrementExpression{Pos: 13},
							&ast.PointerIncrementExpression{Pos: 14},
							&ast.PointerDecrementExpression{Pos: 15},
							&ast.PointerDecrementExpression{Pos: 16},
							&ast.PointerDecrementExpression{Pos: 17},
							&ast.PointerDecrementExpression{Pos: 18},
							&ast.PointerDecrementExpression{Pos: 19},
						},
					},
				},
			},
		},
		{
			source: "+++++[-]",
			expected: &ast.Program{
				Expressions: []ast.Expression{
					&ast.ValueChangeExpression{
						Count: 5,
						Expressions: []ast.Expression{
							&ast.ValueIncrementExpression{Pos: 0},
							&ast.ValueIncrementExpression{Pos: 1},
							&ast.ValueIncrementExpression{Pos: 2},
							&ast.ValueIncrementExpression{Pos: 3},
							&ast.ValueIncrementExpression{Pos: 4},
						},
					},
					&ast.ValueResetExpression{Pos: 5},
				},
			},
		},
		{
			source: ">>+>+>[<<]",
			expected: &ast.Program{
				Expressions: []ast.Expression{
					&ast.PointerMoveExpression{
						Count: 2,
						Expressions: []ast.Expression{
							&ast.PointerIncrementExpression{Pos: 0},
							&ast.PointerIncrementExpression{Pos: 1},
						},
					},
					&ast.ValueChangeExpression{
						Count: 1,
						Expressions: []ast.Expression{
							&ast.ValueIncrementExpression{Pos: 2},
						},
					},
					&ast.PointerMoveExpression{
						Count: 1,
						Expressions: []ast.Expression{
							&ast.PointerIncrementExpression{Pos: 3},
						},
					}, &ast.ValueChangeExpression{
						Count: 1,
						Expressions: []ast.Expression{
							&ast.ValueIncrementExpression{Pos: 4},
						},
					},
					&ast.PointerMoveExpression{
						Count: 1,
						Expressions: []ast.Expression{
							&ast.PointerIncrementExpression{Pos: 5},
						},
					},
					&ast.ZeroSearchExpression{
						StartPosition: 6,
						EndPosition:   9,

						SearchWindow: -2,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.source, func(t *testing.T) {
			p, err := parser.Parse(strings.NewReader(tc.source))
			if err != nil {
				t.Fatal(err)
			}

			o := optimizer.NewOptimizer()
			prog, err := o.Optimize(p)
			if err != nil {
				t.Fatal(err)
			}

			if prog.String() != tc.expected.String() {
				t.Errorf("got: %v, expected: %v", prog.String(), tc.expected.String())
			}

			if !reflect.DeepEqual(prog, tc.expected) {
				t.Errorf("got: %#v, expected: %#v", prog, tc.expected)
			}
		})
	}
}
