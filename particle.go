package main

type Particle struct {
	id        int        // Identificador de la particula
	X, Y      float64    // Posicion de la particula
	radius    float64    // Radio de la particula
	property  int        // Propiedad de la particula
	neighbors []Particle // Vecinos de la particula
}
