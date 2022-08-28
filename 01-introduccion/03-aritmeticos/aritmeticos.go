package main

import "fmt"

func main() {
	a := 20
	b := 10

	// Suma
	result := a + b
	fmt.Println(result)

	// Resta
	result = a - b
	fmt.Println(result)

	// Multiplicacion
	result = a * b
	fmt.Println(result)

	// Division
	const div float32 = 3.0 / 2.0
	fmt.Println(div)

	// Modulo de la division
	result = 3 % 2
	fmt.Println(result)
}
