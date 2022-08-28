package main

import "fmt"

func saludar(nombre string) {
	fmt.Println("Hola,", nombre)
}

func sumar(a, b int) int {
	return a + b
}

func main() {
	saludar("Alex")
	resultado := sumar(10, 20)
	fmt.Printf("Valor de la suma: %d\n", resultado)
}
