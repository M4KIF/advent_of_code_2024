package main

import (
	"fmt"

	"github.com/M4KIF/advent_of_code_2024/01_12_2024_go/puzzle_1/src/data_reader"
	"github.com/M4KIF/advent_of_code_2024/01_12_2024_go/puzzle_1/src/solution"
)

func main() {
	// Getting the data
	data := data_reader.Data{}
	data.Read("/input_data/data.txt")

	// Running the sollution
	solution := solution.Solution{}
	fmt.Println("Result is: ", solution.Solve(data))
}
