package parser

func ParseFiles() ([]TimeStep, *StaticInfo, error) {
	staticFile := "./files/Static100.txt"
	dynamicFile := "./files/Dynamic100.txt"

	staticInfo, err := parseStaticFile(staticFile)
	if err != nil {
		panic(err)
	}

	timeSteps, err := parseDynamicFile(dynamicFile, *staticInfo)
	if err != nil {
		panic(err)
	}

	return timeSteps, staticInfo, nil
}
