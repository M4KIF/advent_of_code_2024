package sollution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DataMock struct {
	LeftArray  []int
	RightArray []int
}

func (d DataMock) GetLeftArray() []int {
	return d.LeftArray
}

func (d DataMock) GetRightArray() []int {
	return d.RightArray
}

func TestSollutionWithSmallInputNoRepetition(t *testing.T) {

	// Mocking the data with precalculated source of truth
	left := []int{1, 2, 4, 5}
	right := []int{1, 1, 1, 1, 2, 5, 5, 5}
	expected_result := 21

	data := DataMock{LeftArray: left, RightArray: right}

	sollution := Sollution{}
	assert.Equal(t, expected_result, sollution.Solve(data))

}

func TestSollutionWithSmallInputWithRepetition(t *testing.T) {

	// Mocking the data with precalculated source of truth
	left := []int{1, 2, 4, 5, 1, 5}
	right := []int{1, 1, 1, 1, 2, 5, 5, 5}
	expected_result := 40

	data := DataMock{LeftArray: left, RightArray: right}

	sollution := Sollution{}
	assert.Equal(t, expected_result, sollution.Solve(data))

}
