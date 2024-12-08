package main

import (
	"github.com/M4KIF/advent_of_code_2024/01_12_2024_go/puzzle_2/src/data_reader"
	"github.com/M4KIF/advent_of_code_2024/01_12_2024_go/puzzle_2/src/solution"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/file_handling"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/logging"
)

func main() {
	// Utilising a default file handling provider
	file_io := file_handling.Default{}

	// Getting the data with dependency injection
	data := data_reader.Create(file_io)

	// Running the sollution
	solution := solution.Solution{data}
	logging.Info("Result", "value", solution.Solve("/input_data/data.txt"))
}
