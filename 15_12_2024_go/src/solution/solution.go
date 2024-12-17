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

func (s *Solution) IsPointSafe(y, x int) bool {
	return (y >= 0 && y < len(s.DataProvider.GetArea()) &&
		x >= 0 && x < len(s.DataProvider.GetArea()[0]))
}

func (s *Solution) FindRobotPosition() (int, int) {
	atMatch, _ := regexp.Compile("[@]{1}")

	for y, line := range s.DataProvider.GetArea() {
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
	if !s.IsPointSafe(y, x) {
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

	y, x := s.FindRobotPosition()

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

	y, x := s.FindRobotPosition()

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
	if !s.IsPointSafe(y, x) {
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
	if !s.IsPointSafe(y, x) {
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

	y, x := s.FindRobotPosition()

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

	y, x := s.FindRobotPosition()

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
		//y, x := s.FindRobotPosition()
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

func (s *Solution) Part2() int {
	return 1
}
