package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"TP1-Sims/parser"
	"TP1-Sims/types"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run main.go <M> <Rc>")
	}

	M, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid M value: %v", err)
	}

	Rc, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatalf("Invalid Rc value: %v", err)
	}

	timestamps, info, err := parser.ParseFiles()
	if err != nil {
		log.Fatalf("Error parsing files: %v", err)
	}

	for _, ts := range timestamps {
		log.Printf("Processing timestamp: %f", ts.Time)

		list := ts.Particles
		processList(list, info.AreaLength, M, Rc)
		output(list)

		log.Printf("Processed timestamp %f", ts.Time)
	}
}

func processList(list []types.Particle, L float64, M int, Rc float64) {
	var cells [][]([]*types.Particle)
	for i := range M {
		cells = append(cells, make([][]*types.Particle, M))
		for j := range M {
			cells[i][j] = make([]*types.Particle, 0)
		}
	}

	for i := range list {
		p := &list[i]
		cellX := int(p.X / (L / float64(M)))
		cellY := int(p.Y / (L / float64(M)))

		cells[cellX][cellY] = append(cells[cellX][cellY], p)
	}

	for i := range cells {
		for j := range cells[i] {
			for _, particle := range cells[i][j] {
				processParticle(M, Rc, cells, particle, i, j)
			}
		}
	}
}

func processParticle(M int, Rc float64, cells [][][]*types.Particle, particle *types.Particle, i, j int) {
	neighbors := []types.Coordinate{
		{X: i - 1, Y: j - 1}, {X: i - 1, Y: j}, {X: i - 1, Y: j + 1},
		{X: i, Y: j - 1}, {X: i, Y: j},
	}

	for curr, coord := range neighbors {
		x, y := coord.X, coord.Y
		self_quadrant := curr == 4

		if x < 0 || x >= M || y < 0 || y >= M {
			continue
		}

		for _, other := range cells[x][y] {
			if (!self_quadrant || particle.Id < other.Id) && particle.BorderDistanceTo(other) < Rc {
				particle.Neighbors = append(particle.Neighbors, other)
				other.Neighbors = append(other.Neighbors, particle)
			}
		}
	}
}

func output(list []types.Particle) {
	// for _, p := range list {
	// 	fmt.Printf("%f %f %f\n", p.Radius, p.X, p.Y)
	// }

	for _, p := range list {
		fmt.Printf("%d", p.Id)
		for _, neighbor := range p.Neighbors {
			fmt.Printf("\t%d", neighbor.Id)
		}
		fmt.Println()
	}
}
