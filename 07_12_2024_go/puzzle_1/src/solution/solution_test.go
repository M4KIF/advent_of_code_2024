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
	OPERATORS = []string{"+", "*"}
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

func TestPermutation(t *testing.T) {
	result_array := []int{100}
	elements_array := [][]int{{50, 50}}

	expected := []string{"+", "*"}

	mockProvider := DataProviderMock{FirstArray: result_array, SecondArray: elements_array}
	sol := NewSolution(mockProvider, "")
	permutations := []string{}
	temp := ""

	sol.Permutations(&permutations, OPERATORS, temp, len(elements_array[0])-1)
	assert.Equal(t, permutations, expected)
}

func TestSingleLineSumCheck(t *testing.T) {
	result_array := []int{100}
	elements_array := [][]int{{50, 50}}

	mockProvider := DataProviderMock{FirstArray: result_array, SecondArray: elements_array}
	sol := NewSolution(mockProvider, "")
	assert.Equal(t, 100, sol.Solve())
}

func TestLineSumCheckForAFiveHun(t *testing.T) {
	result_array := []int{100, 100, 50, 250}
	elements_array := [][]int{{5, 25, 5, 15, 50}, {100}, {2, 2, 2, 2, 2, 10, 15, 5, 10}, {50, 50, 50, 50, 50}}

	mockProvider := DataProviderMock{FirstArray: result_array, SecondArray: elements_array}
	sol := NewSolution(mockProvider, "")
	assert.Equal(t, 500, sol.Solve())
}
