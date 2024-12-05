package solution

import (
	"strconv"

	interfaces "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/data"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/logging"
)

type Solution struct {
	DataProvider interfaces.String2DArray
}

func NewSolution(dataProvider interfaces.String2DArray, path string) *Solution {
	sol := &Solution{DataProvider: dataProvider}
	sol.DataProvider.TakeInput(path)
	return sol
}

/*
First thought is to check each and every spot for possible word location
*/
func is_movement_safe(m_x, m_y, w_x, w_y int) bool {
	// checking for out of bounds
	if w_x < 0 || w_y < 0 || w_x >= m_x || w_y >= m_y {
		return false
	}
	return true
}

func search_part1(w []string, d [][]string, x_m, y_m, s_x, s_y, max_y, max_x, i int) bool {
	var exists []string

	for w_i := range w {
		// If the condition is met, move in search
		if w[w_i] == d[s_y][s_x] {
			// Could be a simple int counting matches, but I've chosen to log the data
			exists = append(exists, string("y"+strconv.Itoa(s_y)+" "+"x"+strconv.Itoa(s_x)))

			// Checks if the length of the word has been reached
			if len(w) == len(exists) {
				return true
			}

			// If the letter matches, move in the wanted direction
			s_y += y_m
			s_x += x_m
			if !is_movement_safe(max_x, max_y, s_x, s_y) {
				logging.Info("XMAS not found in1", "array", exists, "i", i)
				return false
			}
		} else {
			// If the word isn't continous, leave the sh
			logging.Info("XMAS not found in2", "array", exists, "i", i)
			return false
		}
	}
	logging.Info("XMAS found in", "array", exists, "i", i)
	return true
}

func part1_iterative(d [][]string) int {
	/*
		Movements in whatever order
		right(backwards) 1 0
		up 0 -1
		left -1 0
		down 0 1
		up right cross 1 -1
		down right cross 1 1
		up left cross -1 -1
		down left cross -1 1
	*/
	words := 0

	movements_x := []int{1, 0, -1, 0, 1, 1, -1, -1}
	movements_y := []int{0, -1, 0, 1, -1, 1, -1, 1}

	for y := range d {
		for x := range d[y] {
			// potentially each entry is the start of a new word,
			// both normal order or reverse order(thats a lot of iterations!)
			// ouch XD
			for i := 0; i < 8; i++ {
				if search_part1([]string{"X", "M", "A", "S"}, d, movements_x[i], movements_y[i], x, y, len(d), len(d[y]), i) {
					words++
				}
			}
		}
	}
	return words
}

func (s *Solution) SolvePart1() int {
	// Calculating
	data := s.DataProvider.Get2DArray()
	logging.Info("Received the data", "var", data)

	return part1_iterative(data)
}

func search_part2(d [][]string, movements_x []int, movements_y []int, center_x, center_y, max_y, max_x, expected_m, expected_s int) bool {
	existence := map[string]int{}

	for i := 0; i < 4; i++ {
		// Checking all four corners for existence of certain chars
		next_x := center_x + movements_x[i]
		next_y := center_y + movements_y[i]
		if !is_movement_safe(max_x, max_y, next_x, next_y) {
			return false
		}
		existence[d[next_y][next_x]]++
	}

	if existence["M"] == expected_m && existence["S"] == expected_s {
		// If the shape is correct, ie contains 2 M's and 2 S's,
		// then We can exclude the edge-cases
		// ***the diagonal can't be sas or mam.***
		// Here, the exclusion is straight forward, as We have two diagonals
		// If the one is consisting of 2 out of 4 chars + "A", and It is confirmed, that 2 other chars are not the same,
		// then the second diagonal must be consisting of the non "M" chars confirming the edge case
		if d[center_y+movements_y[0]][center_x+movements_x[0]] == "M" && d[center_y+movements_y[3]][center_x+movements_x[3]] == "M" ||
			d[center_y+movements_y[1]][center_x+movements_x[1]] == "M" && d[center_y+movements_y[2]][center_x+movements_x[2]] == "M" {
			return false
		}

		return true
	}
	return false
}

func part2_iterative(d [][]string, M, S int) int {
	/*
		Movements in whatever order
		up right cross 1 -1
		down right cross 1 1
		up left cross -1 -1
		down left cross -1 1
	*/
	words := 0

	movements_x := []int{1, 1, -1, -1}
	movements_y := []int{-1, 1, -1, 1}

	for y := range d {
		for x := range d[y] {
			if d[y][x] == "A" {
				if search_part2(d, movements_x, movements_y, x, y, len(d), len(d[y]), M, S) {
					words++
				}
			}
		}
	}
	return words
}

func (s *Solution) SolvePart2(M, S int) int {
	// Calculating
	data := s.DataProvider.Get2DArray()

	return part2_iterative(data, M, S)
}
