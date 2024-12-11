package main

import (
	"github.com/M4KIF/advent_of_code_2024/10_12_2024_go/src/data_provider"
	"github.com/M4KIF/advent_of_code_2024/10_12_2024_go/src/solution"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/file_handling"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/logging"
)

func main() {
	io_provider := file_handling.Default{}
	dp := data_provider.NewDataProvider(io_provider)

	solution := solution.NewSolution(dp, "/data.txt")
	logging.Info("Part1 result", "value", solution.Solve())
}
