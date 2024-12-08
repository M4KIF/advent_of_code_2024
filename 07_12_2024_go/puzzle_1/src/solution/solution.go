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

func (s *Solution) Permutations(res *[]string, oper []string, temp string, lenght int) {

	if len(temp) == lenght {
		*res = append((*res), temp)
		return
	}

	for i := 0; i < 2; i++ {
		temp += oper[i%2]

		s.Permutations(res, oper, temp, lenght)

		temp = temp[:len(temp)-1]
	}
}

func (s *Solution) Solve() int {

	results := s.DataProvider.GetFirstArray()
	elements := s.DataProvider.GetSecondArray()
	//operators := []string{"+", "*"}

	result := 0
	for i, r := range results {
		// Calculating the subsets of operators that can be used for this problem
		subsets := []string{}
		operators := []string{"*", "+"}
		subset := ""

		s.Permutations(&subsets, operators, subset, len(elements[i])-1)
		logging.Info("Permutations that were calculated out of the given data", "subs", subsets, "result", r, "elements", elements[i], "len", len(elements[i])-1)

		for _, ss := range subsets {
			// Iterating over all subsets in search of a true line
			temp_result := elements[i][0]

			for j, _ := range ss {
				if string(ss[j]) == ADDITION {
					temp_result += elements[i][1+j]
				} else if string(ss[j]) == MULTIPLICATION {
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
