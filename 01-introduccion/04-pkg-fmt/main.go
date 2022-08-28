package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Alex"
	edad := 26

	// %s -> tipo string
	// %d -> tipo numerico
	// %v -> tipo desconocido
	fmt.Printf("Hola, %s y su edad es %d\n", nombre, edad)
	fmt.Printf("Hola, %v y su edad es %v\n", nombre, edad)

	// Sprintf -> string format
	mensaje := fmt.Sprintf("Hola, %s y su edad es %d", nombre, edad)
	fmt.Println(mensaje)

	// %T -> indica el tipo de datos
	fmt.Printf("nombre: %T\n", nombre)

	fmt.Print("Ingrese otro nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("El otro nombre es: ", nombre)
}
