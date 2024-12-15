package main

import (
	"github.com/M4KIF/advent_of_code_2024/middleware/go/file_handling"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/logging"
	"github.com/M4KIF/advent_of_code_2024/project_stubs/go/src/data_provider"
	"github.com/M4KIF/advent_of_code_2024/project_stubs/go/src/solution"
)

func main() {
	// Deps
	crud := file_handling.Default{}
	provider := data_provider.NewDataProvider(crud)
	solution := solution.NewSolution(provider, "/data.txt")

	logging.Info("Result for part 1", "value", solution.Part1())
	logging.Info("Result for part 2", "value", solution.Part2())
}
