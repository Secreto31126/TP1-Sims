package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StaticInfo struct {
	TotalParticles int
	AreaLength     float64
	Radii          []float64
	Properties     []float64
}

func parseStaticFile(path string) (StaticInfo, error) {
	file, err := os.Open(path)
	if err != nil {
		return StaticInfo{}, fmt.Errorf("failed to open static file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var info StaticInfo

	if !scanner.Scan() {
		return info, fmt.Errorf("file empty or invalid")
	}
	info.TotalParticles, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return info, fmt.Errorf("invalid particle count: %w", err)
	}

	if !scanner.Scan() {
		return info, fmt.Errorf("missing area length")
	}
	info.AreaLength, err = strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
	if err != nil {
		return info, fmt.Errorf("invalid area length: %w", err)
	}

	info.Radii = make([]float64, info.TotalParticles)
	info.Properties = make([]float64, info.TotalParticles)

	for i := 0; i < info.TotalParticles && scanner.Scan(); i++ {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			return info, fmt.Errorf("invalid format for particle %d", i+1)
		}
		info.Radii[i], err = strconv.ParseFloat(fields[0], 64)
		if err != nil {
			return info, fmt.Errorf("invalid radius for particle %d: %w", i+1, err)
		}
		info.Properties[i], err = strconv.ParseFloat(fields[1], 64)
		if err != nil {
			return info, fmt.Errorf("invalid property for particle %d: %w", i+1, err)
		}
	}

	return info, nil
}
