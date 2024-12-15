package solution

import (
	"testing"

	"github.com/M4KIF/advent_of_code_2024/13_12_2024_go/src/data_provider"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/file_handling"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationPart1CalcBWith480DataSet(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)

	testSolution := NewSolution(dp, path)

	assert.Equal(t, testSolution.calculateClicksB(testSolution.DataProvider.Get2DArray()[0]), float64(40))
}

func TestIntegrationPart1CalcAWith480DataSet(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)

	testSolution := NewSolution(dp, path)

	assert.Equal(t, testSolution.calculateClicksA(testSolution.DataProvider.Get2DArray()[0], float64(40)), float64(80))
}

func TestIntegrationPart1With480(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)

	testSolution := NewSolution(dp, path)

	assert.Equal(t, testSolution.Part1(), 480)
	assert.NotNil(t, testSolution.DataProvider.Get2DArray())
}
