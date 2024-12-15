package solution

import (
	interfaces "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/data"
)

type Solution struct {
	DataProvider interfaces.Int2DArray
	Path         string
}

func NewSolution(dataProvider interfaces.Int2DArray, path string) *Solution {
	sol := &Solution{DataProvider: dataProvider, Path: path}
	sol.DataProvider.TakeInput(path)
	return sol
}

func (s *Solution) Part1() int {
	return 1
}

func (s *Solution) Part2() int {
	return 1
}
