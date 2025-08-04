//go:build prod
// +build prod

package main

import (
	"io"
	"log"
)

func init() {
	log.SetOutput(io.Discard)
}
