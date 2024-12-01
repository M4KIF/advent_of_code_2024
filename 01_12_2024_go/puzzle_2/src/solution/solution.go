package solution

import (
	"github.com/M4KIF/advent_of_code_2024/middleware/go/file_handling"
	interfaces "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/data"
)

type Solution struct{}

func (s *Solution) Solve(data interfaces.TwoIntArrays) int {
	// Utilising a default file handling provider
	file_io := file_handling.Default{}

	// Collecting the input data
	data.TakeInput("/input_data/data.txt", file_io)

	result := 0

	// Taking pointers into variables for readability
	left := data.GetFirstArray()
	right := data.GetSecondArray()

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
