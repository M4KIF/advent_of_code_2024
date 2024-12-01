package solution

import (
	"testing"

	icrud "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/file"
	"github.com/stretchr/testify/assert"
)

type DataMock struct {
	FirstArray  []int
	SecondArray []int
}

func (d DataMock) TakeInput(path string, file_io icrud.CRUD) bool {
	return true
}

func (d DataMock) GetFirstArray() []int {
	return d.FirstArray
}

func (d DataMock) GetSecondArray() []int {
	return d.SecondArray
}

func TestSolutionWithSmallInputNoRepetition(t *testing.T) {

	// Mocking the data with precalculated source of truth
	left := []int{1, 2, 4, 5}
	right := []int{1, 1, 1, 1, 2, 5, 5, 5}
	expected_result := 21

	data := DataMock{FirstArray: left, SecondArray: right}

	solution := Solution{}
	assert.Equal(t, expected_result, solution.Solve(data))

}

func TestSolutionWithSmallInputWithRepetition(t *testing.T) {

	// Mocking the data with precalculated source of truth
	left := []int{1, 2, 4, 5, 1, 5}
	right := []int{1, 1, 1, 1, 2, 5, 5, 5}
	expected_result := 40

	data := DataMock{FirstArray: left, SecondArray: right}

	solution := Solution{}
	assert.Equal(t, expected_result, solution.Solve(data))

}
