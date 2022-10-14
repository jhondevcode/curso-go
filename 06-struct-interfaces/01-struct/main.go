package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
}

type Empleado struct {
	Persona
	sueldo float64
}

// Metodos
func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\nEdad: %d\n", p.nombre, p.edad)
}

func (p *Persona) SetNombre(nombre string) {
	p.nombre = nombre
}

func main() {
	var p1 Persona = Persona{"Alex", 26}
	// fmt.Println(p1)
	p1.imprimir()

	p1.SetNombre("Juan")
	// fmt.Println(p1)
	p1.imprimir()

	var p2 Persona = Persona{
		nombre: "Roel",
		edad:   27,
	}
	// fmt.Println(p2)
	p2.imprimir()
	p2.SetNombre("Carlos")
	p2.imprimir()

	// Uso ded herencia
	var em1 Empleado = Empleado{
		sueldo: 500,
	}
	em1.nombre = "Pedro"
	em1.edad = 30
	em1.imprimir()
	fmt.Println(em1)
}
