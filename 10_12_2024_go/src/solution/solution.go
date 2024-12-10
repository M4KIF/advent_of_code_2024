package solution

import (
	interfaces "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/data"
)

type Solution struct {
	DataProvider interfaces.String2DArray
}

func NewSolution(dataProvider interfaces.String2DArray) *Solution {
	return &Solution{DataProvider: dataProvider}
}

func helper(d [][]string) int {
	return 1
}

func (s *Solution) Solve() int {
	// Collecting the data
	s.DataProvider.TakeInput("")

	// Calculating
	data := s.DataProvider.Get2DArray()

	return helper(data)
}
