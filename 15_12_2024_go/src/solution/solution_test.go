package solution

import (
	"fmt"
	"strings"
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
	y, x := testSolution.FindRobotPosition(testSolution.DataProvider.GetArea())
	assert.Equal(t, y, 2)
	assert.Equal(t, x, 2)
}

/*
PART 1: TESTS
*/

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
	pushResult := testSolution.MovementXAxisStraight(C_RIGHT)

	// Expected results
	assert.True(t, pushResult)
	assert.Equal(t, testSolution.DataProvider.GetArea()[1], []string{"#", ".", ".", "@", "O", "O", ".", "#"})
}

func TestPushAxisXLeftAtY1X6inDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[1][1] = BLANK
	testSolution.DataProvider.GetArea()[1][2] = BOX
	testSolution.DataProvider.GetArea()[1][3] = BLANK
	testSolution.DataProvider.GetArea()[1][4] = BOX
	testSolution.DataProvider.GetArea()[1][5] = BLANK
	testSolution.DataProvider.GetArea()[1][6] = ROBOT

	// The wanted method invocation
	fmt.Println("line expected: ", testSolution.DataProvider.GetArea()[1])
	pushResult := testSolution.PushAxisXnew(C_LEFT)

	// Expected results
	assert.True(t, pushResult)
	assert.Equal(t, testSolution.DataProvider.GetArea()[1], []string{"#", ".", "O", ".", "O", "@", ".", "#"})
}

func TestPushAxisXLeftAtY1X5inDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[1][1] = BLANK
	testSolution.DataProvider.GetArea()[1][2] = BOX
	testSolution.DataProvider.GetArea()[1][3] = BLANK
	testSolution.DataProvider.GetArea()[1][4] = BOX
	testSolution.DataProvider.GetArea()[1][5] = ROBOT
	testSolution.DataProvider.GetArea()[1][6] = BLANK

	// The wanted method invocation
	fmt.Println("line expected: ", testSolution.DataProvider.GetArea()[1])
	pushResult := testSolution.PushAxisXnew(C_LEFT)

	// Expected results
	assert.True(t, pushResult)
	assert.Equal(t, testSolution.DataProvider.GetArea()[1], []string{"#", ".", "O", "O", "@", ".", ".", "#"})
}

func TestPushAxisXLeftAtY1X5BoToTheRightShouldNotCountinDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[1][1] = BLANK
	testSolution.DataProvider.GetArea()[1][2] = BOX
	testSolution.DataProvider.GetArea()[1][3] = BLANK
	testSolution.DataProvider.GetArea()[1][4] = BOX
	testSolution.DataProvider.GetArea()[1][5] = ROBOT
	testSolution.DataProvider.GetArea()[1][6] = BOX

	// The wanted method invocation
	fmt.Println("line expected: ", testSolution.DataProvider.GetArea()[1])
	pushResult := testSolution.PushAxisXnew(C_LEFT)

	// Expected results
	assert.True(t, pushResult)
	assert.Equal(t, testSolution.DataProvider.GetArea()[1], []string{"#", ".", "O", "O", "@", ".", "O", "#"})
}

func TestPushAxisXLeftAtY1X4TwoPushDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[1][1] = BLANK
	testSolution.DataProvider.GetArea()[1][2] = BOX
	testSolution.DataProvider.GetArea()[1][3] = BOX
	testSolution.DataProvider.GetArea()[1][4] = ROBOT
	testSolution.DataProvider.GetArea()[1][5] = BLANK
	testSolution.DataProvider.GetArea()[1][6] = BOX

	// The wanted method invocation
	fmt.Println("line expected: ", testSolution.DataProvider.GetArea()[1])
	pushResult := testSolution.PushAxisXnew(C_LEFT)

	// Expected results
	assert.True(t, pushResult)
	assert.Equal(t, testSolution.DataProvider.GetArea()[1], []string{"#", "O", "O", "@", ".", ".", "O", "#"})
}

func TestPushAxisXRightAtY1X3inDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[1][1] = BLANK
	testSolution.DataProvider.GetArea()[1][2] = BLANK
	testSolution.DataProvider.GetArea()[1][3] = ROBOT
	testSolution.DataProvider.GetArea()[1][4] = BOX
	testSolution.DataProvider.GetArea()[1][5] = BOX
	testSolution.DataProvider.GetArea()[1][6] = BLANK

	// The wanted method invocation
	// Given #.@O.O.#
	fmt.Println("line expected: ", testSolution.DataProvider.GetArea()[1])
	pushResult := testSolution.PushAxisXnew(C_RIGHT)

	// Expected results
	assert.True(t, pushResult)
	assert.Equal(t, testSolution.DataProvider.GetArea()[1], []string{"#", ".", ".", ".", "@", "O", "O", "#"})
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
	pushResult := testSolution.PushAxisXnew(C_LEFT)

	// Expected results
	assert.True(t, pushResult)
	assert.Equal(t, testSolution.DataProvider.GetArea()[4], []string{"#", ".", "#", "O", "@", ".", ".", "#"})
}

func TestPushAxisYDownAtY1X2PushDownOneinDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[1][2] = BOX
	testSolution.DataProvider.GetArea()[2][2] = ROBOT
	testSolution.DataProvider.GetArea()[3][2] = BLANK
	testSolution.DataProvider.GetArea()[4][2] = BOX
	testSolution.DataProvider.GetArea()[5][2] = BOX
	testSolution.DataProvider.GetArea()[6][2] = BLANK

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
	fmt.Println(testSolution.DataProvider.GetArea()[0][2])
	fmt.Println(testSolution.DataProvider.GetArea()[1][2])
	fmt.Println(testSolution.DataProvider.GetArea()[2][2])
	fmt.Println(testSolution.DataProvider.GetArea()[3][2])
	fmt.Println(testSolution.DataProvider.GetArea()[4][2])
	fmt.Println(testSolution.DataProvider.GetArea()[5][2])
	fmt.Println(testSolution.DataProvider.GetArea()[6][2])
	fmt.Println(testSolution.DataProvider.GetArea()[7][2])

	// The wanted method invocation
	pushResult := testSolution.PushAxisYnew(C_DOWN)

	// Expected results
	assert.True(t, pushResult)
	assert.Equal(t, testSolution.DataProvider.GetArea()[0][2], "#")
	assert.Equal(t, testSolution.DataProvider.GetArea()[1][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[2][2], ".")
	assert.Equal(t, testSolution.DataProvider.GetArea()[3][2], "@")
	assert.Equal(t, testSolution.DataProvider.GetArea()[4][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[5][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[6][2], ".")
	assert.Equal(t, testSolution.DataProvider.GetArea()[7][2], "#")
}

func TestPushAxisYDownAtY2X2PushTwoBoxesinDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[1][2] = BOX
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[3][2] = ROBOT
	testSolution.DataProvider.GetArea()[4][2] = BOX
	testSolution.DataProvider.GetArea()[5][2] = BOX
	testSolution.DataProvider.GetArea()[6][2] = BLANK

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
	fmt.Println(testSolution.DataProvider.GetArea()[0][2])
	fmt.Println(testSolution.DataProvider.GetArea()[1][2])
	fmt.Println(testSolution.DataProvider.GetArea()[2][2])
	fmt.Println(testSolution.DataProvider.GetArea()[3][2])
	fmt.Println(testSolution.DataProvider.GetArea()[4][2])
	fmt.Println(testSolution.DataProvider.GetArea()[5][2])
	fmt.Println(testSolution.DataProvider.GetArea()[6][2])
	fmt.Println(testSolution.DataProvider.GetArea()[7][2])

	// The wanted method invocation
	pushResult := testSolution.PushAxisYnew(C_DOWN)

	fmt.Println(testSolution.DataProvider.GetArea()[0][2])
	fmt.Println(testSolution.DataProvider.GetArea()[1][2])
	fmt.Println(testSolution.DataProvider.GetArea()[2][2])
	fmt.Println(testSolution.DataProvider.GetArea()[3][2])
	fmt.Println(testSolution.DataProvider.GetArea()[4][2])
	fmt.Println(testSolution.DataProvider.GetArea()[5][2])
	fmt.Println(testSolution.DataProvider.GetArea()[6][2])
	fmt.Println(testSolution.DataProvider.GetArea()[7][2])

	// Expected results
	assert.True(t, pushResult)
	assert.Equal(t, testSolution.DataProvider.GetArea()[0][2], "#")
	assert.Equal(t, testSolution.DataProvider.GetArea()[1][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[2][2], ".")
	assert.Equal(t, testSolution.DataProvider.GetArea()[3][2], ".")
	assert.Equal(t, testSolution.DataProvider.GetArea()[4][2], "@")
	assert.Equal(t, testSolution.DataProvider.GetArea()[5][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[6][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[7][2], "#")
}

func TestPushAxisYUPAtY3X2PushTwoBoxesinDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[1][2] = BOX
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[3][2] = ROBOT
	testSolution.DataProvider.GetArea()[4][2] = BOX
	testSolution.DataProvider.GetArea()[5][2] = BOX
	testSolution.DataProvider.GetArea()[6][2] = BLANK

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
	fmt.Println(testSolution.DataProvider.GetArea()[0][2])
	fmt.Println(testSolution.DataProvider.GetArea()[1][2])
	fmt.Println(testSolution.DataProvider.GetArea()[2][2])
	fmt.Println(testSolution.DataProvider.GetArea()[3][2])
	fmt.Println(testSolution.DataProvider.GetArea()[4][2])
	fmt.Println(testSolution.DataProvider.GetArea()[5][2])
	fmt.Println(testSolution.DataProvider.GetArea()[6][2])
	fmt.Println(testSolution.DataProvider.GetArea()[7][2])

	// The wanted method invocation
	pushResult := testSolution.PushAxisYnew(C_TOP)

	fmt.Println(testSolution.DataProvider.GetArea()[0][2])
	fmt.Println(testSolution.DataProvider.GetArea()[1][2])
	fmt.Println(testSolution.DataProvider.GetArea()[2][2])
	fmt.Println(testSolution.DataProvider.GetArea()[3][2])
	fmt.Println(testSolution.DataProvider.GetArea()[4][2])
	fmt.Println(testSolution.DataProvider.GetArea()[5][2])
	fmt.Println(testSolution.DataProvider.GetArea()[6][2])
	fmt.Println(testSolution.DataProvider.GetArea()[7][2])

	// Expected results
	assert.True(t, pushResult)
	assert.Equal(t, testSolution.DataProvider.GetArea()[0][2], "#")
	assert.Equal(t, testSolution.DataProvider.GetArea()[1][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[2][2], "@")
	assert.Equal(t, testSolution.DataProvider.GetArea()[3][2], ".")
	assert.Equal(t, testSolution.DataProvider.GetArea()[4][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[5][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[6][2], ".")
	assert.Equal(t, testSolution.DataProvider.GetArea()[7][2], "#")
}

func TestPushAxisYUPAtY6X2PushTwoBoxesinDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[1][2] = BLANK
	testSolution.DataProvider.GetArea()[2][2] = BOX
	testSolution.DataProvider.GetArea()[3][2] = BOX
	testSolution.DataProvider.GetArea()[4][2] = BOX
	testSolution.DataProvider.GetArea()[5][2] = BOX
	testSolution.DataProvider.GetArea()[6][2] = ROBOT

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
	fmt.Println(testSolution.DataProvider.GetArea()[0][2])
	fmt.Println(testSolution.DataProvider.GetArea()[1][2])
	fmt.Println(testSolution.DataProvider.GetArea()[2][2])
	fmt.Println(testSolution.DataProvider.GetArea()[3][2])
	fmt.Println(testSolution.DataProvider.GetArea()[4][2])
	fmt.Println(testSolution.DataProvider.GetArea()[5][2])
	fmt.Println(testSolution.DataProvider.GetArea()[6][2])
	fmt.Println(testSolution.DataProvider.GetArea()[7][2])

	// The wanted method invocation
	pushResult := testSolution.PushAxisYnew(C_TOP)

	fmt.Println(testSolution.DataProvider.GetArea()[0][2])
	fmt.Println(testSolution.DataProvider.GetArea()[1][2])
	fmt.Println(testSolution.DataProvider.GetArea()[2][2])
	fmt.Println(testSolution.DataProvider.GetArea()[3][2])
	fmt.Println(testSolution.DataProvider.GetArea()[4][2])
	fmt.Println(testSolution.DataProvider.GetArea()[5][2])
	fmt.Println(testSolution.DataProvider.GetArea()[6][2])
	fmt.Println(testSolution.DataProvider.GetArea()[7][2])

	// Expected results
	assert.True(t, pushResult)
	assert.Equal(t, testSolution.DataProvider.GetArea()[0][2], "#")
	assert.Equal(t, testSolution.DataProvider.GetArea()[1][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[2][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[3][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[4][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[5][2], "@")
	assert.Equal(t, testSolution.DataProvider.GetArea()[6][2], ".")
	assert.Equal(t, testSolution.DataProvider.GetArea()[7][2], "#")
}

func TestPushAxisYUPAtY6X2PushOneBoxinDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[1][2] = BOX
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[3][2] = BLANK
	testSolution.DataProvider.GetArea()[4][2] = BLANK
	testSolution.DataProvider.GetArea()[5][2] = BOX
	testSolution.DataProvider.GetArea()[6][2] = ROBOT

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
	fmt.Println(testSolution.DataProvider.GetArea()[0][2])
	fmt.Println(testSolution.DataProvider.GetArea()[1][2])
	fmt.Println(testSolution.DataProvider.GetArea()[2][2])
	fmt.Println(testSolution.DataProvider.GetArea()[3][2])
	fmt.Println(testSolution.DataProvider.GetArea()[4][2])
	fmt.Println(testSolution.DataProvider.GetArea()[5][2])
	fmt.Println(testSolution.DataProvider.GetArea()[6][2])
	fmt.Println(testSolution.DataProvider.GetArea()[7][2])

	// The wanted method invocation
	pushResult := testSolution.PushAxisYnew(C_TOP)

	fmt.Println(testSolution.DataProvider.GetArea()[0][2])
	fmt.Println(testSolution.DataProvider.GetArea()[1][2])
	fmt.Println(testSolution.DataProvider.GetArea()[2][2])
	fmt.Println(testSolution.DataProvider.GetArea()[3][2])
	fmt.Println(testSolution.DataProvider.GetArea()[4][2])
	fmt.Println(testSolution.DataProvider.GetArea()[5][2])
	fmt.Println(testSolution.DataProvider.GetArea()[6][2])
	fmt.Println(testSolution.DataProvider.GetArea()[7][2])

	// Expected results
	assert.True(t, pushResult)
	assert.Equal(t, testSolution.DataProvider.GetArea()[0][2], "#")
	assert.Equal(t, testSolution.DataProvider.GetArea()[1][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[2][2], ".")
	assert.Equal(t, testSolution.DataProvider.GetArea()[3][2], ".")
	assert.Equal(t, testSolution.DataProvider.GetArea()[4][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[5][2], "@")
	assert.Equal(t, testSolution.DataProvider.GetArea()[6][2], ".")
	assert.Equal(t, testSolution.DataProvider.GetArea()[7][2], "#")
}

func TestPushAxisYUPAtY5X2PushOneBoxinDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[1][2] = BOX
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[3][2] = BLANK
	testSolution.DataProvider.GetArea()[4][2] = BOX
	testSolution.DataProvider.GetArea()[5][2] = ROBOT
	testSolution.DataProvider.GetArea()[6][2] = BLANK

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
	fmt.Println(testSolution.DataProvider.GetArea()[0][2])
	fmt.Println(testSolution.DataProvider.GetArea()[1][2])
	fmt.Println(testSolution.DataProvider.GetArea()[2][2])
	fmt.Println(testSolution.DataProvider.GetArea()[3][2])
	fmt.Println(testSolution.DataProvider.GetArea()[4][2])
	fmt.Println(testSolution.DataProvider.GetArea()[5][2])
	fmt.Println(testSolution.DataProvider.GetArea()[6][2])
	fmt.Println(testSolution.DataProvider.GetArea()[7][2])

	// The wanted method invocation
	pushResult := testSolution.PushAxisYnew(C_TOP)

	fmt.Println(testSolution.DataProvider.GetArea()[0][2])
	fmt.Println(testSolution.DataProvider.GetArea()[1][2])
	fmt.Println(testSolution.DataProvider.GetArea()[2][2])
	fmt.Println(testSolution.DataProvider.GetArea()[3][2])
	fmt.Println(testSolution.DataProvider.GetArea()[4][2])
	fmt.Println(testSolution.DataProvider.GetArea()[5][2])
	fmt.Println(testSolution.DataProvider.GetArea()[6][2])
	fmt.Println(testSolution.DataProvider.GetArea()[7][2])

	// Expected results
	assert.True(t, pushResult)
	assert.Equal(t, testSolution.DataProvider.GetArea()[0][2], "#")
	assert.Equal(t, testSolution.DataProvider.GetArea()[1][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[2][2], ".")
	assert.Equal(t, testSolution.DataProvider.GetArea()[3][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[4][2], "@")
	assert.Equal(t, testSolution.DataProvider.GetArea()[5][2], ".")
	assert.Equal(t, testSolution.DataProvider.GetArea()[6][2], ".")
	assert.Equal(t, testSolution.DataProvider.GetArea()[7][2], "#")
}

func TestPushAxisYUPAtY4X2WallBoxinDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[1][2] = BOX
	testSolution.DataProvider.GetArea()[2][2] = WALL
	testSolution.DataProvider.GetArea()[3][2] = BOX
	testSolution.DataProvider.GetArea()[4][2] = ROBOT
	testSolution.DataProvider.GetArea()[5][2] = BLANK
	testSolution.DataProvider.GetArea()[6][2] = BLANK

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
	fmt.Println(testSolution.DataProvider.GetArea()[0][2])
	fmt.Println(testSolution.DataProvider.GetArea()[1][2])
	fmt.Println(testSolution.DataProvider.GetArea()[2][2])
	fmt.Println(testSolution.DataProvider.GetArea()[3][2])
	fmt.Println(testSolution.DataProvider.GetArea()[4][2])
	fmt.Println(testSolution.DataProvider.GetArea()[5][2])
	fmt.Println(testSolution.DataProvider.GetArea()[6][2])
	fmt.Println(testSolution.DataProvider.GetArea()[7][2])

	// The wanted method invocation
	pushResult := testSolution.PushAxisYnew(C_TOP)

	fmt.Println(testSolution.DataProvider.GetArea()[0][2])
	fmt.Println(testSolution.DataProvider.GetArea()[1][2])
	fmt.Println(testSolution.DataProvider.GetArea()[2][2])
	fmt.Println(testSolution.DataProvider.GetArea()[3][2])
	fmt.Println(testSolution.DataProvider.GetArea()[4][2])
	fmt.Println(testSolution.DataProvider.GetArea()[5][2])
	fmt.Println(testSolution.DataProvider.GetArea()[6][2])
	fmt.Println(testSolution.DataProvider.GetArea()[7][2])

	// Expected results
	assert.False(t, pushResult)
	assert.Equal(t, testSolution.DataProvider.GetArea()[0][2], "#")
	assert.Equal(t, testSolution.DataProvider.GetArea()[1][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[2][2], "#")
	assert.Equal(t, testSolution.DataProvider.GetArea()[3][2], "O")
	assert.Equal(t, testSolution.DataProvider.GetArea()[4][2], "@")
	assert.Equal(t, testSolution.DataProvider.GetArea()[5][2], ".")
	assert.Equal(t, testSolution.DataProvider.GetArea()[6][2], ".")
	assert.Equal(t, testSolution.DataProvider.GetArea()[7][2], "#")
}

func TestDownNotCloningOnExampleData(t *testing.T) {

	path := "./test_data/data_3.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)

	testSolution := NewSolution(dp, path)

	// Counting the number of boxes to check for any big errors

	pre_test := 0
	for _, line := range testSolution.DataProvider.GetArea() {
		for _, c := range line {
			if c == "O" {
				pre_test++
			}
		}
	}

	testSolution.Part1()

	post_test := 0
	for _, line := range testSolution.DataProvider.GetArea() {
		for _, c := range line {
			if c == "O" {
				post_test++
			}
		}
	}

	assert.Equal(t, pre_test, post_test)

}

func TestIntegrationPart1CalcBWith480DataSet(t *testing.T) {

	path := "./test_data/data_4.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)

	testSolution := NewSolution(dp, path)

	// Counting the number of boxes to check for any big errors

	pre_test := 0
	for _, line := range testSolution.DataProvider.GetArea() {
		for _, c := range line {
			if c == "O" {
				pre_test++
			}
		}
	}

	test_result := testSolution.Part1()

	post_test := 0
	for _, line := range testSolution.DataProvider.GetArea() {
		for _, c := range line {
			if c == "O" {
				post_test++
			}
		}
	}

	assert.Equal(t, pre_test, post_test)
	assert.Equal(t, test_result, 10092)

}

/*
PART 2: TESTS
*/

func TestConvertPart1AreaToPart2Version(t *testing.T) {

	path := "./test_data/data_5.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)

	testSolution := NewSolution(dp, path)

	pre_test := len(testSolution.DataProvider.GetArea()[0])

	part2_area := testSolution.ConvertPart1AreaToPart2Version()

	post_test := len(part2_area[0])

	assert.NotEqual(t, pre_test, post_test)

	assert.Equal(t, post_test, 14)

	post_test_first_row := strings.Join(part2_area[0], "")
	assert.Equal(t, post_test_first_row, "##############")

	post_test_third_row := strings.Join(part2_area[3], "")
	assert.Equal(t, post_test_third_row, "##....[][]@.##")
}

func TestPushAxisXRightPart2AtY1X4inDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[1][2] = ROBOT

	// Given # # . . @ . [ ] . . [ ] . . # #
	testArea := testSolution.ConvertPart1AreaToPart2Version()
	// Wanted # # . . . @ [ ] . . [ ] . . # #
	fmt.Println("line pretest: ", testArea[1])

	testArea = testSolution.PushAxisXpart2(testArea, C_RIGHT)
	fmt.Println("line posttest: ", testArea[1])

	// Expected results
	assert.Equal(t, testArea[1], []string{"#", "#", ".", ".", ".", "@", "[", "]", ".", ".", "[", "]", ".", ".", "#", "#"})
}

func TestPushAxisXRightPart2AtY1X5inDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[1][2] = ROBOT

	// Given # # . . . @ [ ] . . [ ] . . # #
	testArea := testSolution.ConvertPart1AreaToPart2Version()
	testArea[1][4] = BLANK
	testArea[1][5] = ROBOT

	// Wanted # # . . . . @ [ ] . [ ] . . # #
	fmt.Println("line pretest: ", testArea[1])

	testArea = testSolution.PushAxisXpart2(testArea, C_RIGHT)
	fmt.Println("line posttest: ", testArea[1])

	// Expected results
	assert.Equal(t, testArea[1], []string{"#", "#", ".", ".", ".", ".", "@", "[", "]", ".", "[", "]", ".", ".", "#", "#"})
}

func TestPushAxisXRightPart2AtY1X6inDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[1][2] = ROBOT

	// Given # # . . . . @ [ ] . [ ] . . # #
	testArea := testSolution.ConvertPart1AreaToPart2Version()
	testArea[1][4] = BLANK
	testArea[1][5] = BLANK
	testArea[1][6] = ROBOT
	testArea[1][7] = "["
	testArea[1][8] = "]"

	// Wanted # # . . . . @ [ ] . [ ] . . # #
	fmt.Println("line pretest: ", testArea[1])

	testArea = testSolution.PushAxisXpart2(testArea, C_RIGHT)
	fmt.Println("line posttest: ", testArea[1])

	// Expected results
	assert.Equal(t, testArea[1], []string{"#", "#", ".", ".", ".", ".", ".", "@", "[", "]", "[", "]", ".", ".", "#", "#"})
}

func TestPushAxisXRightPart2AtY1X7inDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[1][2] = ROBOT

	// Given # # . . . . . @ [ ] [ ] . . # #
	testArea := testSolution.ConvertPart1AreaToPart2Version()
	testArea[1][4] = BLANK
	testArea[1][5] = BLANK
	testArea[1][6] = BLANK
	testArea[1][7] = ROBOT
	testArea[1][8] = "["
	testArea[1][9] = "]"

	// Wanted # # . . . . @ [ ] . [ ] . . # #
	fmt.Println("line pretest: ", testArea[1])

	testArea = testSolution.PushAxisXpart2(testArea, C_RIGHT)
	fmt.Println("line posttest: ", testArea[1])

	// Expected results
	assert.Equal(t, testArea[1], []string{"#", "#", ".", ".", ".", ".", ".", ".", "@", "[", "]", "[", "]", ".", "#", "#"})
}

func TestPushAxisXRightPart2AtY1X8inDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[1][2] = ROBOT

	// Given # # . . . . . @ [ ] [ ] . . # #
	testArea := testSolution.ConvertPart1AreaToPart2Version()
	testArea[1][4] = BLANK
	testArea[1][5] = BLANK
	testArea[1][6] = BLANK
	testArea[1][7] = BLANK
	testArea[1][8] = ROBOT
	testArea[1][9] = "["
	testArea[1][10] = "]"
	testArea[1][11] = "["
	testArea[1][12] = "]"

	// Wanted # # . . . . @ [ ] . [ ] . . # #
	fmt.Println("line pretest: ", testArea[1])

	testArea = testSolution.PushAxisXpart2(testArea, C_RIGHT)
	fmt.Println("line posttest: ", testArea[1])

	// Expected results
	assert.Equal(t, testArea[1], []string{"#", "#", ".", ".", ".", ".", ".", ".", ".", "@", "[", "]", "[", "]", "#", "#"})
}

func TestPushAxisXLeftPart2AtY1X13inDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[1][2] = ROBOT

	// Given # # . . . . . @ [ ] [ ] . . # #
	testArea := testSolution.ConvertPart1AreaToPart2Version()
	testArea[1][2] = BLANK
	testArea[1][3] = BLANK
	testArea[1][4] = BLANK
	testArea[1][5] = BLANK
	testArea[1][6] = "["
	testArea[1][7] = "]"
	testArea[1][8] = "["
	testArea[1][9] = "]"
	testArea[1][10] = "."
	testArea[1][11] = "."
	testArea[1][12] = "."
	testArea[1][13] = "@"

	// Wanted # # . . . . @ [ ] . [ ] . . # #
	fmt.Println("line pretest: ", testArea[1])

	testArea = testSolution.PushAxisXpart2(testArea, C_LEFT)
	fmt.Println("line posttest: ", testArea[1])

	// Expected results
	assert.Equal(t, testArea[1], []string{"#", "#", ".", ".", ".", ".", "[", "]", "[", "]", ".", ".", "@", ".", "#", "#"})
}

func TestPushAxisXLeftPart2AtY1X12inDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[1][2] = ROBOT

	// Given # # . . . . . @ [ ] [ ] . . # #
	testArea := testSolution.ConvertPart1AreaToPart2Version()
	testArea[1][2] = BLANK
	testArea[1][3] = BLANK
	testArea[1][4] = BLANK
	testArea[1][5] = BLANK
	testArea[1][6] = "["
	testArea[1][7] = "]"
	testArea[1][8] = "["
	testArea[1][9] = "]"
	testArea[1][10] = "."
	testArea[1][11] = "."
	testArea[1][12] = "@"
	testArea[1][13] = "."

	// Wanted # # . . . . @ [ ] . [ ] . . # #
	fmt.Println("line pretest: ", testArea[1])

	testArea = testSolution.PushAxisXpart2(testArea, C_LEFT)
	fmt.Println("line posttest: ", testArea[1])

	// Expected results
	assert.Equal(t, testArea[1], []string{"#", "#", ".", ".", ".", ".", "[", "]", "[", "]", ".", "@", ".", ".", "#", "#"})
}

func TestPushAxisXLeftPart2AtY1X11inDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[1][2] = ROBOT

	// Given # # . . . . . @ [ ] [ ] . . # #
	testArea := testSolution.ConvertPart1AreaToPart2Version()
	testArea[1][2] = BLANK
	testArea[1][3] = BLANK
	testArea[1][4] = BLANK
	testArea[1][5] = BLANK
	testArea[1][6] = "["
	testArea[1][7] = "]"
	testArea[1][8] = "["
	testArea[1][9] = "]"
	testArea[1][10] = "."
	testArea[1][11] = "@"
	testArea[1][12] = "."
	testArea[1][13] = "."

	// Wanted # # . . . . @ [ ] . [ ] . . # #
	fmt.Println("line pretest: ", testArea[1])

	testArea = testSolution.PushAxisXpart2(testArea, C_LEFT)
	fmt.Println("line posttest: ", testArea[1])

	// Expected results
	assert.Equal(t, testArea[1], []string{"#", "#", ".", ".", ".", ".", "[", "]", "[", "]", "@", ".", ".", ".", "#", "#"})
}

func TestPushAxisXLeftPart2AtY1X10Push2BoxesinDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[1][2] = ROBOT

	// Given # # . . . . . @ [ ] [ ] . . # #
	testArea := testSolution.ConvertPart1AreaToPart2Version()
	testArea[1][2] = BLANK
	testArea[1][3] = BLANK
	testArea[1][4] = BLANK
	testArea[1][5] = BLANK
	testArea[1][6] = "["
	testArea[1][7] = "]"
	testArea[1][8] = "["
	testArea[1][9] = "]"
	testArea[1][10] = "@"
	testArea[1][11] = "."
	testArea[1][12] = "."
	testArea[1][13] = "."

	// Wanted # # . . . . @ [ ] . [ ] . . # #
	fmt.Println("line pretest: ", testArea[1])

	testArea = testSolution.PushAxisXpart2(testArea, C_LEFT)
	fmt.Println("line posttest: ", testArea[1])

	// Expected results
	assert.Equal(t, testArea[1], []string{"#", "#", ".", ".", ".", "[", "]", "[", "]", "@", ".", ".", ".", ".", "#", "#"})
}

func TestPushAxisXLeftPart2AtY1X10Push3BoxesinDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK
	testSolution.DataProvider.GetArea()[1][2] = ROBOT

	// Given # # . . . . . @ [ ] [ ] . . # #
	testArea := testSolution.ConvertPart1AreaToPart2Version()
	testArea[1][2] = BLANK
	testArea[1][3] = BLANK
	testArea[1][4] = "["
	testArea[1][5] = "]"
	testArea[1][6] = "["
	testArea[1][7] = "]"
	testArea[1][8] = "["
	testArea[1][9] = "]"
	testArea[1][10] = "@"
	testArea[1][11] = "."
	testArea[1][12] = "."
	testArea[1][13] = "."

	// Wanted # # . . . . @ [ ] . [ ] . . # #
	fmt.Println("line pretest: ", testArea[1])

	testArea = testSolution.PushAxisXpart2(testArea, C_LEFT)
	fmt.Println("line posttest: ", testArea[1])

	// Expected results
	assert.Equal(t, testArea[1], []string{"#", "#", ".", "[", "]", "[", "]", "[", "]", "@", ".", ".", ".", ".", "#", "#"})
}

func TestPushAxisXLeftPart2AtY1X8WallNoMovementinDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK

	// Given # # . . . . . @ [ ] [ ] . . # #
	testArea := testSolution.ConvertPart1AreaToPart2Version()
	testArea[1][2] = "["
	testArea[1][3] = "]"
	testArea[1][4] = "["
	testArea[1][5] = "]"
	testArea[1][6] = "["
	testArea[1][7] = "]"
	testArea[1][8] = "@"
	testArea[1][9] = "."
	testArea[1][10] = "."
	testArea[1][11] = "."
	testArea[1][12] = BLANK
	testArea[1][13] = BLANK

	// Wanted # # . . . . @ [ ] . [ ] . . # #
	fmt.Println("line pretest: ", testArea[1])

	testArea = testSolution.PushAxisXpart2(testArea, C_LEFT)
	fmt.Println("line posttest: ", testArea[1])

	// Expected results
	assert.Equal(t, testArea[1], []string{"#", "#", "[", "]", "[", "]", "[", "]", "@", ".", ".", ".", ".", ".", "#", "#"})
}

func TestPushAxisYpart2DownAtY3X2PushDownOneinDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK

	// Given # # . . . . . @ [ ] [ ] . . # #
	testArea := testSolution.ConvertPart1AreaToPart2Version()

	// Altering the data to suit the test
	testArea[1][2] = BLANK
	testArea[2][2] = BLANK
	testArea[3][2] = ROBOT
	testArea[4][2] = BOX_LEFT
	testArea[4][3] = BOX_RIGHT
	testArea[5][2] = BLANK
	testArea[6][2] = BLANK

	/*
		#
		.
		.
		@
		[ ]
		.
		.
		#
	*/
	for _, line := range testArea {
		fmt.Println(line)
	}

	// The wanted method invocation
	testArea = testSolution.PushAxisYpart2(testArea, C_DOWN)

	for _, line := range testArea {
		fmt.Println(line)
	}

	// Expected results
	assert.Equal(t, testArea[0][2], "#")
	assert.Equal(t, testArea[1][2], ".")
	assert.Equal(t, testArea[2][2], ".")
	assert.Equal(t, testArea[3][2], ".")
	assert.Equal(t, testArea[4][2], "@")
	assert.Equal(t, testArea[5][2], "[")
	assert.Equal(t, testArea[5][3], "]")
	assert.Equal(t, testArea[6][2], ".")
	assert.Equal(t, testArea[7][2], "#")
}

func TestPushAxisYpart2DownAtY3X2PushDownTripleHalfStackWallNoMovementInDescExample(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK

	// Given # # . . . . . @ [ ] [ ] . . # #
	testArea := testSolution.ConvertPart1AreaToPart2Version()

	// Altering the data to suit the test
	testArea[1][2] = BLANK
	testArea[2][2] = BLANK
	testArea[3][2] = ROBOT
	testArea[4][2] = BOX_LEFT
	testArea[4][3] = BOX_RIGHT
	testArea[5][1] = BOX_LEFT
	testArea[5][2] = BOX_RIGHT
	testArea[5][3] = BOX_LEFT
	testArea[5][4] = BOX_RIGHT
	testArea[6][2] = BLANK

	/*
			#
			.
			.
			@
			[ ]
		  [ ] [ ]
			.
			#
	*/
	for _, line := range testArea {
		fmt.Println(line)
	}

	// The wanted method invocation
	testArea = testSolution.PushAxisYpart2(testArea, C_DOWN)

	for _, line := range testArea {
		fmt.Println(line)
	}

	// Expected results
	assert.Equal(t, testArea[0][2], "#")
	assert.Equal(t, testArea[1][2], ".")
	assert.Equal(t, testArea[2][2], ".")
	assert.Equal(t, testArea[3][2], "@")
	assert.Equal(t, testArea[4][2], "[")
	assert.Equal(t, testArea[5][2], "]")
	assert.Equal(t, testArea[5][3], "[")
	assert.Equal(t, testArea[6][2], ".")
	assert.Equal(t, testArea[7][2], "#")
}

func TestPushAxisYpart2DownAtY3X2PushDownTripleHalfStackInDescExample(t *testing.T) {

	path := "./test_data/data_6.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK

	// Given # # . . . . . @ [ ] [ ] . . # #
	testArea := testSolution.ConvertPart1AreaToPart2Version()

	// Altering the data to suit the test
	testArea[1][3] = BLANK
	testArea[2][3] = BLANK
	testArea[3][3] = ROBOT
	testArea[4][3] = BOX_LEFT
	testArea[4][4] = BOX_RIGHT
	testArea[5][2] = BOX_LEFT
	testArea[5][3] = BOX_RIGHT
	testArea[5][4] = BOX_LEFT
	testArea[5][5] = BOX_RIGHT
	testArea[6][3] = BLANK

	/*
			#
			.
			.
			@
			[ ]
		  [ ] [ ]
			.
			#
	*/
	for _, line := range testArea {
		fmt.Println(line)
	}

	// The wanted method invocation
	testArea = testSolution.PushAxisYpart2(testArea, C_DOWN)

	for _, line := range testArea {
		fmt.Println(line)
	}

	// Expected results
	assert.Equal(t, testArea[0][3], "#")
	assert.Equal(t, testArea[1][3], ".")
	assert.Equal(t, testArea[2][3], ".")
	assert.Equal(t, testArea[3][3], ".")
	assert.Equal(t, testArea[4][3], "@")
	assert.Equal(t, testArea[5][3], "[")
	assert.Equal(t, testArea[5][4], "]")
	assert.Equal(t, testArea[6][2], "[")
	assert.Equal(t, testArea[6][3], "]")
	assert.Equal(t, testArea[6][4], "[")
	assert.Equal(t, testArea[6][5], "]")
	assert.Equal(t, testArea[7][3], "#")
}

func TestPushAxisYpart2UpAtY3X2PushDownTripleHalfStackInDescExample(t *testing.T) {

	path := "./test_data/data_5.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK

	// Given # # . . . . . @ [ ] [ ] . . # #
	testArea := testSolution.ConvertPart1AreaToPart2Version()

	// Blanking the space
	testArea[3][10] = BLANK
	testArea[3][9] = BLANK
	testArea[3][8] = BLANK
	testArea[3][7] = BLANK
	testArea[3][6] = BLANK
	testArea[3][5] = BLANK
	testArea[3][4] = BLANK
	testArea[4][10] = BLANK
	testArea[4][9] = BLANK
	testArea[4][8] = BLANK
	testArea[4][7] = BLANK
	testArea[4][6] = BLANK
	testArea[4][5] = BLANK
	testArea[4][4] = BLANK

	// Altering the data to suit the test
	testArea[3][5] = ROBOT
	testArea[2][5] = BOX_LEFT
	testArea[2][6] = BOX_RIGHT
	testArea[3][6] = BOX_LEFT
	testArea[3][7] = BOX_RIGHT
	testArea[2][7] = BOX_LEFT
	testArea[2][8] = BOX_RIGHT

	/*
			#
			.
			.
			@
			[ ]
		  [ ] [ ]
			.
			#
	*/

	for _, line := range testArea {
		fmt.Println(line)
	}

	// The wanted method invocation
	testArea = testSolution.PushAxisYpart2(testArea, C_TOP)

	for _, line := range testArea {
		fmt.Println(line)
	}

	// Expected results
	assert.Equal(t, testArea[2][5], "@")
	assert.Equal(t, testArea[1][5], "[")
	assert.Equal(t, testArea[1][6], "]")
	assert.Equal(t, testArea[2][7], "[")
	assert.Equal(t, testArea[2][8], "]")
}

func TestPushAxisYpart2DownAtY3X2PushDownTripleHalfStackWithMultipleconnectedInDescExample(t *testing.T) {

	path := "./test_data/data_6.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK

	// Given # # . . . . . @ [ ] [ ] . . # #
	testArea := testSolution.ConvertPart1AreaToPart2Version()

	// Altering the data to suit the test
	testArea[1][3] = BLANK
	testArea[2][3] = BLANK
	testArea[3][3] = ROBOT
	testArea[4][3] = BOX_LEFT
	testArea[4][4] = BOX_RIGHT
	testArea[5][2] = BOX_LEFT
	testArea[5][3] = BOX_RIGHT
	testArea[5][4] = BOX_LEFT
	testArea[5][5] = BOX_RIGHT
	testArea[5][6] = BOX_LEFT
	testArea[5][7] = BOX_RIGHT
	testArea[6][6] = BOX_LEFT
	testArea[6][7] = BOX_RIGHT
	testArea[6][3] = BLANK

	/*
			#
			.
			.
			@
			[ ]
		  [ ] [ ]
			.
			#
	*/
	for _, line := range testArea {
		fmt.Println(line)
	}

	// The wanted method invocation
	testArea = testSolution.PushAxisYpart2(testArea, C_DOWN)

	for _, line := range testArea {
		fmt.Println(line)
	}

	// Expected results
	assert.Equal(t, testArea[0][3], "#")
	assert.Equal(t, testArea[1][3], ".")
	assert.Equal(t, testArea[2][3], ".")
	assert.Equal(t, testArea[3][3], ".")
	assert.Equal(t, testArea[4][3], "@")
	assert.Equal(t, testArea[5][3], "[")
	assert.Equal(t, testArea[5][4], "]")
	assert.Equal(t, testArea[5][6], "[")
	assert.Equal(t, testArea[5][7], "]")
	assert.Equal(t, testArea[6][6], "[")
	assert.Equal(t, testArea[6][7], "]")
	assert.Equal(t, testArea[6][2], "[")
	assert.Equal(t, testArea[6][3], "]")
	assert.Equal(t, testArea[6][4], "[")
	assert.Equal(t, testArea[6][5], "]")
	assert.Equal(t, testArea[7][3], "#")
}

/*
	This edge case is killing It atm

# # # # # # # # # # # # # # # # # # # #]
[# # . . . . [ ] . . . . [ ] . . [ ] # #]
[# # . . . . . . . . . . . . [ ] . . # #]
[# # . . [ ] [ ] . . . . [ ] . . [ ] # #]
[# # [ ] . [ ] . . . [ ] . . [ ] . . # #]
[# # [ ] # # . . . . [ ] . . . . . . # #]
[# # . @ [ ] . . . . . . . . [ ] . . # #]
[# # . . . . . [ ] . . [ ] . [ ] [ ] # #]
[# # . . . . . . . . [ ] . . . . . . # #]
[# # # # # # # # # # # # # # # # # # # #]
*/
func TestPushAxisYpart2UpAtY6X3TwoSimpleStackedWithEdgeCaseInDescExample(t *testing.T) {

	path := "./test_data/data_6.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK

	// Given # # . . . . . @ [ ] [ ] . . # #
	testArea := testSolution.ConvertPart1AreaToPart2Version()

	// Altering the data to suit the test
	testArea[1][2] = BLANK
	testArea[1][3] = BLANK
	testArea[2][2] = BLANK
	testArea[2][3] = BLANK
	testArea[3][2] = BLANK
	testArea[3][3] = BLANK
	testArea[3][4] = BOX_LEFT
	testArea[3][5] = BOX_RIGHT
	testArea[4][2] = BOX_LEFT
	testArea[4][3] = BOX_RIGHT
	testArea[4][5] = BOX_LEFT
	testArea[4][6] = BOX_RIGHT
	testArea[5][2] = BOX_LEFT
	testArea[5][3] = BOX_RIGHT
	testArea[6][2] = BLANK
	testArea[6][3] = ROBOT
	testArea[7][2] = BLANK
	testArea[7][3] = BLANK
	testArea[7][2] = BLANK
	testArea[7][3] = BLANK

	for _, line := range testArea {
		fmt.Println(line)
	}

	// The wanted method invocation
	testArea = testSolution.PushAxisYpart2(testArea, C_TOP)

	for _, line := range testArea {
		fmt.Println(line)
	}

	// Expected results
	assert.Equal(t, testArea[0][2], "#")
	assert.Equal(t, testArea[0][3], "#")
	assert.Equal(t, testArea[1][2], ".")
	assert.Equal(t, testArea[1][3], ".")
	assert.Equal(t, testArea[2][2], ".")
	assert.Equal(t, testArea[2][3], ".")
	assert.Equal(t, testArea[3][2], "[")
	assert.Equal(t, testArea[3][3], "]")
	assert.Equal(t, testArea[3][4], "[")
	assert.Equal(t, testArea[3][5], "]")
	assert.Equal(t, testArea[4][2], "[")
	assert.Equal(t, testArea[4][3], "]")
	assert.Equal(t, testArea[5][2], ".")
	assert.Equal(t, testArea[5][3], "@")
}

/*
	That's propably not the way to do It.

Ie. hit the wall with the head until It crumbles.
The wall, not the head.
But I won't go on reddit, even If I'm 3 days behind
because of this logic.
[# # # # # # # # # # # # # # # # # # # #]
[# # . ] . . [ ] . . . . . . [ ] [ ] # #]
[# # [ ] . . . . . . . . . . . [ ] . # #]
[# # . . . . . . . . . . . . [ ] [ ] # #]
[# # . . . . . . . . . . . . . . [ ] # #]
[# # . . # # [ ] . . [ ] [ ] [ . @ . # #]
[# # . . . [ ] . . . [ ] . [ ] ] . . # #]
[# # . . . . . [ ] . . [ ] . [ ] [ ] # #]
[# # . . . . . . . . [ ] . . . . . . # #]
[# # # # # # # # # # # # # # # # # # # #]
*/
func TestPushAxisYpart2UpAtY5X4PentaStackWithWallEdgeCaseInDescExample(t *testing.T) {

	path := "./test_data/data_6.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)
	testSolution := NewSolution(dp, path)

	// Altering the data to suit the test
	testSolution.DataProvider.GetArea()[2][2] = BLANK

	// Given # # . . . . . @ [ ] [ ] . . # #
	testArea := testSolution.ConvertPart1AreaToPart2Version()

	// Altering the data to suit the test
	testArea[1][2] = BOX_LEFT
	testArea[1][3] = BOX_RIGHT
	testArea[1][4] = BOX_LEFT
	testArea[1][5] = BOX_RIGHT
	testArea[2][2] = BLANK
	testArea[2][3] = BOX_LEFT
	testArea[2][4] = BOX_RIGHT
	testArea[2][5] = BLANK
	testArea[3][2] = BOX_LEFT
	testArea[3][3] = BOX_RIGHT
	testArea[3][4] = BOX_LEFT
	testArea[3][5] = BOX_RIGHT
	testArea[4][2] = BLANK
	testArea[4][3] = BLANK
	testArea[4][4] = BOX_LEFT
	testArea[4][5] = BOX_RIGHT
	testArea[5][4] = ROBOT
	testArea[5][5] = BLANK

	for _, line := range testArea {
		fmt.Println(line)
	}

	// The wanted method invocation
	testArea = testSolution.PushAxisYpart2(testArea, C_TOP)

	for _, line := range testArea {
		fmt.Println(line)
	}

	// Expected results
	assert.Equal(t, testArea[0][2], "#")
	assert.Equal(t, testArea[0][3], "#")
	assert.Equal(t, testArea[1][2], ".")
	assert.Equal(t, testArea[1][3], ".")
	assert.Equal(t, testArea[2][2], ".")
	assert.Equal(t, testArea[2][3], ".")
	assert.Equal(t, testArea[3][2], "[")
	assert.Equal(t, testArea[3][3], "]")
	assert.Equal(t, testArea[3][4], "[")
	assert.Equal(t, testArea[3][5], "]")
	assert.Equal(t, testArea[4][2], "[")
	assert.Equal(t, testArea[4][3], "]")
	assert.Equal(t, testArea[5][2], ".")
	assert.Equal(t, testArea[5][3], "@")
}

func TestIntegrationPart2CalcBWith480DataSet(t *testing.T) {

	path := "./test_data/data_4.txt"

	io_provider := file_handling.Test{}
	dp := data_provider.NewDataProvider(io_provider)

	testSolution := NewSolution(dp, path)

	// Counting the number of boxes to check for any big errors

	pre_test := 0
	for _, line := range testSolution.DataProvider.GetArea() {
		for _, c := range line {
			if c == "[" {
				pre_test++
			}
			if c == "]" {
				pre_test++
			}
		}
	}

	test_result := testSolution.Part2()

	post_test := 0
	for _, line := range testSolution.DataProvider.GetArea() {
		for _, c := range line {
			if c == "[" {
				pre_test++
			}
			if c == "]" {
				pre_test++
			}
		}
	}

	assert.Equal(t, pre_test, post_test)
	assert.Equal(t, test_result, 9021)

}
