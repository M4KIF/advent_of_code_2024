package solution

import (
	"testing"

	"github.com/M4KIF/advent_of_code_2024/middleware/go/file_handling"
	"github.com/M4KIF/advent_of_code_2024/project_stubs/go/src/data_provider"
)

func TestIntegrationPart1CalcBWith480DataSet(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)

	testSolution := NewSolution(dp, path)

}
