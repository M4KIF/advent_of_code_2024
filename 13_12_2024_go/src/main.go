package main

import (
	"github.com/M4KIF/advent_of_code_2024/13_12_2024_go/src/data_provider"
	"github.com/M4KIF/advent_of_code_2024/13_12_2024_go/src/solution"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/file_handling"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/logging"
)

func main() {
	// Deps
	crud := file_handling.Default{}
	provider := data_provider.NewDataProvider(crud)
	solution := solution.NewSolution(provider, "/data.txt")

	logging.Info("Result for part 1", "value", solution.Part1())
}
