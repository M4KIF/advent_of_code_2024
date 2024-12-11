package solution

import (
	"testing"

	"github.com/M4KIF/advent_of_code_2024/10_12_2024_go/src/data_provider"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/file_handling"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationSolutionWithFileProvidingCapab(t *testing.T) {

	path := "./test_data/test_data_1.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)

	testSolution := NewSolution(dp, path)

	assert.Equal(t, testSolution.Part1(), 1)
	assert.NotNil(t, testSolution.DataProvider.Get2DArray())
}

func TestIntegMethodHelper(t *testing.T) {

	path := "./test_data/test_data_1.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)

	testSolution := NewSolution(dp, path)

	assert.Equal(t, testSolution.Part1(), 1)
}

func TestPart1ExampleData(t *testing.T) {

	path := "./test_data/test_data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)

	testSolution := NewSolution(dp, path)

	assert.Equal(t, testSolution.Part1(), 36)
}

func TestPart2ExampleData(t *testing.T) {

	path := "./test_data/test_data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)

	testSolution := NewSolution(dp, path)

	assert.Equal(t, 81, testSolution.Part2())
}
