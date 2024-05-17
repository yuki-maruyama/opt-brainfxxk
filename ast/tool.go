package ast

import (
	"fmt"
	"reflect"
)

var nodeCountMap = make(map[string]int)

func printAST(expr Expression, indent string, isLast bool) {
	nodeType := reflect.TypeOf(expr).Elem().Name()

	nodeCountMap[nodeType]++

	fmt.Printf("%s", indent)
	if isLast {
		fmt.Printf("└─ ")
		indent += "   "
	} else {
		fmt.Printf("├─ ")
		indent += "|  "
	}
	fmt.Printf("%s: %s\n", nodeType, expr.String())

	if whileExpr, ok := expr.(*WhileExpression); ok {
		for i, child := range whileExpr.Body {
			printAST(child, indent, i == len(whileExpr.Body)-1)
		}
	}
}

func PrintASTList(exprList []Expression) {
	for i, expr := range exprList {
		printAST(expr, "", i == len(exprList)-1)
	}
}