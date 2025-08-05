package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"TP1-Sims/parser"
	"TP1-Sims/types"
)

type Job struct {
	types.Coordinate
	Particle *types.Particle
}

func main() {
	if len(os.Args) < 4 {
		fmt.Print("Usage: go run main.go <M> <Rc> <Size>")
		os.Exit(1)
	}

	M, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid M value: %v", err)
		os.Exit(1)
	}

	Rc, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Printf("Invalid Rc value: %v", err)
		os.Exit(1)
	}

	timestamps, info, err := parser.ParseFiles(os.Args[3])
	if err != nil {
		fmt.Printf("Error parsing files: %v", err)
		os.Exit(1)
	}

	for _, ts := range timestamps {
		log.Printf("Processing timestamp: %f", ts.Time)

		list := ts.Particles
		processList(list, info.AreaLength, M, Rc)
		output(list)

		log.Printf("Processed timestamp %f", ts.Time)
	}
}

func processList(list []*types.Particle, L float64, M int, Rc float64) {
	var cells [][]([]*types.Particle)
	for i := range M {
		cells = append(cells, make([][]*types.Particle, M))
		for j := range M {
			cells[i][j] = make([]*types.Particle, 0)
		}
	}

	for i := range list {
		p := list[i]
		cellX := int(p.X / (L / float64(M)))
		cellY := int(p.Y / (L / float64(M)))

		cells[cellX][cellY] = append(cells[cellX][cellY], p)
	}

	var wg sync.WaitGroup
	jobs := make(chan Job, len(list))

	for range 1000 {
		wg.Add(1)

		go func() {
			defer wg.Done()
			processParticle(M, Rc, cells, jobs)
		}()
	}

	for i := range cells {
		for j := range cells[i] {
			for _, particle := range cells[i][j] {
				jobs <- Job{
					Coordinate: types.Coordinate{X: i, Y: j},
					Particle:   particle,
				}
			}
		}
	}

	close(jobs)
	wg.Wait()
}

func processParticle(M int, Rc float64, cells [][][]*types.Particle, jobs <-chan Job) {
	for job := range jobs {
		i, j, particle := job.X, job.Y, job.Particle

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
					// This fortunately doesn't cause a deadlock
					particle.AddNeighbor(other)
					other.AddNeighbor(particle)
				}
			}
		}
	}
}

func output(list []*types.Particle) {
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
