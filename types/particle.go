package types

import (
	"math"
	"sync"
)

type Particle struct {
	Id        int         // Identificador de la particula
	X, Y      float64     // Posicion de la particula
	Radius    float64     // Radio de la particula
	Property  float64     // Propiedad de la particula
	Neighbors []*Particle // Vecinos de la particula
	mutex     *sync.Mutex
}

func NewParticle(id int, x, y, radius, property float64) *Particle {
	return &Particle{
		Id:        id,
		X:         x,
		Y:         y,
		Radius:    radius,
		Property:  property,
		Neighbors: make([]*Particle, 0),
		mutex:     &sync.Mutex{},
	}
}

func (p *Particle) BorderDistanceTo(other *Particle, L float64, loop bool) float64 {
	dx := math.Abs(p.X - other.X)
	dy := math.Abs(p.Y - other.Y)

	if loop {
		if dx > L/2 {
			dx = L - dx
		}
		if dy > L/2 {
			dy = L - dy
		}
	}

	return math.Max(math.Sqrt(dx*dx+dy*dy)-(p.Radius+other.Radius), 0)
}

// Add a neighbor to the particle's neighbors list
// It doesn't validate if the neighbor is already present
func (p *Particle) AddNeighbor(neighbor *Particle) {
	if neighbor == nil || neighbor.Id == p.Id {
		return
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.Neighbors = append(p.Neighbors, neighbor)
}
