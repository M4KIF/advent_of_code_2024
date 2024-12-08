package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DataProviderMock struct {
	FirstArray  []int
	SecondArray [][]int
}

var (
	OPERATORS = []string{"+", "*", "||"}
)

func (d DataProviderMock) TakeInput(path string) bool {
	return true
}

func (d DataProviderMock) GetFirstArray() []int {
	return d.FirstArray
}

func (d DataProviderMock) GetSecondArray() [][]int {
	return d.SecondArray
}

func TestPermutationPart1(t *testing.T) {
	result_array := []int{100}
	elements_array := [][]int{{50, 50}}

	expected := []string{"1", "2"}

	mockProvider := DataProviderMock{FirstArray: result_array, SecondArray: elements_array}
	sol := NewSolution(mockProvider, "")
	permutations := []string{}
	temp := ""

	sol.Permutations(&permutations, temp, len(elements_array[0])-1, PART_1)
	assert.Equal(t, permutations, expected)
}

func TestPermutationPart2(t *testing.T) {
	result_array := []int{100}
	elements_array := [][]int{{50, 50}}

	expected := []string{"1", "2", "3"}

	mockProvider := DataProviderMock{FirstArray: result_array, SecondArray: elements_array}
	sol := NewSolution(mockProvider, "")
	permutations := []string{}
	temp := ""

	sol.Permutations(&permutations, temp, len(elements_array[0])-1, PART_2)
	assert.Equal(t, permutations, expected)
}

// Yes
func TestPermutationPart2Long(t *testing.T) {
	result_array := []int{100}
	elements_array := [][]int{{50, 50, 200, 500, 100, 22, 3, 54, 6}}

	expected := []string{"1", "2", "3"}

	mockProvider := DataProviderMock{FirstArray: result_array, SecondArray: elements_array}
	sol := NewSolution(mockProvider, "")
	permutations := []string{}
	temp := ""

	sol.Permutations(&permutations, temp, len(elements_array[0])-1, PART_2)
	assert.NotEqual(t, permutations, expected)
}

func TestPart2SingleLineSumCheck(t *testing.T) {
	result_array := []int{100}
	elements_array := [][]int{{50, 50}}

	mockProvider := DataProviderMock{FirstArray: result_array, SecondArray: elements_array}
	sol := NewSolution(mockProvider, "")
	assert.Equal(t, 100, sol.SolvePart1())
}

func TestPart1LineSumCheckForAFiveHun(t *testing.T) {
	result_array := []int{100, 100, 50, 250}
	elements_array := [][]int{{5, 25, 5, 15, 50}, {100}, {2, 2, 2, 2, 2, 10, 15, 5, 10}, {50, 50, 50, 50, 50}}

	mockProvider := DataProviderMock{FirstArray: result_array, SecondArray: elements_array}
	sol := NewSolution(mockProvider, "")
	assert.Equal(t, 500, sol.SolvePart1())
}

func TestPart2LineSumCheckForAFiveHun(t *testing.T) {
	result_array := []int{100, 100, 50, 250, 156}
	elements_array := [][]int{{5, 25, 5, 15, 50}, {100}, {2, 2, 2, 2, 2, 10, 15, 5, 10}, {50, 50, 50, 50, 50}, {15, 6}}

	mockProvider := DataProviderMock{FirstArray: result_array, SecondArray: elements_array}
	sol := NewSolution(mockProvider, "")
	assert.Equal(t, 656, sol.SolvePart2())
}
