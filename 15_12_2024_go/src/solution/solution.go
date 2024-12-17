package solution

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strings"

	interfaces "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/data"
)

type Solution struct {
	DataProvider interfaces.String2Dand1Darray
	Path         string
}

var (
	// Materials
	WALL  = "#"
	ROBOT = "@"
	BOX   = "O"
	BLANK = "."

	BOX_LEFT  = "["
	BOX_RIGHT = "]"

	// Directions chars
	C_TOP   = "^"
	C_RIGHT = ">"
	C_DOWN  = "v"
	C_LEFT  = "<"

	// Directions indexes
	I_TOP   = 2
	I_RIGHT = 0
	I_DOWN  = 3
	I_LEFT  = 1

	// Movement Arrays
	movement_y = []int{0, 0, -1, 1}
	movement_x = []int{1, -1, 0, 0}
)

func NewSolution(dataProvider interfaces.String2Dand1Darray, path string) *Solution {
	sol := &Solution{DataProvider: dataProvider, Path: path}
	sol.DataProvider.TakeInput(path)
	return sol
}

func (s *Solution) IsPointSafe(area [][]string, y, x int) bool {
	return (y >= 0 && y < len(area) &&
		x >= 0 && x < len(area[0]))
}

func (s *Solution) FindRobotPosition(area [][]string) (int, int) {
	atMatch, _ := regexp.Compile("[@]{1}")

	for y, line := range area {
		if x := slices.IndexFunc(line, func(c string) bool {
			return atMatch.Match([]byte(c))
		}); x != -1 {
			return y, x
		}
	}

	return -1, -1
}

func (s *Solution) SearchForWallAndBoxesAtY(y, x int, direction string) (int, int) {
	// Oribl
	if !s.IsPointSafe(s.DataProvider.GetArea(), y, x) {
		return -1, -1
	}

	boxes_count := 0

	wallMatch, _ := regexp.Compile("#")
	boxesMatch, _ := regexp.Compile("O")

	switch direction {
	case C_TOP:
		for t_y := y; t_y >= 0; t_y-- {
			if boxesMatch.Match([]byte(s.DataProvider.GetArea()[t_y][x])) {
				boxes_count++
			}
			if wallMatch.Match([]byte(s.DataProvider.GetArea()[t_y][x])) {
				return t_y, boxes_count
			}
		}
	case C_DOWN:
		for t_y := y; t_y < len(s.DataProvider.GetArea()); t_y++ {
			if boxesMatch.Match([]byte(s.DataProvider.GetArea()[t_y][x])) {
				boxes_count++
			}
			if wallMatch.Match([]byte(s.DataProvider.GetArea()[t_y][x])) {
				return t_y, boxes_count
			}
		}
	}

	return -1, -1
}

func (s *Solution) PushAxisY(d string) bool {

	y, x := s.FindRobotPosition(s.DataProvider.GetArea())

	wall_y, boxes_count := s.SearchForWallAndBoxesAtY(y, x, d)

	delta_y := int(math.Abs(math.Abs(float64(wall_y)) - math.Abs(float64(y))))

	blank_spaces := delta_y - boxes_count

	fmt.Println("y: ", y, " x:", x, " wall_x: ", wall_y, " boxes_count: ",
		boxes_count, " delta_x: ", delta_y, " blanks: ", blank_spaces)

	// Stands near/at a wall
	if delta_y <= 1 {
		return false
	}

	// Already pushed what was possible to push
	if blank_spaces == 0 {
		return false
	}

	switch d {
	case C_TOP:
		// if boxes_count <= 0 || s.DataProvider.GetArea()[y+movement_y[I_TOP]][x] != BOX {
		// 	s.DataProvider.GetArea()[y][x] = BLANK
		// 	s.DataProvider.GetArea()[y+movement_y[I_TOP]][x] = ROBOT
		// 	return true
		// }
		// for i := wall_y + 1; i <= y; i++ {
		// 	if boxes_count > 0 {
		// 		s.DataProvider.GetArea()[i][x] = BOX
		// 		boxes_count--
		// 	} else if boxes_count == 0 {
		// 		s.DataProvider.GetArea()[i][x] = ROBOT
		// 		boxes_count--
		// 	} else {
		// 		s.DataProvider.GetArea()[i][x] = BLANK
		// 	}
		// }
		// fmt.Println("line expected: ", s.DataProvider.GetArea()[y])

		// Convert to a vertical slice. TOP-DOWN
		vertical_line := ""
		for _, line := range s.DataProvider.GetArea() {
			vertical_line += line[x]
		}

		boxesMatch, _ := regexp.Compile("O+")
		boxes := boxesMatch.FindAllString(vertical_line, -1)
		indexes := boxesMatch.FindAllStringIndex(vertical_line, -1)

		if boxes == nil {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y][x+movement_x[I_RIGHT]] = ROBOT
			return true
		}

		// Filtering the indexes with garbage collection
		filter := indexes[:0]
		filtered_boxes := boxes[:0]

		for i, candidate := range indexes {
			if candidate[1] <= y {
				filter = append(filter, candidate)
				filtered_boxes = append(filtered_boxes, boxes[i])
			}
		}

		for i := len(filter); i < len(indexes); i++ {
			indexes[i] = nil // or the zero value of T
		}

		// If all of the boxes are on the opposite side, move and fajrant
		if len(filter) == 0 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y+movement_y[I_TOP]][x] = ROBOT
			return true
		}

		fmt.Println("TOPSTART! ", filter)
		fmt.Println(filtered_boxes)

		// If the boxes are further than 1 from the robot, move the robot and fajrant
		if filter[len(filter)-1][1] != y {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y+movement_y[I_TOP]][x] = ROBOT
			fmt.Println("ASDASDSDASD")
			return true
		} else if s.DataProvider.GetArea()[y][filter[len(filter)-1][0]-1] == WALL {
			return false
		} else {
			// Move the robot on the first found index
			s.DataProvider.GetArea()[filter[len(filter)-1][1]-1][x] = ROBOT
			s.DataProvider.GetArea()[filter[len(filter)-1][1]][x] = BLANK

			fmt.Println("ASDASD")
			// Shift the boxes by one to the right
			for i := 0; i < len(filtered_boxes[len(filtered_boxes)-1]); i++ {
				s.DataProvider.GetArea()[filter[len(filter)-1][0]-1+i][x] = BOX
			}

		}
	case C_DOWN:
		vertical_line := ""
		for _, line := range s.DataProvider.GetArea() {
			vertical_line += line[x]
		}

		boxesMatch, _ := regexp.Compile("O+")
		boxes := boxesMatch.FindAllString(vertical_line, -1)
		indexes := boxesMatch.FindAllStringIndex(vertical_line, -1)

		if boxes == nil {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y][x+movement_x[I_RIGHT]] = ROBOT
			return true
		}

		// Filtering the indexes with garbage collection
		filter := indexes[:0]
		filtered_boxes := boxes[:0]

		for i, candidate := range indexes {
			if candidate[0] > y {
				filter = append(filter, candidate)
				filtered_boxes = append(filtered_boxes, boxes[i])
			}
		}

		for i := len(filter); i < len(indexes); i++ {
			indexes[i] = nil // or the zero value of T
		}

		// If all of the boxes are on the opposite side, move and fajrant
		if len(filter) == 0 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y+movement_y[I_DOWN]][x] = ROBOT
			return true
		}

		fmt.Println("DOWNSTART! ", filter)
		fmt.Println(filtered_boxes)

		// If the boxes are further than 1 from the robot, move the robot and fajrant
		if math.Abs(float64(filter[0][0]-y)) != 1 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y+movement_y[I_DOWN]][x] = ROBOT
			fmt.Println("ASDASDSDASD")
			return true
		} else if s.DataProvider.GetArea()[filter[0][1]][x] == WALL {
			fmt.Println("ASDASDSDASDsdfsadf")
			return false
		} else {
			// Move the robot on the first found index
			s.DataProvider.GetArea()[filter[0][0]][x] = ROBOT
			s.DataProvider.GetArea()[filter[0][0]-1][x] = BLANK

			fmt.Println("ASDASD")
			// Shift the boxes by one to the right
			for i := 0; i < len(filtered_boxes[len(filtered_boxes)-1]); i++ {
				s.DataProvider.GetArea()[filter[0][0]+1+i][x] = BOX
			}

		}
	}

	return true
}

func (s *Solution) PushAxisYnew(d string) bool {

	y, x := s.FindRobotPosition(s.DataProvider.GetArea())

	switch d {
	case C_TOP:
		if s.DataProvider.GetArea()[y+movement_y[I_TOP]][x] == WALL {
			return false
		}
		// if boxes_count <= 0 || s.DataProvider.GetArea()[y+movement_y[I_TOP]][x] != BOX {
		// 	s.DataProvider.GetArea()[y][x] = BLANK
		// 	s.DataProvider.GetArea()[y+movement_y[I_TOP]][x] = ROBOT
		// 	return true
		// }
		// for i := wall_y + 1; i <= y; i++ {
		// 	if boxes_count > 0 {
		// 		s.DataProvider.GetArea()[i][x] = BOX
		// 		boxes_count--
		// 	} else if boxes_count == 0 {
		// 		s.DataProvider.GetArea()[i][x] = ROBOT
		// 		boxes_count--
		// 	} else {
		// 		s.DataProvider.GetArea()[i][x] = BLANK
		// 	}
		// }
		// fmt.Println("line expected: ", s.DataProvider.GetArea()[y])

		// Convert to a vertical slice. TOP-DOWN
		vertical_line := ""
		for _, line := range s.DataProvider.GetArea() {
			vertical_line += line[x]
		}

		boxesMatch, _ := regexp.Compile("O+")
		boxes := boxesMatch.FindAllString(vertical_line, -1)
		indexes := boxesMatch.FindAllStringIndex(vertical_line, -1)

		// if len(boxes) == 0 {
		// 	s.DataProvider.GetArea()[y][x] = BLANK
		// 	s.DataProvider.GetArea()[y+movement_y[I_TOP]][x] = ROBOT
		// 	return true
		// }

		// Filtering the indexes with garbage collection
		filter := indexes[:0]
		filtered_boxes := boxes[:0]

		for i, candidate := range indexes {
			if candidate[1] <= y {
				filter = append(filter, candidate)
				filtered_boxes = append(filtered_boxes, boxes[i])
			}
		}

		for i := len(filter); i < len(indexes); i++ {
			indexes[i] = nil // or the zero value of T
		}

		// If all of the boxes are on the opposite side, move and fajrant
		if len(filter) == 0 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y+movement_y[I_TOP]][x] = ROBOT
			return true
		}

		//fmt.Println("TOPSTART! ", filter)
		//fmt.Println(filtered_boxes)
		//fmt.Println("Edge case", s.DataProvider.GetArea()[filter[len(filter)-1][0]-1][x])

		// If the boxes are further than 1 from the robot, move the robot and fajrant
		if filter[len(filter)-1][1] != y {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y+movement_y[I_TOP]][x] = ROBOT
			//fmt.Println("ASDASDSDASD")
			return true
		} else if s.DataProvider.GetArea()[filter[len(filter)-1][0]-1][x] == WALL {
			//fmt.Println("ASDASDASFEFWW#CWECWE")
			return false
		} else {
			// Move the robot on the first found index
			s.DataProvider.GetArea()[filter[len(filter)-1][1]-1][x] = ROBOT
			s.DataProvider.GetArea()[filter[len(filter)-1][1]][x] = BLANK

			//fmt.Println("ASDASDD")
			// Shift the boxes by one to the right
			for i := 0; i < len(filtered_boxes[len(filtered_boxes)-1]); i++ {
				s.DataProvider.GetArea()[filter[len(filter)-1][0]-1+i][x] = BOX
			}

		}
	case C_DOWN:
		if s.DataProvider.GetArea()[y+movement_y[I_DOWN]][x] == WALL {
			return false
		}
		vertical_line := ""
		for _, line := range s.DataProvider.GetArea() {
			vertical_line += line[x]
		}

		boxesMatch, _ := regexp.Compile("O+")
		boxes := boxesMatch.FindAllString(vertical_line, -1)
		indexes := boxesMatch.FindAllStringIndex(vertical_line, -1)

		// if len(boxes) == 0 {
		// 	s.DataProvider.GetArea()[y][x] = BLANK
		// 	s.DataProvider.GetArea()[y+movement_y[I_DOWN]][x] = ROBOT
		// 	return true
		// }

		// Filtering the indexes with garbage collection
		filter := indexes[:0]
		filtered_boxes := boxes[:0]

		for i, candidate := range indexes {
			if candidate[0] > y {
				filter = append(filter, candidate)
				filtered_boxes = append(filtered_boxes, boxes[i])
			}
		}

		for i := len(filter); i < len(indexes); i++ {
			indexes[i] = nil // or the zero value of T
		}

		// If all of the boxes are on the opposite side, move and fajrant
		if len(filter) == 0 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y+movement_y[I_DOWN]][x] = ROBOT
			return true
		}

		//fmt.Println("DOWNSTART! ", filter)
		//fmt.Println(filtered_boxes)

		// If the boxes are further than 1 from the robot, move the robot and fajrant
		if math.Abs(float64(filter[0][0]-y)) != 1 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y+movement_y[I_DOWN]][x] = ROBOT
			//fmt.Println("ASDASDSDASD")
			return true
		} else if s.DataProvider.GetArea()[filter[0][1]][x] == WALL {
			//fmt.Println("ASDASDSDASDsdfsadf")
			return false
		} else {
			// Move the robot on the first found index
			s.DataProvider.GetArea()[filter[0][0]][x] = ROBOT
			s.DataProvider.GetArea()[filter[0][0]-1][x] = BLANK

			//fmt.Println("ASDASDAAA")
			// Shift the boxes by one to the right
			for i := 0; i < len(filtered_boxes[0]); i++ {
				s.DataProvider.GetArea()[filter[0][0]+1+i][x] = BOX
			}

		}
	}

	return true
}

func (s *Solution) SearchForWallAndBoxesAtX(y, x int, direction string) (int, int) {
	// Oribl
	if !s.IsPointSafe(s.DataProvider.GetArea(), y, x) {
		return -1, -1
	}

	boxes_count := 0

	wallMatch, _ := regexp.Compile("#")
	boxesMatch, _ := regexp.Compile("O")

	switch direction {
	case C_RIGHT:
		for t_x := x; t_x < len(s.DataProvider.GetArea()[y]); t_x++ {
			if boxesMatch.Match([]byte(s.DataProvider.GetArea()[y][t_x])) {
				boxes_count++
			}
			if wallMatch.Match([]byte(s.DataProvider.GetArea()[y][t_x])) {
				return t_x, boxes_count
			}
		}
	case C_LEFT:
		for t_x := x; t_x >= 0; t_x-- {
			if boxesMatch.Match([]byte(s.DataProvider.GetArea()[y][t_x])) {
				boxes_count++
			}
			if wallMatch.Match([]byte(s.DataProvider.GetArea()[y][t_x])) {
				return t_x, boxes_count
			}
		}
	}

	return -1, -1
}

func (s *Solution) SearchForWallAndBoxesAtXw(y, x int, direction string) (int, int) {
	// Oribl
	if !s.IsPointSafe(s.DataProvider.GetArea(), y, x) {
		return -1, -1
	}

	boxes_count := 0

	wallMatch, _ := regexp.Compile("#")
	boxesMatch, _ := regexp.Compile("O+")

	for _, c := range "TEST" {
		fmt.Println("RUNE", c)
	}

	switch direction {
	case C_RIGHT:
		slices.Concat(s.DataProvider.GetArea()[y])
		for t_x := x; t_x < len(s.DataProvider.GetArea()[y]); t_x++ {
			if boxesMatch.Match([]byte(s.DataProvider.GetArea()[y][t_x])) {
				boxes_count++
			}
			if wallMatch.Match([]byte(s.DataProvider.GetArea()[y][t_x])) {
				return t_x, boxes_count
			}
		}
	case C_LEFT:
		for t_x := x; t_x >= 0; t_x-- {
			if boxesMatch.Match([]byte(s.DataProvider.GetArea()[y][t_x])) {
				boxes_count++
			}
			if wallMatch.Match([]byte(s.DataProvider.GetArea()[y][t_x])) {
				return t_x, boxes_count
			}
		}
	}

	return -1, -1
}

func (s *Solution) PushAxisXnew(d string) bool {
	/*
		I see It this way.
		- Seek the position of the robot
		- Seek the nearest wall to the specified direction
		- While seeking the wall, count boxes that can be shifted
		- Calculate the delta distance between the robot and the wall
		- delta will be equal to empty space + robot + boxes.
		- rearrange the room in the span of this delta to first have the boxes
		- then robot
		- then empty space
	*/

	y, x := s.FindRobotPosition(s.DataProvider.GetArea())

	switch d {
	case C_RIGHT:
		if s.DataProvider.GetArea()[y][x+movement_x[I_RIGHT]] == WALL {
			return false
		}
		boxesMatch, _ := regexp.Compile("O+")
		boxes := boxesMatch.FindAllString(strings.Join(s.DataProvider.GetArea()[y], ""), -1)
		indexes := boxesMatch.FindAllStringIndex(strings.Join(s.DataProvider.GetArea()[y], ""), -1)

		if len(boxes) == 0 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y][x+movement_x[I_RIGHT]] = ROBOT
			return true
		}

		// Filtering the indexes with garbage collection
		filter := indexes[:0]
		filtered_boxes := boxes[:0]

		for i, candidate := range indexes {
			if candidate[0] > x {
				filter = append(filter, candidate)
				filtered_boxes = append(filtered_boxes, boxes[i])
			}
		}

		for i := len(filter); i < len(indexes); i++ {
			indexes[i] = nil // or the zero value of T
		}

		// If all of the boxes are on the opposite side, move and fajrant
		if len(filter) == 0 || len(filtered_boxes) == 0 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y][x+movement_x[I_RIGHT]] = ROBOT
			return true
		}

		// If the boxes are further than 1 from the robot, move the robot and fajrant
		if math.Abs(float64(filter[0][0]-x)) != 1 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y][x+movement_x[I_RIGHT]] = ROBOT
			return true
		}

		if s.DataProvider.GetArea()[y][filter[0][1]] == WALL {
			// If the index succeeding the found box(es) is a wall, get out
			return false
		}

		// Move the robot on the first found index
		s.DataProvider.GetArea()[y][indexes[0][0]] = ROBOT
		s.DataProvider.GetArea()[y][x] = BLANK

		// Shift the boxes by one to the right
		for i := 0; i < len(filtered_boxes[0]); i++ {
			s.DataProvider.GetArea()[y][filter[0][1]-i] = BOX
		}

		// // That is wrong
		// for i := wall_x - 1; i >= x; i-- {
		// 	if boxes_count > 0 {
		// 		s.DataProvider.GetArea()[y][i] = BOX
		// 		boxes_count--
		// 	} else if boxes_count == 0 {
		// 		s.DataProvider.GetArea()[y][i] = ROBOT
		// 		boxes_count--
		// 	} else {
		// 		s.DataProvider.GetArea()[y][i] = BLANK
		// 	}
		// }
		//fmt.Println("line expected: ", s.DataProvider.GetArea()[y])
	case C_LEFT:
		if s.DataProvider.GetArea()[y][x+movement_x[I_LEFT]] == WALL {
			return false
		}

		boxesMatch, _ := regexp.Compile("O+")
		boxes := boxesMatch.FindAllString(strings.Join(s.DataProvider.GetArea()[y], ""), -1)
		indexes := boxesMatch.FindAllStringIndex(strings.Join(s.DataProvider.GetArea()[y], ""), -1)

		// Filtering the indexes with garbage collection
		filter := indexes[:0]
		filtered_boxes := boxes[:0]

		for i, candidate := range indexes {
			if candidate[1] <= x {
				filter = append(filter, candidate)
				filtered_boxes = append(filtered_boxes, boxes[i])
			}
		}

		for i := len(filter); i < len(indexes); i++ {
			indexes[i] = nil // or the zero value of T
		}

		fmt.Println(filter)
		fmt.Println(filtered_boxes)

		// If all of the boxes are on the opposite side, move and fajrant
		if len(filter) == 0 || len(filtered_boxes) == 0 {
			//fmt.Println("AASDASDSDASDQWW")
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y][x+movement_x[I_LEFT]] = ROBOT
			return true
		}

		if math.Abs(float64(filter[len(filter)-1][1]-1-x)) != 1 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y][x+movement_x[I_LEFT]] = ROBOT
			return true
		}

		// If the index succeeding the found box(es) is a wall, get out
		if s.DataProvider.GetArea()[y][filter[len(filter)-1][0]-1] == WALL {
			//fmt.Println("ASDASDQWW")
			return false
		}

		// Move the robot on the first found index
		s.DataProvider.GetArea()[y][filter[len(filter)-1][1]-1] = ROBOT
		s.DataProvider.GetArea()[y][x] = BLANK

		//fmt.Println("asdasd131231231a")
		// Shift the boxes by one to the right
		for i := 0; i < len(filtered_boxes[len(filtered_boxes)-1]); i++ {
			s.DataProvider.GetArea()[y][filter[len(filter)-1][1]-2-i] = BOX
		}

	}

	return true
}

func (s *Solution) PushAxisX(d string) bool {
	/*
		I see It this way.
		- Seek the position of the robot
		- Seek the nearest wall to the specified direction
		- While seeking the wall, count boxes that can be shifted
		- Calculate the delta distance between the robot and the wall
		- delta will be equal to empty space + robot + boxes.
		- rearrange the room in the span of this delta to first have the boxes
		- then robot
		- then empty space
	*/

	y, x := s.FindRobotPosition(s.DataProvider.GetArea())

	wall_x, boxes_count := s.SearchForWallAndBoxesAtX(y, x, d)

	delta_x := int(math.Abs(math.Abs(float64(wall_x)) - math.Abs(float64(x))))

	blank_spaces := delta_x - boxes_count

	fmt.Println("y: ", y, " x:", x, " wall_x: ", wall_x, " boxes_count: ",
		boxes_count, " delta_x: ", delta_x, " blanks: ", blank_spaces)

	// Stands near/at a wall
	if delta_x <= 1 {
		return false
	}

	// Already pushed what was possible to push
	if blank_spaces == 0 {
		return false
	}

	if boxes_count <= 0 {

	}

	switch d {
	case C_RIGHT:
		boxesMatch, _ := regexp.Compile("O+")
		boxes := boxesMatch.FindAllString(strings.Join(s.DataProvider.GetArea()[y], ""), -1)
		indexes := boxesMatch.FindAllStringIndex(strings.Join(s.DataProvider.GetArea()[y], ""), -1)

		if boxes == nil {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y][x+movement_x[I_RIGHT]] = ROBOT
			return true
		}

		// Filtering the indexes with garbage collection
		filter := indexes[:0]
		filtered_boxes := boxes[:0]

		for i, candidate := range indexes {
			if candidate[0] > x {
				filter = append(filter, candidate)
				filtered_boxes = append(filtered_boxes, boxes[i])
			}
		}

		for i := len(filter); i < len(indexes); i++ {
			indexes[i] = nil // or the zero value of T
		}

		// If all of the boxes are on the opposite side, move and fajrant
		if len(filter) == 0 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y][x+movement_x[I_RIGHT]] = ROBOT
			return true
		}

		// If the index succeeding the found box(es) is a wall, get out
		if s.DataProvider.GetArea()[y][filter[0][1]] == WALL {
			return false
		}

		// If the boxes are further than 1 from the robot, move the robot and fajrant
		if math.Abs(float64(filter[0][0]-x)) != 1 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y][x+movement_x[I_RIGHT]] = ROBOT
			return true
		} else {
			// Move the robot on the first found index
			s.DataProvider.GetArea()[y][indexes[0][0]] = ROBOT
			s.DataProvider.GetArea()[y][x] = BLANK

			// Shift the boxes by one to the right
			for i := 0; i < len(filtered_boxes[0]); i++ {
				s.DataProvider.GetArea()[y][filter[0][1]-i] = BOX
			}

		}

		// // That is wrong
		// for i := wall_x - 1; i >= x; i-- {
		// 	if boxes_count > 0 {
		// 		s.DataProvider.GetArea()[y][i] = BOX
		// 		boxes_count--
		// 	} else if boxes_count == 0 {
		// 		s.DataProvider.GetArea()[y][i] = ROBOT
		// 		boxes_count--
		// 	} else {
		// 		s.DataProvider.GetArea()[y][i] = BLANK
		// 	}
		// }
		fmt.Println("line expected: ", s.DataProvider.GetArea()[y])
	case C_LEFT:
		boxesMatch, _ := regexp.Compile("O+")
		boxes := boxesMatch.FindAllString(strings.Join(s.DataProvider.GetArea()[y], ""), -1)
		indexes := boxesMatch.FindAllStringIndex(strings.Join(s.DataProvider.GetArea()[y], ""), -1)

		if boxes == nil {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y][x+movement_x[I_LEFT]] = ROBOT
			return true
		}

		// Filtering the indexes with garbage collection
		filter := indexes[:0]
		filtered_boxes := boxes[:0]

		for i, candidate := range indexes {
			if candidate[1] <= x {
				filter = append(filter, candidate)
				filtered_boxes = append(filtered_boxes, boxes[i])
			}
		}

		for i := len(filter); i < len(indexes); i++ {
			indexes[i] = nil // or the zero value of T
		}

		fmt.Println(filter)
		fmt.Println(filtered_boxes)

		// If all of the boxes are on the opposite side, move and fajrant
		if len(filter) == 0 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y][x+movement_x[I_LEFT]] = ROBOT
			return true
		}

		if len(filtered_boxes) == 0 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y][x+movement_x[I_LEFT]] = ROBOT
			return true
		}

		// If the index succeeding the found box(es) is a wall, get out
		if s.DataProvider.GetArea()[y][filter[len(filter)-1][0]-1] == WALL {
			return false
		}

		// If the boxes are further than 1 from the robot, move the robot and fajrant
		if math.Abs(float64(filter[len(filter)-1][1]-1-x)) != 1 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y][x+movement_x[I_LEFT]] = ROBOT
			return true
		} else {
			// Move the robot on the first found index
			s.DataProvider.GetArea()[y][filter[len(filter)-1][1]-1] = ROBOT
			s.DataProvider.GetArea()[y][x] = BLANK

			fmt.Println("asdasd131231231a")
			// Shift the boxes by one to the right
			for i := 0; i < len(filtered_boxes[len(filtered_boxes)-1]); i++ {
				s.DataProvider.GetArea()[y][filter[len(filter)-1][0]-1] = BOX
			}

		}
		// if boxes_count <= 0 || s.DataProvider.GetArea()[y][x+movement_x[I_LEFT]] != BOX {
		// 	s.DataProvider.GetArea()[y][x] = BLANK
		// 	s.DataProvider.GetArea()[y][x+movement_x[I_LEFT]] = ROBOT
		// 	return true
		// }

		// // That is wrong
		// for i := wall_x + 1; i <= x; i++ {
		// 	if boxes_count > 0 {
		// 		s.DataProvider.GetArea()[y][i] = BOX
		// 		boxes_count--
		// 	} else if boxes_count == 0 {
		// 		s.DataProvider.GetArea()[y][i] = ROBOT
		// 		boxes_count--
		// 	} else {
		// 		s.DataProvider.GetArea()[y][i] = BLANK
		// 	}
		// }
	}

	return true
}

func (s *Solution) Part1() int {
	for _, command := range s.DataProvider.GetCommands() {
		// if i > 5 {
		// 	break
		// }
		//y, x := s.FindRobotPosition(s.DataProvider.GetArea())
		switch command {
		case C_TOP:
			//
			fmt.Println("Command top")
			// if s.DataProvider.GetArea()[y+movement_y[I_TOP]][x] == WALL {
			// 	continue
			// }
			s.PushAxisYnew(C_TOP)

		case C_RIGHT:
			//
			fmt.Println("Command right")
			// if s.DataProvider.GetArea()[y][x+movement_x[I_RIGHT]] == WALL {
			// 	continue
			// }
			s.PushAxisXnew(C_RIGHT)

		case C_DOWN:
			//
			fmt.Println("Command down")
			// if s.DataProvider.GetArea()[y+movement_y[I_DOWN]][x] == WALL {
			// 	continue
			// }
			s.PushAxisYnew(C_DOWN)

		case C_LEFT:
			//
			fmt.Println("Command left")
			// if s.DataProvider.GetArea()[y][x+movement_x[I_LEFT]] == WALL {
			// 	continue
			// }
			s.PushAxisXnew(C_LEFT)
		}
		// for _, line := range s.DataProvider.GetArea() {
		// 	fmt.Println(line)
		// }
	}

	res := 0

	for y, line := range s.DataProvider.GetArea() {
		for x, c := range line {
			if c == "O" {
				res += y*100 + x
			}
		}
	}

	return res
}

/*
PART 2
*/

func (s *Solution) ConvertPart1AreaToPart2Version() [][]string {

	part2_area := [][]string{}

	// Line by line
	for _, line := range s.DataProvider.GetArea() {
		temp := []string{}
		for x, c := range line {
			// Altering the walls
			if c == "#" && x < len(line) {
				temp = append(temp, "#")
				temp = append(temp, "#")
			} else if c == "O" {
				temp = append(temp, "[")
				temp = append(temp, "]")
			} else if c == "." {
				temp = append(temp, ".")
				temp = append(temp, ".")
			} else if c == "@" {
				temp = append(temp, "@")
				temp = append(temp, ".")
			}

		}
		part2_area = append(part2_area, temp)
	}

	return part2_area
}

/*


   # Flood fill (with pattern checking(?))
   def find_group(self, plant, points: dict, stack: list, visited):
       if len(stack) > 0:
           vertice = stack.pop()

           if self.area[vertice[0]][vertice[1]] == plant:
               points[(vertice[0], vertice[1])] = 1
               visited[vertice[0]][vertice[1]]+=1
           else:
               return

           for move in self.movement:

               n_y = vertice[0] + move[0]
               n_x = vertice[1] + move[1]
               if (n_y >=0 and n_y < len(self.area) and n_x >= 0 and n_x < len(self.area[0])):
                   if self.area[n_y][n_x] == plant and points.get((n_y, n_x)) is None:
                       stack.append((n_y, n_x))
           self.find_group(plant, points, stack, visited)

*/

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Push(point T) {
	s.items = append(s.items, point)
}

// LIFO
func (s *Stack[T]) Pop() T {
	// That is the drawback of generics, safety check in code, not in DS
	point := s.items[0]
	s.items = s.items[1:]
	return point
}

func (s *Solution) IsInBoxesGroup(area [][]string, y, x, starting_y int, dir string) bool {

	if dir == C_DOWN {
		if y <= starting_y {
			return false
		}
		if area[y][x] == BOX_LEFT && area[y+1][x] == BOX_LEFT ||
			area[y][x] == BOX_RIGHT && area[y+1][x] == BOX_RIGHT {
			return true
		}

		// Halfstacking down-movement
		if area[y][x] == BOX_LEFT && area[y-1][x+1] == BOX_LEFT ||
			area[y][x] == BOX_RIGHT && area[y-1][x] == BOX_LEFT ||
			area[y][x] == BOX_RIGHT && area[y-1][x-1] == BOX_RIGHT ||
			area[y][x] == BOX_LEFT && area[y-1][x] == BOX_RIGHT {
			return true
		}
	}

	if dir == C_TOP {
		if y >= starting_y {
			return false
		}

		if area[y][x] == BOX_LEFT && area[y-1][x] == BOX_LEFT ||
			area[y][x] == BOX_RIGHT && area[y-1][x] == BOX_RIGHT {
			return true
		}

		// Halfstacking up-movement
		if area[y][x] == BOX_LEFT && area[y+1][x-1] == BOX_LEFT ||
			area[y][x] == BOX_LEFT && area[y+1][x] == BOX_RIGHT ||
			area[y][x] == BOX_RIGHT && area[y+1][x+1] == BOX_RIGHT ||
			area[y][x] == BOX_RIGHT && area[y+1][x] == BOX_LEFT {
			return true
		}
	}

	// // Up down columns, without half-stacking
	// if area[y][x] == BOX_LEFT && area[y-1][x] == BOX_LEFT ||
	// 	area[y][x] == BOX_RIGHT && area[y-1][x] == BOX_RIGHT ||
	// 	area[y][x] == BOX_LEFT && area[y+1][x] == BOX_LEFT ||
	// 	area[y][x] == BOX_RIGHT && area[y+1][x] == BOX_RIGHT {
	// 	return true
	// }

	// // Halfstacking down-movement
	// if area[y][x] == BOX_LEFT && area[y-1][x+1] == BOX_LEFT ||
	// 	area[y][x] == BOX_RIGHT && area[y-1][x] == BOX_LEFT ||
	// 	area[y][x] == BOX_RIGHT && area[y-1][x-1] == BOX_RIGHT ||
	// 	area[y][x] == BOX_LEFT && area[y-1][x] == BOX_RIGHT {
	// 	return true
	// }

	// // Halfstacking up-movement
	// if area[y][x] == BOX_LEFT && area[y+1][x-1] == BOX_LEFT ||
	// 	area[y][x] == BOX_LEFT && area[y+1][x] == BOX_RIGHT ||
	// 	area[y][x] == BOX_RIGHT && area[y+1][x+1] == BOX_RIGHT ||
	// 	area[y][x] == BOX_RIGHT && area[y+1][x] == BOX_LEFT {
	// 	return true
	// }

	// Closing
	if area[y][x] == BOX_LEFT && area[y][x+1] == BOX_RIGHT ||
		area[y][x] == BOX_RIGHT && area[y][x-1] == BOX_LEFT {
		return true
	}
	return false
}

func (s *Solution) FloodFillTheYaxis(area [][]string, points *[][]int, stack *Stack[[]int], dir string, s_y int) {
	//fmt.Println("Stack length: ", len(stack.items), " contents", stack.items)
	if !stack.IsEmpty() {
		p := stack.Pop()
		//fmt.Println("PROBA 1", p, " ", area[p[0]][p[1]])

		if s.IsInBoxesGroup(area, p[0], p[1], s_y, dir) {
			if !slices.ContainsFunc((*points), func(checked []int) bool {
				if checked[0] == p[0] && checked[1] == p[1] {
					return true
				}
				return false
			}) {
				(*points) = append((*points), p)
			}
			//fmt.Println("PROBA 2", (*points))
		} else {
			return
		}

		for i := 0; i < 4; i++ {
			// if dir == C_TOP && i == I_DOWN {
			// 	continue
			// }
			// if dir == C_DOWN && i == I_TOP {
			// 	continue
			// }
			n_y := p[0] + movement_y[i]
			n_x := p[1] + movement_x[i]

			if s.IsPointSafe(area, n_y, n_x) {
				//fmt.Println("PROBA 3", n_y, " ", n_x)
				if area[n_y][n_x] == BOX_LEFT || area[n_y][n_x] == BOX_RIGHT {
					if !slices.ContainsFunc((*points), func(checked []int) bool {
						if checked[0] == n_y && checked[1] == n_x {
							return true
						}
						return false
					}) {
						stack.Push([]int{n_y, n_x})
					}
				}
			}
		}
		s.FloodFillTheYaxis(area, points, stack, dir, s_y)
	}
}

// Ze dreaded thing.
/*
Now seriously, I plan to:
- approach It with flood filling the area insearch of interconnected boxes,
- sort the points and pick them depending on the direction
going up (the lowest index) / going down (the highest index)
- check whether those chosen points don't have a wall in front of them.
- If not, increment every height by one
- Move the lazy robot's a... to the right spot.
*/
func (s *Solution) PushAxisYpart2(area [][]string, d string) [][]string {

	y, x := s.FindRobotPosition(area)

	switch d {
	case C_TOP:
		if area[y+movement_y[I_TOP]][x] == WALL {
			return area
		}

		//fmt.Println("TESTING @ZE IDEA")
		points := [][]int{}
		stack := Stack[[]int]{}
		stack.Push([]int{y + movement_y[I_TOP], x})
		s.FloodFillTheYaxis(area, &points, &stack, C_TOP, y)
		//fmt.Println("Len points", len(points))

		if len(points) == 0 {
			area[y][x] = BLANK
			area[y+movement_y[I_TOP]][x] = ROBOT
			return area
		}

		slices.SortFunc(points, func(f []int, s []int) int {
			if f[0] < s[0] {
				return -1
			}
			return 1
		})

		for _, p := range points {
			//fmt.Println("Point", "(", p[0], ",", p[1], ")", " ", area[p[0]][p[1]])

			if area[p[0]-1][p[1]] == WALL {
				// Wypierdalać
				//fmt.Println("Wypuerdalac")
				return area
			}
		}

		//Swapping the point according to the sequence in the array
		swapping := Stack[string]{}

		for _, p := range points {
			swapping.Push(area[p[0]-1][p[1]])
			area[p[0]-1][p[1]] = area[p[0]][p[1]]
			area[p[0]][p[1]] = swapping.Pop()
		}
		area[y-1][x] = ROBOT
		area[y][x] = BLANK

		// ///////////
		// if area[y+movement_y[I_TOP]][x] == WALL {
		// 	return area
		// }

		// fmt.Println("TESTING @ZE IDEA")
		// points := [][]int{}
		// stack := Stack[[]int]{}
		// stack.Push([]int{y + movement_y[I_TOP], x})
		// s.FloodFillTheYaxis(area, &points, &stack)
		// fmt.Println(points)

		// slices.SortFunc(points, func(f []int, s []int) int {
		// 	if f[0] > s[1] {
		// 		return 1
		// 	}
		// 	return -1
		// })

		// // Convert to a vertical slice. TOP-DOWN
		// vertical_line := ""
		// for _, line := range area {
		// 	vertical_line += line[x]
		// }

		// boxesMatch, _ := regexp.Compile(`[\[\]]+`)
		// boxes := boxesMatch.FindAllString(vertical_line, -1)
		// indexes := boxesMatch.FindAllStringIndex(vertical_line, -1)

		// // Filtering the indexes with garbage collection
		// filter := indexes[:0]
		// filtered_boxes := boxes[:0]

		// for i, candidate := range indexes {
		// 	if candidate[1] <= y {
		// 		filter = append(filter, candidate)
		// 		filtered_boxes = append(filtered_boxes, boxes[i])
		// 	}
		// }

		// for i := len(filter); i < len(indexes); i++ {
		// 	indexes[i] = nil // or the zero value of T
		// }

		// // If all of the boxes are on the opposite side, move and fajrant
		// if len(filter) == 0 {
		// 	area[y][x] = BLANK
		// 	area[y+movement_y[I_TOP]][x] = ROBOT
		// 	return area
		// }

		// //fmt.Println("TOPSTART! ", filter)
		// //fmt.Println(filtered_boxes)
		// //fmt.Println("Edge case", s.DataProvider.GetArea()[filter[len(filter)-1][0]-1][x])

		// /* How to piece together the flood fill alg with this logic>? */

		// // If the boxes are further than 1 from the robot, move the robot and fajrant
		// if filter[len(filter)-1][1] != y {
		// 	area[y][x] = BLANK
		// 	area[y+movement_y[I_TOP]][x] = ROBOT
		// 	//fmt.Println("ASDASDSDASD")
		// 	return area
		// } else if area[filter[len(filter)-1][0]-1][x] == WALL {
		// 	//fmt.Println("ASDASDASFEFWW#CWECWE")
		// 	return area
		// } else {
		// 	// Move the robot on the first found index
		// 	area[filter[len(filter)-1][1]-1][x] = ROBOT
		// 	area[filter[len(filter)-1][1]][x] = BLANK

		// 	//fmt.Println("ASDASDD")
		// 	// Shift the boxes by one to the right
		// 	for i := 0; i < len(filtered_boxes[len(filtered_boxes)-1]); i++ {
		// 		area[filter[len(filter)-1][0]-1+i][x] = BOX
		// 	}

		//}
	case C_DOWN:
		if area[y+movement_y[I_DOWN]][x] == WALL {
			return area
		}

		//fmt.Println("TESTING @ZE IDEA")
		points := [][]int{}
		stack := Stack[[]int]{}
		stack.Push([]int{y + movement_y[I_DOWN], x})
		s.FloodFillTheYaxis(area, &points, &stack, C_DOWN, y)
		//fmt.Println("Len points", len(points))

		if len(points) == 0 {
			area[y][x] = BLANK
			area[y+movement_y[I_DOWN]][x] = ROBOT
			return area
		}

		slices.SortFunc(points, func(f []int, s []int) int {
			if f[0] > s[0] {
				return -1
			}
			return 1
		})

		for _, p := range points {
			//fmt.Println("Point", "(", p[0], ",", p[1], ")", " ", area[p[0]][p[1]])

			if area[p[0]+1][p[1]] == WALL {
				// Wypierdalać
				//fmt.Println("Wypuerdalac")
				return area
			}
		}

		//Swapping the point according to the sequence in the array
		swapping := Stack[string]{}

		for _, p := range points {
			swapping.Push(area[p[0]+1][p[1]])
			area[p[0]+1][p[1]] = area[p[0]][p[1]]
			area[p[0]][p[1]] = swapping.Pop()
		}
		area[y+1][x] = ROBOT
		area[y][x] = BLANK

		// vertical_line := ""
		// for _, line := range area {
		// 	vertical_line += line[x]
		// }

		// boxesMatch, _ := regexp.Compile(`[\[\]]+`)
		// boxes := boxesMatch.FindAllString(vertical_line, -1)
		// indexes := boxesMatch.FindAllStringIndex(vertical_line, -1)

		// // Filtering the indexes with garbage collection
		// filter := indexes[:0]
		// filtered_boxes := boxes[:0]

		// for i, candidate := range indexes {
		// 	if candidate[0] > y {
		// 		filter = append(filter, candidate)
		// 		filtered_boxes = append(filtered_boxes, boxes[i])
		// 	}
		// }

		// for i := len(filter); i < len(indexes); i++ {
		// 	indexes[i] = nil // or the zero value of T
		// }

		// // If all of the boxes are on the opposite side, move and fajrant
		// if len(filter) == 0 {
		// 	area[y][x] = BLANK
		// 	area[y+movement_y[I_DOWN]][x] = ROBOT
		// 	return area
		// }

		// //fmt.Println("DOWNSTART! ", filter)
		// //fmt.Println(filtered_boxes)

		// // If the boxes are further than 1 from the robot, move the robot and fajrant
		// if math.Abs(float64(filter[0][0]-y)) != 1 {
		// 	area[y][x] = BLANK
		// 	area[y+movement_y[I_DOWN]][x] = ROBOT
		// 	//fmt.Println("ASDASDSDASD")
		// 	return area
		// } else if area[filter[0][1]][x] == WALL {
		// 	//fmt.Println("ASDASDSDASDsdfsadf")
		// 	return area
		// } else {
		// 	// Move the robot on the first found index
		// 	area[filter[0][0]][x] = ROBOT
		// 	area[filter[0][0]-1][x] = BLANK

		// 	//fmt.Println("ASDASDAAA")
		// 	// Shift the boxes by one to the right
		// 	for i := 0; i < len(filtered_boxes[0]); i++ {
		// 		area[filter[0][0]+1+i][x] = BOX
		// 	}

		// }
	}

	return area
}

func (s *Solution) PushAxisXpart2(area [][]string, d string) [][]string {
	/*
		I see It this way.
		- Seek the position of the robot
		- Seek the nearest wall to the specified direction
		- While seeking the wall, count boxes that can be shifted
		- Calculate the delta distance between the robot and the wall
		- delta will be equal to empty space + robot + boxes.
		- rearrange the room in the span of this delta to first have the boxes
		- then robot
		- then empty space
	*/

	y, x := s.FindRobotPosition(area)
	//fmt.Println(y, x, "EHASD")

	switch d {
	case C_RIGHT:
		if area[y][x+movement_x[I_RIGHT]] == WALL {
			//fmt.Println("ASDAGEDWW")
			return area
		}
		boxesMatch, _ := regexp.Compile(`[\[\]]+`)
		boxes := boxesMatch.FindAllString(strings.Join(area[y], ""), -1)
		indexes := boxesMatch.FindAllStringIndex(strings.Join(area[y], ""), -1)

		// Filtering the indexes with garbage collection
		filter := indexes[:0]
		filtered_boxes := boxes[:0]

		for i, candidate := range indexes {
			if candidate[0] > x {
				filter = append(filter, candidate)
				filtered_boxes = append(filtered_boxes, boxes[i])
			}
		}

		for i := len(filter); i < len(indexes); i++ {
			indexes[i] = nil // or the zero value of T
		}

		// If all of the boxes are on the opposite side, move and fajrant
		if len(filter) == 0 || len(filtered_boxes) == 0 {
			area[y][x] = BLANK
			area[y][x+movement_x[I_RIGHT]] = ROBOT
			//fmt.Println("HELO!")
			return area
		}

		//fmt.Println(filter)
		//fmt.Println(filtered_boxes)

		// If the boxes are further than 1 from the robot, move the robot and fajrant
		if math.Abs(float64(filter[0][0]-x)) != 1 {
			//fmt.Println("JELO")
			area[y][x] = BLANK
			area[y][x+movement_x[I_RIGHT]] = ROBOT
			return area
		}

		if area[y][filter[0][1]] == WALL {
			// If the index succeeding the found box(es) is a wall, get out
			return area
		}

		// Move the robot on the first found index
		area[y][indexes[0][0]] = ROBOT
		area[y][x] = BLANK

		// Shift the boxes by one to the right
		for i := 0; i < len(filtered_boxes[0]); i += 2 {
			area[y][filter[0][1]-i] = "]"
			area[y][filter[0][1]-1-i] = "["
		}

		// // That is wrong
		// for i := wall_x - 1; i >= x; i-- {
		// 	if boxes_count > 0 {
		// 		s.DataProvider.GetArea()[y][i] = BOX
		// 		boxes_count--
		// 	} else if boxes_count == 0 {
		// 		s.DataProvider.GetArea()[y][i] = ROBOT
		// 		boxes_count--
		// 	} else {
		// 		s.DataProvider.GetArea()[y][i] = BLANK
		// 	}
		// }
		//fmt.Println("line expected: ", s.DataProvider.GetArea()[y])
	case C_LEFT:
		if area[y][x+movement_x[I_LEFT]] == WALL {
			return area
		}

		boxesMatch, _ := regexp.Compile(`[\[\]]+`)
		boxes := boxesMatch.FindAllString(strings.Join(area[y], ""), -1)
		indexes := boxesMatch.FindAllStringIndex(strings.Join(area[y], ""), -1)

		// Filtering the indexes with garbage collection
		filter := indexes[:0]
		filtered_boxes := boxes[:0]

		for i, candidate := range indexes {
			if candidate[1] <= x {
				filter = append(filter, candidate)
				filtered_boxes = append(filtered_boxes, boxes[i])
			}
		}

		for i := len(filter); i < len(indexes); i++ {
			indexes[i] = nil // or the zero value of T
		}

		//fmt.Println(filter)
		//fmt.Println(filtered_boxes)

		// If all of the boxes are on the opposite side, move and fajrant
		if len(filter) == 0 || len(filtered_boxes) == 0 {
			//fmt.Println("AASDASDSDASDQWW")
			area[y][x] = BLANK
			area[y][x+movement_x[I_LEFT]] = ROBOT
			return area
		}

		if math.Abs(float64(filter[len(filter)-1][1]-1-x)) != 1 {
			area[y][x] = BLANK
			area[y][x+movement_x[I_LEFT]] = ROBOT
			return area
		}

		// If the index succeeding the found box(es) is a wall, get out
		if area[y][filter[len(filter)-1][0]-1] == WALL {
			//fmt.Println("ASDASDQWW")
			return area
		}

		// Move the robot on the first found index
		area[y][filter[len(filter)-1][1]-1] = ROBOT
		area[y][x] = BLANK

		//fmt.Println("asdasd131231231a")
		// Shift the boxes by one to the right
		for i := 0; i < len(filtered_boxes[len(filtered_boxes)-1]); i += 2 {
			area[y][filter[len(filter)-1][1]-2-i] = "]"
			area[y][filter[len(filter)-1][1]-3-i] = "["
		}

	}

	return area
}

func (s *Solution) Part2() int {
	area := s.ConvertPart1AreaToPart2Version()

	for _, line := range area {
		fmt.Println(line)
	}

	for _, command := range s.DataProvider.GetCommands() {
		// if i > 5 {
		// 	break
		// }
		//y, x := s.FindRobotPosition(s.DataProvider.GetArea())

		//fmt.Println(area)
		switch command {
		case C_TOP:
			//
			//fmt.Println("Command top")
			// if s.DataProvider.GetArea()[y+movement_y[I_TOP]][x] == WALL {
			// 	continue
			// }
			area = s.PushAxisYpart2(area, C_TOP)

		case C_RIGHT:
			//
			//fmt.Println("Command right")
			// if s.DataProvider.GetArea()[y][x+movement_x[I_RIGHT]] == WALL {
			// 	continue
			// }
			area = s.PushAxisXpart2(area, C_RIGHT)

		case C_DOWN:
			//
			//fmt.Println("Command down")
			// if s.DataProvider.GetArea()[y+movement_y[I_DOWN]][x] == WALL {
			// 	continue
			// }
			area = s.PushAxisYpart2(area, C_DOWN)

		case C_LEFT:
			//
			//fmt.Println("Command left")
			// if s.DataProvider.GetArea()[y][x+movement_x[I_LEFT]] == WALL {
			// 	continue
			// }
			area = s.PushAxisXpart2(area, C_LEFT)
		}
	}

	for _, line := range area {
		fmt.Println(line)
	}

	res := 0

	// Change for regex style
	for y, line := range area {
		boxesMatch, _ := regexp.Compile(`[\[\]]`)
		indexes := boxesMatch.FindAllStringIndex(strings.Join(line, ""), -1)
		for _, index := range indexes {
			res += y*100 + index[0]
		}
	}

	return res
}
