package main

import "fmt"

func main() {
	var a int32
	fmt.Print("a: ")
	fmt.Scanf("%d", &a)

	var b int32
	fmt.Print("b: ")
	fmt.Scanf("%d", &b)

	suma := a + b
	fmt.Printf("La suma de %d y %d es %d\n", a, b, suma)
}
