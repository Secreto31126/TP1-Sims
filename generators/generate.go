package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run main.go <N> <L>")
	}

	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid N value: %v", err)
	}

	L, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid L value: %v", err)
	}

	if len(os.Args) == 3 {
		static(N, L)
	} else {
		dynamic(N, L)
	}
}

func static(N, L int) {
	fmt.Printf("%d\n%d\n", N, L)
	for range N {
		fmt.Printf("%f\t%f\n", rand.Float64()*0.02 + 0.24, rand.Float64()*0.02 + 0.24)
	}
}

func dynamic(N, L int) {
	fmt.Printf("0\n")
	for range N {
		fmt.Printf("%.8e\t%.8e\n", rand.Float64()*float64(L), rand.Float64()*float64(L))
	}
}
