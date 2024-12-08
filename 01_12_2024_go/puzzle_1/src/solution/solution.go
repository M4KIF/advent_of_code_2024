package solution

import (
	"math"
	"slices"

	"github.com/M4KIF/advent_of_code_2024/01_12_2024_go/puzzle_1/src/io"
)

type Solution struct{}

func (s *Solution) Solve(data io.Data) int {
	result := 0

	// Taking pointers into variables for readability
	left := data.GetLeftArray()
	right := data.GetRightArray()

	// Sorting in ascending order
	slices.Sort(left)
	slices.Sort(right)

	// Both arrays should be of equal size
	for i := 0; i < len(left); i++ {
		result += int(math.Abs(float64(left[i] - right[i])))
	}

	return result
}
