package mensajes

import "fmt"

// constante privada
const message string = "Hola desde mi constante"

/* ***** Funciones publicas ***** */
func Imprimir() {
	fmt.Println(message)
	privateFunction()
}

func Hola() {
	fmt.Println("Hola desde el packete mensajes")
}

/* ***** Funciones privadas ***** */
func privateFunction() {
	fmt.Println("Hola desde la funcion privada")
}
