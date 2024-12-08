package solution

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type IoMock struct{}

func (io *IoMock) Open(s string) (*os.File, error) {
	return nil, nil
}

type DataMock struct {
	FirstArray  []int
	SecondArray []int
}

func (d DataMock) TakeInput(path string) bool {
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

	//data := DataMock{FirstArray: left, SecondArray: right}

	solution := Solution{DataMock{FirstArray: left, SecondArray: right}}
	assert.Equal(t, expected_result, solution.Solve("path"))

}

func TestSolutionWithSmallInputWithRepetition(t *testing.T) {

	// Mocking the data with precalculated source of truth
	left := []int{1, 2, 4, 5, 1, 5}
	right := []int{1, 1, 1, 1, 2, 5, 5, 5}
	expected_result := 40

	solution := Solution{DataMock{FirstArray: left, SecondArray: right}}
	assert.Equal(t, expected_result, solution.Solve("path"))

}
