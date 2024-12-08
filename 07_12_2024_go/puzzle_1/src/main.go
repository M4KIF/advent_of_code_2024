package main

import (
	"github.com/M4KIF/advent_of_code_2024/07_12_2024_go/puzzle_1/src/data_provider"
	"github.com/M4KIF/advent_of_code_2024/07_12_2024_go/puzzle_1/src/solution"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/file_handling"
)

func main() {
	// Initialising the dependencies
	io_provider := file_handling.Default{}
	data_provider := data_provider.NewDataProvider(io_provider)

	// Instantiating the main 'business' logic
	solution := solution.NewSolution(data_provider, "/input_data/data_final.txt")

	// Doing the job
	print("result: ", solution.Solve())
}
