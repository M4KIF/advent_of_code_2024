package solution

import (
	"strconv"

	interfaces "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/data"
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
	return &Solution{DataProvider: dataProvider, Path: path, M_Y: Y_UDLR, M_X: X_UDLR}
}

func (s *Solution) translate_path(area [][]int, path []int) bool {
	translated_path := ""

	for i := 0; i < len(path)-1; i += 2 {
		translated_path += strconv.Itoa(area[path[0+i]][path[1+i]])
	}
	return translated_path == "0123456789"
}

func (s *Solution) path_find_greedy(paths map[string]int, area [][]int, path []int, y, x int) bool {
	// What will be the base case of this algorithm?
	// When the translated path made of indexes gathered over revursive calls
	// will match the "0123456789" string
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

				s.path_find_greedy(paths, area, path, y+s.M_Y[i], x+s.M_X[i])

				path = path[:len(path)-2]
			}
		}

	}

	return false
}

func (s *Solution) helper() int {
	height_map := s.DataProvider.Get2DArray()
	// So, by having a big height map
	// I want to search for starting points
	// and pathfind to the highest spot
	// returning the ways I can achieve
	// the highest spot from the selected starting point

	// The simplest here will be to loop over the data
	// As I will achieve both the path find and the coordinates retrieval
	value := 0

	for row := 0; row < len(height_map); row++ {
		for col := 0; col < len(height_map[0]); col++ {
			if height_map[row][col] == 0 {
				paths := map[string]int{}
				empty_path := []int{row, col}
				s.path_find_greedy(paths, height_map, empty_path, row, col)
				value += len(paths)
			}
		}
	}

	return value
}

func (s *Solution) Solve() int {
	// Collecting the data
	s.DataProvider.TakeInput(s.Path)

	return s.helper()
}
