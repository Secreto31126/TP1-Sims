package types

import "math"

type Particle struct {
	Id        int         // Identificador de la particula
	X, Y      float64     // Posicion de la particula
	Radius    float64     // Radio de la particula
	Property  float64     // Propiedad de la particula
	Neighbors []*Particle // Vecinos de la particula
}

func NewParticle(id int, x, y, radius, property float64) *Particle {
	return &Particle{
		Id:        id,
		X:         x,
		Y:         y,
		Radius:    radius,
		Property:  property,
		Neighbors: make([]*Particle, 0),
	}
}

func (p Particle) BorderDistanceTo(other *Particle) float64 {
	dx := p.X - other.X
	dy := p.Y - other.Y
	return math.Sqrt(dx*dx+dy*dy) - (p.Radius + other.Radius)
}
