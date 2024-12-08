package solution

import (
	interfaces "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/data"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/logging"
)

var (
	ADDITION       = "+"
	MULTIPLICATION = "*"
)

type Solution struct {
	DataProvider interfaces.SingleDimIntTwoDimIntArrays
}

func NewSolution(dataProvider interfaces.SingleDimIntTwoDimIntArrays, path string) *Solution {
	sol := &Solution{DataProvider: dataProvider}
	sol.DataProvider.TakeInput(path)
	return sol
}

func (s *Solution) operSubsets(res *[][]string, oper []string, n int, curr []string) {

	if len(curr) == n {
		*res = append(*res, curr)
		return
	}

	for i := range len(oper) {
		curr = append(curr, oper[i])

		s.operSubsets(res, oper, n, curr)

		curr = curr[:len(curr)-1]
	}
}

func (s *Solution) Solve() int {

	results := s.DataProvider.GetFirstArray()
	elements := s.DataProvider.GetSecondArray()
	operators := []string{"+", "*"}

	result := 0
	for i, r := range results {
		// Calculating the subsets of operators that can be used for this problem
		subsets := [][]string{}
		subset := []string{}

		s.operSubsets(&subsets, operators, len(elements[i])-1, subset)
		logging.Info("Subsets that were calculated out of the given data", "subs", subsets, "result", r)

		for _, ss := range subsets {
			// Iterating over all subsets in search of a true line
			temp_result := elements[i][0]

			for j, o := range ss {
				if o == ADDITION {
					temp_result += elements[i][1+j]
				} else if o == MULTIPLICATION {
					temp_result *= elements[i][1+j]
				}
			}

			if temp_result == r {
				result += r
				break
			}
		}
	}

	return result
}
