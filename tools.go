package mathparser

import (
	"slices"
	"strings"
)

var replace = [][]string{{" ", ""}, {"--", "+"}, {"+-", "-"}, {"**", "^"}, {"-+", "-"}}

func fix(expr string) string {
	for _, rep := range replace {
		expr = strings.ReplaceAll(expr, rep[0], rep[1])
	}
	return expr
}

func add(exit, stack []string, opAdd string) ([]string, []string) {
	if opAdd == "(" {
		return exit, append(stack, opAdd)
	}

	if opAdd == ")" {
		for i := len(stack) - 1; i >= 0; i-- {
			op := stack[i]
			if op == "(" {
				stack = stack[:len(stack)-1]
				return exit, stack
			}

			exit = append(exit, op)
			stack = stack[:len(stack)-1]
		}
	}

	for i := len(stack) - 1; i >= 0; i-- {
		op := stack[i]
		if token[op] < token[opAdd] {
			return exit, append(stack, opAdd)
		}

		exit = append(exit, op)
		stack = stack[:len(stack)-1]
	}

	if len(stack) == 0 {
		return exit, append(stack, opAdd)
	}

	return exit, stack
}

func train(expr []string) []string {
	var (
		exit  []string
		stack []string
	)

	for _, ent := range expr {
		if _, ok := token[ent]; !ok {
			exit = append(exit, ent)
			continue
		}

		exit, stack = add(exit, stack, ent)
	}

	slices.Reverse(stack)
	exit = append(exit, stack...)
	return exit
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
