package solution

import (
	"fmt"
	"strconv"

	interfaces "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/data"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/logging"
)

var (
	// Directions in which We can make steps
	UDLR   = 4
	Y_UDLR = []int{-1, 1, 0, 0}
	X_UDLR = []int{0, 0, -1, 1}
)

type Solution struct {
	M_Y          []int
	M_X          []int
	DataProvider interfaces.Int2DArray
	Path         string
}

func NewSolution(dataProvider interfaces.Int2DArray, path string) *Solution {
	sol := &Solution{DataProvider: dataProvider, Path: path, M_Y: Y_UDLR, M_X: X_UDLR}
	sol.DataProvider.TakeInput(path)
	return sol
}

func (s *Solution) translate_path(area [][]int, path []int) bool {
	translated_path := ""

	for i := 0; i < len(path)-1; i += 2 {
		translated_path += strconv.Itoa(area[path[0+i]][path[1+i]])
	}
	return translated_path == "0123456789"
}

func (s *Solution) path_find_greedy_part_1(paths map[string]int, area [][]int, path []int, y, x int) bool {
	// What will be the base case of this algorithm?
	// When the translated path made of indexes gathered over revursive calls
	// will match the "0123456789" string
	// SCORING
	// It will add the "9" coordinates which the path's lead to
	if s.translate_path(area, path) {
		paths[strconv.Itoa(y)+strconv.Itoa(x)] = 1
		return true
	}

	// Recursive check
	for i := 0; i < UDLR; i++ {

		// "Is next step safe"
		if y+s.M_Y[i] >= 0 && y+s.M_Y[i] < len(area) && x+s.M_X[i] >= 0 && x+s.M_X[i] < len(area[0]) {
			if area[y+s.M_Y[i]][x+s.M_X[i]]-area[y][x] == 1 {
				path = append(path, y+s.M_Y[i])
				path = append(path, x+s.M_X[i])

				s.path_find_greedy_part_1(paths, area, path, y+s.M_Y[i], x+s.M_X[i])

				path = path[:len(path)-2]
			}
		}

	}

	return false
}

func (s *Solution) Part1() int {
	height_map := s.DataProvider.Get2DArray()
	// So, by having a big height map
	// I want to search for starting points
	// and pathfind to the highest spot
	// returning the requested scoring ways I can achieve
	// the highest spot from the selected starting point
	value := 0

	for row := 0; row < len(height_map); row++ {
		for col := 0; col < len(height_map[0]); col++ {
			if height_map[row][col] == 0 {
				paths := map[string]int{}
				empty_path := []int{row, col}
				s.path_find_greedy_part_1(paths, height_map, empty_path, row, col)
				value += len(paths)
			}
		}
	}

	return value
}

func (s *Solution) join_int_path_to_string(path []int) string {
	str := ""
	for i := 0; i < len(path); i++ {
		str += strconv.Itoa(path[i])
	}
	return str
}

func (s *Solution) path_find_greedy_part_2(paths map[string]int, area [][]int, path []int, y, x int) bool {
	// What will be the base case of this algorithm?
	// When the translated path made of indexes gathered over revursive calls
	// will match the "0123456789" string
	// SCORING
	// It will add the whole path that has lead to "9"
	if s.translate_path(area, path) {
		logging.Info("Path to join", "path", path)
		paths[s.join_int_path_to_string(path)] = 1
		return true
	}

	// Recursive check
	for i := 0; i < UDLR; i++ {

		// "Is next step safe"
		if y+s.M_Y[i] >= 0 && y+s.M_Y[i] < len(area) && x+s.M_X[i] >= 0 && x+s.M_X[i] < len(area[0]) {
			if area[y+s.M_Y[i]][x+s.M_X[i]]-area[y][x] == 1 {
				path = append(path, y+s.M_Y[i])
				path = append(path, x+s.M_X[i])

				s.path_find_greedy_part_2(paths, area, path, y+s.M_Y[i], x+s.M_X[i])

				path = path[:len(path)-2]
			}
		}

	}

	return false
}

func (s *Solution) Part2() int {
	height_map := s.DataProvider.Get2DArray()
	// So, by having a big height map
	// I want to search for starting points
	// and pathfind to the highest spot
	// returning the requested scoring ways I can achieve
	// the highest spot from the selected starting point
	value := 0

	for row := 0; row < len(height_map); row++ {
		for col := 0; col < len(height_map[0]); col++ {
			if height_map[row][col] == 0 {
				paths := map[string]int{}
				empty_path := []int{row, col}
				s.path_find_greedy_part_2(paths, height_map, empty_path, row, col)
				fmt.Println(paths)
				value += len(paths)
			}
		}
	}

	return value
}
