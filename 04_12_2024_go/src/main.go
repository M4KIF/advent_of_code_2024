package main

import (
	"github.com/M4KIF/advent_of_code_2024/04_12_2024_go/src/data_provider"
	"github.com/M4KIF/advent_of_code_2024/04_12_2024_go/src/solution"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/file_handling"
)

func main() {
	// Initialising the dependencies
	io_provider := file_handling.Default{}
	data_provider := data_provider.NewDataProvider(io_provider)

	// Instantiating the main 'business' logic
	solution := solution.NewSolution(data_provider, "/input_data/input_final.txt")

	// Doing the job
	//print("result part1: ", solution.SolvePart1())
	print("\nresult part2: ", solution.SolvePart2(2, 2))
}
