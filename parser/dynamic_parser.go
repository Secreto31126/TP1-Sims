package parser

import (
	"TP1-Sims/types"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TimeStep struct {
	Time      float64
	Particles []types.Particle
}

func parseDynamicFile(path string, staticInfo StaticInfo) ([]TimeStep, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open dynamic file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var timeSteps []TimeStep

	for scanner.Scan() {
		timeLine := scanner.Text()
		timeVal, err := strconv.ParseFloat(strings.TrimSpace(timeLine), 64)
		if err != nil {
			return nil, fmt.Errorf("expected time value, got: %s", timeLine)
		}

		particles := make([]types.Particle, staticInfo.TotalParticles)
		for i := 0; i < staticInfo.TotalParticles; i++ {
			if !scanner.Scan() {
				return nil, fmt.Errorf("incomplete particle data at time %.2f", timeVal)
			}
			fields := strings.Fields(scanner.Text())
			if len(fields) < 2 {
				return nil, fmt.Errorf("invalid particle data at time %.2f: %v", timeVal, fields)
			}
			x, _ := strconv.ParseFloat(fields[0], 64)
			y, _ := strconv.ParseFloat(fields[1], 64)

			particles[i] = types.Particle{
				Id:       i,
				X:        x,
				Y:        y,
				Radius:   staticInfo.Radii[i],
				Property: staticInfo.Properties[i],
			}
		}
		timeSteps = append(timeSteps, TimeStep{Time: timeVal, Particles: particles})
	}

	return timeSteps, nil
}
