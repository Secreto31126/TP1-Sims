package types

import "math"

type Particle struct {
	Id        int        // Identificador de la particula
	X, Y      float64    // Posicion de la particula
	Radius    float64    // Radio de la particula
	Property  float64    // Propiedad de la particula
	Neighbors []Particle // Vecinos de la particula
}

func (p Particle) BorderDistanceTo(other Particle) float64 {
	dx := p.X - other.X
	dy := p.Y - other.Y
	return math.Sqrt(dx*dx+dy*dy) - (p.Radius + other.Radius)
}
