package main

import (
	"fmt"

	"github.com/M4KIF/advent_of_code_2024/01_12_2024/iter_1/data_reader"
	"github.com/M4KIF/advent_of_code_2024/01_12_2024/iter_1/sollution"
)

func main() {
	// Getting the data
	data := data_reader.Data{}
	data.Read("/input_data/data.txt")

	// Running the sollution
	sollution := sollution.Sollution{}
	fmt.Println("Result is: ", sollution.Solve(data))
}
