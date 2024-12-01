package main

import (
	"github.com/M4KIF/advent_of_code_2024/01_12_2024_go/puzzle_2/src/data_reader"
	"github.com/M4KIF/advent_of_code_2024/01_12_2024_go/puzzle_2/src/solution"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/logging"
)

func main() {
	// Getting the data with dependency injection
	data := data_reader.Create()

	// Running the sollution
	solution := solution.Solution{}
	logging.Info("Result", "value", solution.Solve(data))
}
