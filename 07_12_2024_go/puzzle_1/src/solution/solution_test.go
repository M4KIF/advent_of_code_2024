package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DataProviderMock struct {
	FirstArray  []int
	SecondArray [][]int
}

func (d DataProviderMock) TakeInput(path string) bool {
	return true
}

func (d DataProviderMock) GetFirstArray() []int {
	return d.FirstArray
}

func (d DataProviderMock) GetSecondArray() [][]int {
	return d.SecondArray
}

func SingleLineSumCheck(t *testing.T) {
	result_array := []int{100}
	elements_array := [][]int{{50, 50}}

	mockProvider := DataProviderMock{FirstArray: result_array, SecondArray: elements_array}
	sol := NewSolution(mockProvider, "")

	assert.Equal(t, sol.Solve(), 100)
}
