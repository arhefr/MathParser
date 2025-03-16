package mathparser

import (
	"fmt"
	"slices"
)

func Parser(expr string) []string {
	var (
		ent  string
		exit []string
	)

	expr = fix(expr)
	for i := range expr {
		sym := string(expr[i])

		if _, ok := token[sym]; !ok || (i == 0 && (sym == "-" || sym == "+")) {
			ent += sym

			if i == len(expr)-1 {
				exit = append(exit, ent)
				ent = ""
				break
			}

			if _, ok := token[string(expr[i+1])]; ok {
				exit = append(exit, ent)
				ent = ""
			}
			continue
		}
		exit = append(exit, sym)
	}

	return exit
}

func InfixPrefix(exprStr string) []string {
	expr := Parser(exprStr)
	slices.Reverse(expr)

	for i, ent := range expr {
		if ent == ")" {
			expr[i] = "("
		}

		if ent == "(" {
			expr[i] = ")"
		}
	}

	res := train(expr)
	slices.Reverse(res)
	return res
}

func InfixPostfix(exprStr string) []string {
	return train(Parser(exprStr))
}

func PrefixInfix(expr []string) (string, error) {
	slices.Reverse(expr)
	exprStr, err := PostfixInfix(expr)
	if err != nil {
		return "", err
	}

	return exprStr, nil
}

func PostfixInfix(expr []string) (string, error) {
	var res []string

	for _, ent := range expr {
		if _, ok := token[ent]; !ok {
			res = slices.Insert(res, 0, ent)
			continue
		}

		if len(res) < 2 {
			return "", fmt.Errorf("error incorrect math expression")
		}

		op1, op2 := res[0], res[1]
		res = res[2:]
		res = slices.Insert(res, 0, "("+op2+ent+op1+")")
	}

	return res[0], nil
}
