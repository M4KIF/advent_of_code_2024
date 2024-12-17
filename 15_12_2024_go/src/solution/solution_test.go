package solution

import (
	"fmt"
	"testing"

	"github.com/M4KIF/advent_of_code_2024/15_12_2024_go/src/data_provider"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/file_handling"
	"github.com/stretchr/testify/assert"
)

func TestFindRobotPositionDescExampleBeginning(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)

	testSolution := NewSolution(dp, path)
	y, x := testSolution.FindRobotPosition()
	assert.Equal(t, y, 2)
	assert.Equal(t, x, 2)
}

func TestPushAxisXRightAtY1X2inDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[1][2] = ROBOT

	// The wanted method invocation
	// Given #.@O.O.#
	fmt.Println("line expected: ", testSolution.DataProvider.GetArea()[1])
	pushResult := testSolution.PushAxisX(C_RIGHT)

	// Expected results
	assert.True(t, pushResult)
	assert.Equal(t, testSolution.DataProvider.GetArea()[1], []string{"#", ".", ".", "@", "O", "O", ".", "#"})
}

func TestPushAxisXLeftAtY4X5inDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[4][5] = ROBOT

	// The wanted method invocation
	// Given #.#.O@.#
	fmt.Println("line expected: ", testSolution.DataProvider.GetArea()[4])
	pushResult := testSolution.PushAxisX(C_LEFT)

	// Expected results
	assert.True(t, pushResult)
	assert.Equal(t, testSolution.DataProvider.GetArea()[4], []string{"#", ".", "#", "O", "@", ".", ".", "#"})
}

func TestPushAxisYDownAtY1X4inDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[1][4] = ROBOT

	/*
		#
		@
		O
		O
		O
		O
		.
		#
	*/
	fmt.Println(testSolution.DataProvider.GetArea()[0][4])
	fmt.Println(testSolution.DataProvider.GetArea()[1][4])
	fmt.Println(testSolution.DataProvider.GetArea()[2][4])
	fmt.Println(testSolution.DataProvider.GetArea()[3][4])
	fmt.Println(testSolution.DataProvider.GetArea()[4][4])
	fmt.Println(testSolution.DataProvider.GetArea()[5][4])
	fmt.Println(testSolution.DataProvider.GetArea()[6][4])
	fmt.Println(testSolution.DataProvider.GetArea()[7][4])

	// The wanted method invocation
	pushResult := testSolution.PushAxisY(C_DOWN)

	// Expected results
	assert.True(t, pushResult)
	assert.Equal(t, testSolution.DataProvider.GetArea()[0][4], "#")
	assert.Equal(t, testSolution.DataProvider.GetArea()[1][4], ".")
	assert.Equal(t, testSolution.DataProvider.GetArea()[2][4], "@")
	assert.Equal(t, testSolution.DataProvider.GetArea()[3][4], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[4][4], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[5][4], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[6][4], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[7][4], "#")
}

func TestIntegrationPart1CalcBWith480DataSet(t *testing.T) {

	path := "./test_data/data_3.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)

	testSolution := NewSolution(dp, path)
	assert.Equal(t, testSolution.Part1(), 1)

}
