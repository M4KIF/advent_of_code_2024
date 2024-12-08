package solution

import (
	interfaces "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/data"
)

type Solution struct {
	DataProvider interfaces.TwoIntArrays
}

func NewSolution(dataProvider interfaces.TwoIntArrays) *Solution {
	return &Solution{dataProvider}
}

func (s *Solution) Solve(path string) int {
	s.DataProvider.TakeInput(path)

	result := 0

	// Taking pointers into variables for readability
	left := s.DataProvider.GetFirstArray()
	right := s.DataProvider.GetSecondArray()

	memo := map[int]int{}

	for i := 0; i < len(left); i++ {

		if entryValue, exists := memo[left[i]]; !exists {
			// If not exists in the memo, calculates and memoises
			appearances := 0

			for j := 0; j < len(right); j++ {
				if left[i] == right[j] {
					appearances++
				}
			}

			// Adding to the memo and appending the result with newest memo entry
			memo[left[i]] = appearances

			result += memo[left[i]] * left[i]
		} else {
			// If exists in the memo skips calc and increments the result
			result += entryValue * left[i]
		}

	}
	return result
}
