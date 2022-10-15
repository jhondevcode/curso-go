package main

import (
	"fmt"
	"paquetes/figuras"
	"paquetes/mensajes"
	"paquetes/models"

	"github.com/donvito/hellomod"
)

func main() {
	mensajes.Hola()
	mensajes.Imprimir()

	cuadrado := figuras.Cuadrado{
		Lado: 8,
	}
	figuras.Medidas(&cuadrado)

	circulo := figuras.Circulo{
		Radio: 5,
	}
	figuras.Medidas(&circulo)

	persona := models.Persona{}
	persona.Constructor("Alex", 26)
	fmt.Println(persona.ToString())

	persona.SetEdad(27)
	persona.SetNombre("Roel")
	fmt.Println(persona.ToString())

	hellomod.SayHello()
}
