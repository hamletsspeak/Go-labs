package main

import (
	"github.com/hamletsspeak/Oblako52/internal/worker"
)

func main() {
	worker.RunWorkers(3)
}
