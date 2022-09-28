package main

import (
	"fmt"
	"strings"
)

func Reverse(input string) string {
	src := strings.Split(input, "")
	out := make([]string, 0)
	for ptr := len(src) - 1; ptr >= 0; ptr-- {
		out = append(out, src[ptr])
	}
	return strings.Join(out, "")
}

func EsPalindromo(input string) bool {
	input = strings.ReplaceAll(strings.ToLower(input), " ", "")
	return strings.Compare(input, Reverse(input)) == 0
}

func Verificar(input string) {
	if EsPalindromo(input) {
		fmt.Printf("%s es palindromo\n", input)
	} else {
		fmt.Printf("%s no es palindromo\n", input)
	}
}

func main() {
	Verificar("Oso")
	Verificar("Anita lava la tina")
	Verificar("Oso baboso")
}
