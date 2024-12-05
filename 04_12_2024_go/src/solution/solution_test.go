package solution

import (
	"os"
	"testing"

	"github.com/M4KIF/advent_of_code_2024/04_12_2024_go/src/data_provider"
	"github.com/stretchr/testify/assert"
)

type IoMock struct {
}

func (i IoMock) Open(path string) (*os.File, error) {
	if file, e := os.Open(path); e != nil {
		return nil, e
	} else {
		return file, nil
	}
}

func TestSolutionPart1RightScan(t *testing.T) {

	solution := NewSolution(data_provider.NewDataProvider(IoMock{}), "./test_data/part_1_01.txt")
	assert.Equal(t, solution.SolvePart1(), 1)
}

func TestSolutionPart1LeftScan(t *testing.T) {

	solution := NewSolution(data_provider.NewDataProvider(IoMock{}), "./test_data/part_1_02.txt")
	assert.Equal(t, solution.SolvePart1(), 1)
}

func TestSolutionPart1DownScan(t *testing.T) {

	solution := NewSolution(data_provider.NewDataProvider(IoMock{}), "./test_data/part_1_03.txt")
	assert.Equal(t, solution.SolvePart1(), 1)
}

func TestSolutionPart1UpScan(t *testing.T) {

	solution := NewSolution(data_provider.NewDataProvider(IoMock{}), "./test_data/part_1_04.txt")
	assert.Equal(t, solution.SolvePart1(), 1)
}

func TestSolutionPart1DownRightScan(t *testing.T) {

	solution := NewSolution(data_provider.NewDataProvider(IoMock{}), "./test_data/part_1_05.txt")
	assert.Equal(t, solution.SolvePart1(), 1)
}

func TestSolutionPart1UpRightScan(t *testing.T) {

	solution := NewSolution(data_provider.NewDataProvider(IoMock{}), "./test_data/part_1_06.txt")
	assert.Equal(t, solution.SolvePart1(), 1)
}

func TestSolutionPart1DownLeftScan(t *testing.T) {

	solution := NewSolution(data_provider.NewDataProvider(IoMock{}), "./test_data/part_1_07.txt")
	assert.Equal(t, solution.SolvePart1(), 1)
}

func TestSolutionPart1UpLeftScan(t *testing.T) {

	solution := NewSolution(data_provider.NewDataProvider(IoMock{}), "./test_data/part_1_08.txt")
	assert.Equal(t, solution.SolvePart1(), 1)
}

func TestSolutionPart2UpRightScan(t *testing.T) {

	solution := NewSolution(data_provider.NewDataProvider(IoMock{}), "./test_data/part_2_01.txt")
	assert.Equal(t, solution.SolvePart2(1, 0), 1)
}

func TestSolutionPart2DownRightScan(t *testing.T) {

	solution := NewSolution(data_provider.NewDataProvider(IoMock{}), "./test_data/part_2_02.txt")
	assert.Equal(t, solution.SolvePart2(1, 0), 1)
}

func TestSolutionPart2UpLeftScan(t *testing.T) {

	solution := NewSolution(data_provider.NewDataProvider(IoMock{}), "./test_data/part_2_03.txt")
	assert.Equal(t, solution.SolvePart2(1, 0), 1)
}

func TestSolutionPart2DownLeftScan(t *testing.T) {

	solution := NewSolution(data_provider.NewDataProvider(IoMock{}), "./test_data/part_2_04.txt")
	assert.Equal(t, solution.SolvePart2(1, 0), 1)
}

func TestSolutionPart2EdgeCase1(t *testing.T) {

	solution := NewSolution(data_provider.NewDataProvider(IoMock{}), "./test_data/part_2_edge_case_1.txt")
	assert.Equal(t, solution.SolvePart2(2, 2), 0)
}

func TestSolutionPart2EdgeCase2(t *testing.T) {

	solution := NewSolution(data_provider.NewDataProvider(IoMock{}), "./test_data/part_2_edge_case_2.txt")
	assert.Equal(t, solution.SolvePart2(2, 2), 0)
}
