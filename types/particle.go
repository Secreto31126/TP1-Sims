package types

type Particle struct {
	Id        int        // Identificador de la particula
	X, Y      float64    // Posicion de la particula
	Radius    float64    // Radio de la particula
	Property  float64    // Propiedad de la particula
	Neighbors []Particle // Vecinos de la particula
}
