package main

import "fmt"

func ModificarValor(cadena *string) {
	fmt.Printf("%p\n", cadena)
	*cadena = "Hola desde la funcion"
}

func main() {
	var cadena string = "Hola Mundo desde GO"
	fmt.Printf("%p\n", &cadena)
	fmt.Println("Antes de la funcion:", cadena)

	ModificarValor(&cadena)
	fmt.Println("Despues de la funcion: ", cadena)
	fmt.Printf("%p\n", &cadena)
}
