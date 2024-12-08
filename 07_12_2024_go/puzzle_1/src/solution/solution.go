package solution

import (
	"strconv"

	interfaces "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/data"
)

var (
	ADDITION       = "1"
	MULTIPLICATION = "2"
	CONCAT         = "3"

	PART_1 = 2
	PART_2 = 3
)

type Solution struct {
	DataProvider interfaces.SingleDimIntTwoDimIntArrays
}

func NewSolution(dataProvider interfaces.SingleDimIntTwoDimIntArrays, path string) *Solution {
	sol := &Solution{DataProvider: dataProvider}
	sol.DataProvider.TakeInput(path)
	return sol
}

func (s *Solution) Permutations(res *[]string, temp string, lenght int, mode int) {

	if len(temp) == lenght {
		*res = append((*res), temp)
		return
	}

	for i := 0; i < mode; i++ {
		temp += strconv.Itoa((i % mode) + 1)

		s.Permutations(res, temp, lenght, mode)

		temp = temp[:len(temp)-1]
	}
}

func (s *Solution) SolvePart1() int {
	results := s.DataProvider.GetFirstArray()
	elements := s.DataProvider.GetSecondArray()

	result := 0
	for i, r := range results {
		// Calculating the subsets of operators that can be used for this problem
		subsets := []string{}
		subset := ""

		s.Permutations(&subsets, subset, len(elements[i])-1, PART_1)

		for _, ss := range subsets {
			temp_result := elements[i][0]

			for j := range ss {
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

func (s *Solution) SolvePart2() int {
	results := s.DataProvider.GetFirstArray()
	elements := s.DataProvider.GetSecondArray()

	result := 0
	for i, r := range results {
		subsets := []string{}
		subset := ""

		s.Permutations(&subsets, subset, len(elements[i])-1, PART_2)

		for _, ss := range subsets {
			temp_result := elements[i][0]

			for j := range ss {
				if string(ss[j]) == ADDITION {
					temp_result += elements[i][1+j]
					//logging.Info("Eins", "tmep", string(ss[j]))
				} else if string(ss[j]) == MULTIPLICATION {
					temp_result *= elements[i][1+j]
					//logging.Info("Duo", "tmep", string(ss[j]))
				} else if string(ss[j]) == CONCAT {
					temp := strconv.Itoa(temp_result)
					temp += strconv.Itoa(elements[i][1+j])
					temp_result, _ = strconv.Atoi(temp)
					//logging.Info("POLICEI!", "tmep", string(ss[j]), "temp", temp, "res", temp_result)
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
