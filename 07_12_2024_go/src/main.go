package main

import (
	"fmt"

	"github.com/M4KIF/advent_of_code_2024/07_12_2024_go/src/data_provider"
	"github.com/M4KIF/advent_of_code_2024/07_12_2024_go/src/solution"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/file_handling"
)

func main() {
	// Initialising the dependencies
	io_provider := file_handling.Default{}
	data_provider := data_provider.NewDataProvider(io_provider)

	// Instantiating the main 'business' logic
	solution := solution.NewSolution(data_provider, "/input_data/data_final.txt")

	// Doing the job
	fmt.Println("Results part 1: ", solution.SolvePart1())
	fmt.Println("Results part 2: ", solution.SolvePart2())
}
