package parser

import "fmt"

func ParseFiles() ([]TimeStep, *StaticInfo, error) {
	staticFile := "./files/Static100.txt"
	dynamicFile := "./files/Dynamic100.txt"

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
