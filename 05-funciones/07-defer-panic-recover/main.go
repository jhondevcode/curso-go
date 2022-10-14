package main

import (
	"fmt"
	"os"
)

func main() {
	// Function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Ups!, al parecer el programa finalizo de manera inesperada")
		}
	}()

	if file, err := os.Open("hola.txt"); err != nil {
		panic("No fue posible abrir el archivo")
	} else {
		defer func() {
			fmt.Println("Cerrando el archivo")
			file.Close()
		}()
		content := make([]byte, 256)
		long, _ := file.Read(content)
		f_content := string(content[:long])
		fmt.Println(f_content)
	}
}
