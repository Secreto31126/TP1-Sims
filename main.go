package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
	"time"

	"TP1-Sims/parser"
	"TP1-Sims/types"
)

type Job struct {
	types.Coordinate
	Particle *types.Particle
}

func main() {
	start := time.Now()
	if len(os.Args) < 4 {
		fmt.Print("Usage: go run main.go <M> <Rc> <Size> [cell|brute] [loop]\n")
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

	searchMode := "cell"
	if len(os.Args) >= 5 {
		searchMode = os.Args[4]
		if searchMode != "cell" && searchMode != "brute" {
			fmt.Println("Invalid search mode. Use 'cell' or 'brute'")
			os.Exit(1)
		}
	}

	loop := false
	if len(os.Args) >= 6 && os.Args[5] == "loop" {
		loop = true
	}

	var listProcessingTime time.Duration = 0
	for _, ts := range timestamps {
		log.Printf("Processing timestamp: %f", ts.Time)

		list := ts.Particles
		checkpointTime := time.Now()
		processList(list, info.AreaLength, M, Rc, loop, searchMode)
		listProcessingTime += time.Since(checkpointTime)
		output(list)

		log.Printf("Processed timestamp %f", ts.Time)
	}
	log.Printf("List processed in time: %s", listProcessingTime)
	log.Printf("Total runtime: %s", time.Since(start))
}

func processList(list []*types.Particle, L float64, M int, Rc float64, loop bool, searchMode string) {
	for _, p := range list {
		p.Neighbors = nil
	}

	switch searchMode {
	case "cell":
		processCellIndex(list, L, M, Rc, loop)
	case "brute":
		processBruteForce(list, Rc, loop, L)
	}
}

func processCellIndex(list []*types.Particle, L float64, M int, Rc float64, loop bool) {
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
			processParticle(M, Rc, cells, jobs, loop)
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

func processBruteForce(particles []*types.Particle, Rc float64, loop bool, L float64) {
	var wg sync.WaitGroup
	for i := range particles {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			pi := particles[i]
			for j := range particles {
				if i == j {
					continue
				}
				pj := particles[j]

				dist := pi.BorderDistanceTo(pj)
				if loop {
					// Apply minimum image convention for wrapping
					dx := math.Abs(pi.X - pj.X)
					if dx > L/2 {
						dx = L - dx
					}

					dy := math.Abs(pi.Y - pj.Y)
					if dy > L/2 {
						dy = L - dy
					}

					dist = math.Sqrt(dx*dx+dy*dy) - pi.Radius - pj.Radius
				}

				if dist < Rc {
					// Avoid duplicate additions
					if pi.Id < pj.Id {
						pi.AddNeighbor(pj)
						pj.AddNeighbor(pi)
					}
				}
			}
		}(i)
	}
	wg.Wait()
}

func processParticle(M int, Rc float64, cells [][][]*types.Particle, jobs <-chan Job, loop bool) {
	for job := range jobs {
		i, j, particle := job.X, job.Y, job.Particle

		neighbors := []types.Coordinate{
			{X: i - 1, Y: j - 1}, {X: i - 1, Y: j}, {X: i - 1, Y: j + 1},
			{X: i, Y: j - 1}, {X: i, Y: j},
		}

		for curr, coord := range neighbors {
			x, y := coord.X, coord.Y
			self_quadrant := curr == 4

			if loop {
				x = (x + M) % M
				y = (y + M) % M
			} else {
				if x < 0 || x >= M || y < 0 || y >= M {
					continue
				}
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
	for _, p := range list {
		fmt.Printf("%f %f %f\n%d", p.Radius, p.X, p.Y, p.Id)
		for _, neighbor := range p.Neighbors {
			fmt.Printf("\t%d", neighbor.Id)
		}
		fmt.Println()
	}
}
