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
			source:  "+++++----->>>>><<<<<",
			expected: &ast.Program{
				Expressions: []ast.Expression{
					&ast.MultipleValueIncrementExpression{Count: 5},
					&ast.MultipleValueDecrementExpression{Count: 5},
					&ast.MultiplePointerIncrementExpression{Count: 5},
					&ast.MultiplePointerDecrementExpression{Count: 5},
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
				t.Errorf("got: %v, expected: %v", prog, tc.expected)
			}
		})
	}
}
