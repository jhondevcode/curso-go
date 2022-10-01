package main

import "fmt"

func main() {
	fmt.Println("Inicio del programa")

	func() {
		fmt.Println("Hola desde funcion anonima")
	}()

	MyFunc := func(nombre string) string {
		return fmt.Sprintf("Hola %s, desde la terminal\n", nombre)
	}

	fmt.Println(MyFunc("Lenovo"))
	fmt.Printf("%T\n", MyFunc)

	var Suma func(...int) int = func(vars ...int) int {
		var sum int = 0
		for _, v := range vars {
			sum += v
		}
		return sum
	}

	fmt.Println(Suma(10, 20, 30, 40, 50))
	fmt.Printf("%T\n", Suma)

	fmt.Println("Terminando programa")
}
