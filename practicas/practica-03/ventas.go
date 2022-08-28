package main

import "fmt"

func main() {
	var precio float32

	fmt.Print("Valor de venta: ")
	fmt.Scanf("%f", &precio)

	igv := 18.0 / 100.0

	precio_final := precio + (precio * float32(igv))
	fmt.Printf("El precio final de venta es %f\n", precio_final)
}
