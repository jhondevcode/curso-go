package main

import "fmt"

func Sumar(nombre string, data ...int) (string, int) {
	// data []int
	mensaje := fmt.Sprintf("La suma de %s es:", nombre)
	var suma int = 0
	for _, value := range data {
		suma += value
	}
	return mensaje, suma
}

func main() {
	mensaje, result := Sumar("Alex", 1, 2, 3, 4)
	fmt.Println(mensaje, result)
	fmt.Println(Sumar("Roel", 10, -20, 30, -40))
}
