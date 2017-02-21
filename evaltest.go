package main

import (
	"fmt"
	"github.com/adonovan/gopl.io/ch7/eval"
	//	"math"
)

func TestEval() {
	tests := []struct {
		expr string
		env  eval.Env
		want string
	}{
		{"1 + 2 * 3 * 5", eval.Env{}, "-4"},
		//		{"sqrt(A / pi)", eval.Env{"A": 87616, "pi": math.Pi}, "167"},
		//		{"pow(x, 3) + pow(y, 3)", eval.Env{"x": 12, "y": 1}, "1729"},
		//		{"pow(x, 3) + pow(y, 3)", eval.Env{"x": 9, "y": 10}, "1729"},
		//		{"5 / 9 * (F - 32)", eval.Env{"F": -40}, "-40"},
		//		{"5 / 9 * (F - 32)", eval.Env{"F": 32}, "0"},
		//		{"5 / 9 * (F - 32)", eval.Env{"F": 212}, "100"},
		//		//!-Eval
		//		// additional tests that don't appear in the book
		//		{"-1 + -x", eval.Env{"x": 1}, "-2"},
		//		{"-1 - x", eval.Env{"x": 1}, "-2"},
		//		//!+Eval
	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := eval.Parse(test.expr)
		if err != nil {
			fmt.Printf("Parse error %v\n", err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			fmt.Printf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}

func main() {
	TestEval()
}
