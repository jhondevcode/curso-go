package main

import "fmt"

/**
 * Los slices son estructuras de datos mutables con capacidades de almacenamiento variable
 * Caracteristicas: punteros, longitud, capacidad
 *
 * Array: var data [3]int = [3]int{10, 20, 30}
 * Slice: var data []int = []int{10, 20, 30}
 */
func main() {
	numeros := []int{1, 2, 3}
	fmt.Println(numeros, len(numeros))

	// Agregar numeros
	numeros = append(numeros, 4)
	numeros = append(numeros, 5)
	numeros = append(numeros, 6, 7, 8, 9)

	fmt.Println(numeros, len(numeros))

	// sub slice
	sub_slice := numeros[:2]
	numeros[0] = 100

	fmt.Println(sub_slice)
	fmt.Println(numeros)

	var meses []string = []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio"}
	fmt.Printf("Len: %d, Cap: %v, %p\n", len(meses), cap(meses), &meses)
	meses = append(meses, "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre")
	fmt.Printf("Len: %d, Cap: %v, %p\n", len(meses), cap(meses), &meses)
}
