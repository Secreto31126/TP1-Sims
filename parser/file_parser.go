package parser

import "fmt"

func ParseFiles(size string) ([]TimeStep, *StaticInfo, error) {
	staticFile := "./files/Static" + size + ".txt"
	dynamicFile := "./files/Dynamic" + size + ".txt"

	staticInfo, err := parseStaticFile(staticFile)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing static file: %w", err)
	}

	timeSteps, err := parseDynamicFile(dynamicFile, *staticInfo)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing dynamic file: %w", err)
	}

	return timeSteps, staticInfo, nil
}
