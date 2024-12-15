package solution

import (
	"math"

	interfaces "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/data"
)

type Solution struct {
	DataProvider interfaces.Int2DArray
	Path         string
}

var (
	UPPER_CLICKS_BOUNDARY = 100
	A_CLICK_COST          = 3
	B_CLICK_COST          = 1
)

func NewSolution(dataProvider interfaces.Int2DArray, path string) *Solution {
	sol := &Solution{DataProvider: dataProvider, Path: path}
	sol.DataProvider.TakeInput(path)
	return sol
}

func (s *Solution) calculateClicksA(data []int, b float64) float64 {
	return (float64(data[4])/float64(data[0]) - b*(float64(data[2])/float64(data[0])))
}

func (s *Solution) calculateClicksB(data []int) float64 {
	return ((float64(data[5])*float64(data[0]) - float64(data[4])*float64(data[1])) /
		(float64(data[3])*float64(data[0]) - float64(data[2])*float64(data[1])))
}

func (s *Solution) Part1() int {
	tokens := 0

	// Calculating
	data := s.DataProvider.Get2DArray()
	for _, game := range data {
		// fmt.Println("Trying the game nr: " + strconv.Itoa(i) + " data: " +
		// 	"[" + "x1: " + strconv.Itoa(game[0]) + "," + "y1: " +
		// 	strconv.Itoa(game[1]) + "," + "x2: " + strconv.Itoa(game[2]) +
		// 	"," + "y2: " + strconv.Itoa(game[3]) + "," + "rx: " +
		// 	strconv.Itoa(game[4]) + "," + "ry: " + strconv.Itoa(game[5]) + "]")

		// Each game calculation begins with calculating the amount of B moves
		clicks_B := s.calculateClicksB(game)
		if int(math.Round(clicks_B)) < 0 {
			//fmt.Println(i, "B less than 0", "value", clicks_B)
			continue
		} else if math.Abs(math.Round(clicks_B)-clicks_B) > float64(0.01) {
			//fmt.Println(i, " B Not an integer", "diff", math.Round(clicks_B)-clicks_B, "val", clicks_B)
			continue
		} else if int(math.Round(clicks_B)) > UPPER_CLICKS_BOUNDARY {
			//fmt.Println(i, "B clicks above 100", "value", clicks_B)
			continue
		}

		clicks_A := s.calculateClicksA(game, clicks_B)
		if int(math.Round(clicks_A)) < 0 {
			//fmt.Println(i, "A less than 0", "value", clicks_A)
			continue
		} else if math.Abs(math.Round(clicks_A)-clicks_A) > float64(0.06) {
			//fmt.Println(i, "B Not an integer", "diff", math.Round(clicks_A)-clicks_A, "val", clicks_A)
			continue
		} else if int(math.Round(clicks_A)) > UPPER_CLICKS_BOUNDARY {
			//fmt.Println(i, "A clicks above 100", "value", clicks_A)
			continue
		}

		tokens += int(math.Round(clicks_A))*A_CLICK_COST + int(math.Round(clicks_B))*B_CLICK_COST
	}

	return tokens
}

func (s *Solution) calculateClicksCorrectedA(data []int, b float64) float64 {
	return (float64(data[4]+10000000000000)/float64(data[0]) - b*(float64(data[2])/float64(data[0])))
}

func (s *Solution) calculateClicksCorrectedB(data []int) float64 {
	return ((float64(data[5]+10000000000000)*float64(data[0]) - float64(data[4]+10000000000000)*float64(data[1])) /
		(float64(data[3])*float64(data[0]) - float64(data[2])*float64(data[1])))
}

func (s *Solution) Part2() int {
	tokens := 0

	// Calculating
	data := s.DataProvider.Get2DArray()
	for _, game := range data {

		// Each game calculation begins with calculating the amount of B moves
		clicks_B := s.calculateClicksCorrectedB(game)
		if int(math.Round(clicks_B)) < 0 {
			//fmt.Println(i, "B less than 0", "value", clicks_B)
			continue
		} else if math.Abs(math.Round(clicks_B)-clicks_B) > float64(0.01) {
			//fmt.Println(i, " B Not an integer", "diff", math.Round(clicks_B)-clicks_B, "val", clicks_B)
			continue
		}

		clicks_A := s.calculateClicksCorrectedA(game, clicks_B)
		if int(math.Round(clicks_A)) < 0 {
			continue
		} else if math.Abs(math.Round(clicks_A)-clicks_A) > float64(0.06) {
			continue
		}

		tokens += int(math.Round(clicks_A))*A_CLICK_COST + int(math.Round(clicks_B))*B_CLICK_COST
	}

	return tokens
}
