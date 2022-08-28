package main

import "fmt"

/**
 * Los arreglos son inmutable y de tamaño fijo
 */
func main() {
	var numeros [5]int // arreglo inmutable
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30

	fmt.Println(numeros)
	fmt.Println(numeros[1])

	nombres := [4]string{"Alex", "Roel", "Juan", "Te amo lizeth"}
	fmt.Println(nombres)

	colores := [...]string{"Rojo", "Verde", "Negro", "Azul"}
	fmt.Println(colores, len(colores))

	// Arreglos con indices definidos
	monedas := [...]string{0: "Dolar", 2: "Soles", 3: "Pesos", 5: "Euro"}
	fmt.Println(monedas, len(monedas))

	subMoneda := monedas[:3]
	fmt.Println(subMoneda)
}
