package utils

import (
	"os"
)

func ReadFile() {
	for i := 0; i < 5; i++ {
		b, _ := os.Open("Makefile")
		defer func() { b.Close() }()
	}
}
