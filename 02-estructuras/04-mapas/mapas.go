package main

import "fmt"

/**
 * Estructura de un mapa:
 *
 */
func main() {
	dias := make(map[int]string)
	dias[0] = "Lunes"
	dias[1] = "Martes"
	dias[2] = "Miercoles"
	dias[3] = "Jueves"
	dias[4] = "Viernes"
	dias[5] = "Sabado"
	dias[6] = "Domingo"
	fmt.Println(dias)

	dias[7] = "undefined"
	fmt.Println(dias)
	delete(dias, 1)
	fmt.Println(dias, len(dias))

	// Nuevo Mapa
	estudiantes := make(map[string][]int)
	estudiantes["Alex"] = []int{13, 15, 16}
	estudiantes["Roel"] = []int{14, 13, 17}
	fmt.Println(estudiantes)
	fmt.Println(estudiantes["Alex"][1])
}
