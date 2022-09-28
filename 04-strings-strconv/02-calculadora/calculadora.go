package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Evaluar(input string) int32 {
	caracteres := strings.Split(input, "+")
	var suma int32 = 0
	for _, caracter := range caracteres {
		dato, err := strconv.Atoi(caracter)
		if err != nil {
			fmt.Println(err)
			suma += 0
		} else {
			suma += int32(dato)
		}
	}
	return suma
}

func main() {
	var expresion string
	fmt.Print("=>")
	fmt.Scanln(&expresion)
	fmt.Println(Evaluar(expresion))
}
