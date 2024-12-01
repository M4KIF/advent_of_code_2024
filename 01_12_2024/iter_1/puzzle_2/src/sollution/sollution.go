package sollution

import (
	"github.com/M4KIF/advent_of_code_2024/01_12_2024/iter_1/io"
)

type Sollution struct{}

func (s *Sollution) Solve(data io.Data) int {
	result := 0

	// Taking pointers into variables for readability
	left := data.GetLeftArray()
	right := data.GetRightArray()

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
