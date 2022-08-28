package main

import "fmt"

func main() {
	var a int32
	var b int32

	// cofiente y residuo
	fmt.Print("a: ")
	fmt.Scanf("%d", &a)

	fmt.Print("b: ")
	fmt.Scanf("%d", &b)

	cociente := a / b
	residuo := a % b

	fmt.Printf("El cocienter de dividor %d y %d es %d\n", a, b, cociente)
	fmt.Printf("El residuo de dividir %d y %d es %d\n", a, b, residuo)
}
