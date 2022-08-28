package main

import "fmt"

func main() {
	var nombre1 string = "Alex"
	var nombre2 string = "Roel"
	edad := 26
	fmt.Println(nombre1, nombre2, edad)

	var a int     // valor por defecto: 0
	var b float64 // valor por defecto: 0
	var c string  // valor por defecto: ""
	var d bool    // valor por defecto: false

	fmt.Println(a, b, c, d)

	const pi = 3.141592
	fmt.Println(pi)
}
