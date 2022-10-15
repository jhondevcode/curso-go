package figuras

import "math"

type Cuadrado struct {
	Lado float64
}

func (c *Cuadrado) area() float64 {
	return math.Pow(c.Lado, 2)
}

func (c *Cuadrado) perimetro() float64 {
	return 4 * c.Lado
}
