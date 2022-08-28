package main

import "fmt"

func main() {
	var numeros []int = make([]int, 3, 3)
	numeros[0] = 100
	numeros[0] = 200
	numeros[0] = 300
	fmt.Printf("Len: %d, Cap: %d, %p\n", len(numeros), cap(numeros), &numeros)
	numeros = append(numeros, 400)
	fmt.Printf("Len: %d, Cap: %d, %p\n", len(numeros), cap(numeros), &numeros)
}
