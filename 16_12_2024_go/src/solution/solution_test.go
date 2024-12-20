package solution

import (
	"testing"

	"github.com/M4KIF/advent_of_code_2024/16_12_2024_go/src/data_provider"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/file_handling"
	"github.com/stretchr/testify/assert"
)

func TestPart1With15By15DataFromExample(t *testing.T) {
	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	assert.Equal(t, testSolution.Part1(), 7036)
}

func TestPart1With16By16DataFromExample2(t *testing.T) {
	path := "./test_data/data_3.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	assert.Equal(t, testSolution.Part1(), 11048)
}
