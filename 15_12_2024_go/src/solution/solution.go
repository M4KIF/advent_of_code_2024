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

	movement_y = []int{1, 1, 1, -1, -1, -1}
	movement_x = []int{1, -1, 0, 0, 1, -1}
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

// func (s *Solution) SearchForWallAndBoxesAtY(y, x int, direction string) (int, int) {
// 	// Oribl
// 	if !s.IsPointSafe(s.DataProvider.GetArea(), y, x) {
// 		return -1, -1
// 	}

// 	boxes_count := 0

// 	wallMatch, _ := regexp.Compile("#")
// 	boxesMatch, _ := regexp.Compile("O")

// 	switch direction {
// 	case C_TOP:
// 		for t_y := y; t_y >= 0; t_y-- {
// 			if boxesMatch.Match([]byte(s.DataProvider.GetArea()[t_y][x])) {
// 				boxes_count++
// 			}
// 			if wallMatch.Match([]byte(s.DataProvider.GetArea()[t_y][x])) {
// 				return t_y, boxes_count
// 			}
// 		}
// 	case C_DOWN:
// 		for t_y := y; t_y < len(s.DataProvider.GetArea()); t_y++ {
// 			if boxesMatch.Match([]byte(s.DataProvider.GetArea()[t_y][x])) {
// 				boxes_count++
// 			}
// 			if wallMatch.Match([]byte(s.DataProvider.GetArea()[t_y][x])) {
// 				return t_y, boxes_count
// 			}
// 		}
// 	}

// 	return -1, -1
// }

func (s *Solution) MovementYAxisStraight(d string) bool {

	y, x := s.FindRobotPosition(s.DataProvider.GetArea())

	switch d {
	case C_TOP:
		if s.DataProvider.GetArea()[y-1][x] == WALL {
			return false
		}

		// Convert to a vertical slice. TOP-DOWN
		vertical_line := ""
		for _, line := range s.DataProvider.GetArea() {
			vertical_line += line[x]
		}

		boxesMatch, _ := regexp.Compile("O+")
		boxes := boxesMatch.FindAllString(vertical_line, -1)
		indexes := boxesMatch.FindAllStringIndex(vertical_line, -1)

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
			s.DataProvider.GetArea()[y-1][x] = ROBOT
			return true
		}

		// If the boxes are further than 1 from the robot, move the robot and fajrant
		if filter[len(filter)-1][1] != y {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y-1][x] = ROBOT
			return true
		} else if s.DataProvider.GetArea()[filter[len(filter)-1][0]-1][x] == WALL {
			return false
		} else {
			// Move the robot on the first found index
			s.DataProvider.GetArea()[filter[len(filter)-1][1]-1][x] = ROBOT
			s.DataProvider.GetArea()[filter[len(filter)-1][1]][x] = BLANK

			// Shift the boxes by one to the right
			for i := 0; i < len(filtered_boxes[len(filtered_boxes)-1]); i++ {
				s.DataProvider.GetArea()[filter[len(filter)-1][0]-1+i][x] = BOX
			}
		}
	case C_DOWN:
		if s.DataProvider.GetArea()[y+1][x] == WALL {
			return false
		}
		vertical_line := ""
		for _, line := range s.DataProvider.GetArea() {
			vertical_line += line[x]
		}

		boxesMatch, _ := regexp.Compile("O+")
		boxes := boxesMatch.FindAllString(vertical_line, -1)
		indexes := boxesMatch.FindAllStringIndex(vertical_line, -1)

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
			s.DataProvider.GetArea()[y+1][x] = ROBOT
			return true
		}

		// If the boxes are further than 1 from the robot, move the robot and fajrant
		if math.Abs(float64(filter[0][0]-y)) != 1 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y+1][x] = ROBOT
			return true
		} else if s.DataProvider.GetArea()[filter[0][1]][x] == WALL {
			return false
		} else {
			// Move the robot on the first found index
			s.DataProvider.GetArea()[filter[0][0]][x] = ROBOT
			s.DataProvider.GetArea()[filter[0][0]-1][x] = BLANK

			// Shift the boxes by one to the right
			for i := 0; i < len(filtered_boxes[0]); i++ {
				s.DataProvider.GetArea()[filter[0][0]+1+i][x] = BOX
			}
		}
	}

	return true
}

// func (s *Solution) SearchForWallAndBoxesAtX(y, x int, direction string) (int, int) {
// 	// Oribl
// 	if !s.IsPointSafe(s.DataProvider.GetArea(), y, x) {
// 		return -1, -1
// 	}

// 	boxes_count := 0

// 	wallMatch, _ := regexp.Compile("#")
// 	boxesMatch, _ := regexp.Compile("O")

// 	switch direction {
// 	case C_RIGHT:
// 		for t_x := x; t_x < len(s.DataProvider.GetArea()[y]); t_x++ {
// 			if boxesMatch.Match([]byte(s.DataProvider.GetArea()[y][t_x])) {
// 				boxes_count++
// 			}
// 			if wallMatch.Match([]byte(s.DataProvider.GetArea()[y][t_x])) {
// 				return t_x, boxes_count
// 			}
// 		}
// 	case C_LEFT:
// 		for t_x := x; t_x >= 0; t_x-- {
// 			if boxesMatch.Match([]byte(s.DataProvider.GetArea()[y][t_x])) {
// 				boxes_count++
// 			}
// 			if wallMatch.Match([]byte(s.DataProvider.GetArea()[y][t_x])) {
// 				return t_x, boxes_count
// 			}
// 		}
// 	}

// 	return -1, -1
// }

func (s *Solution) MovementXAxisStraight(d string) bool {
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
		if s.DataProvider.GetArea()[y][x+1] == WALL {
			return false
		}
		boxesMatch, _ := regexp.Compile("O+")
		boxes := boxesMatch.FindAllString(strings.Join(s.DataProvider.GetArea()[y], ""), -1)
		indexes := boxesMatch.FindAllStringIndex(strings.Join(s.DataProvider.GetArea()[y], ""), -1)

		if len(boxes) == 0 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y][x+1] = ROBOT
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
			s.DataProvider.GetArea()[y][x+1] = ROBOT
			return true
		}

		// If the boxes are further than 1 from the robot, move the robot and fajrant
		if math.Abs(float64(filter[0][0]-x)) != 1 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y][x+1] = ROBOT
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
	case C_LEFT:
		if s.DataProvider.GetArea()[y][x-1] == WALL {
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

		// If all of the boxes are on the opposite side, move and fajrant
		if len(filter) == 0 || len(filtered_boxes) == 0 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y][x-1] = ROBOT
			return true
		}

		if math.Abs(float64(filter[len(filter)-1][1]-1-x)) != 1 {
			s.DataProvider.GetArea()[y][x] = BLANK
			s.DataProvider.GetArea()[y][x-1] = ROBOT
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

		// Shift the boxes by one to the right
		for i := 0; i < len(filtered_boxes[len(filtered_boxes)-1]); i++ {
			s.DataProvider.GetArea()[y][filter[len(filter)-1][1]-2-i] = BOX
		}
	}

	return true
}

func (s *Solution) Part1() int {
	for _, command := range s.DataProvider.GetCommands() {
		switch command {
		case C_TOP:
			s.MovementYAxisStraight(C_TOP)
		case C_RIGHT:
			s.MovementXAxisStraight(C_RIGHT)
		case C_DOWN:
			s.MovementYAxisStraight(C_DOWN)
		case C_LEFT:
			s.MovementXAxisStraight(C_LEFT)
		}
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
	// That is the drawback of generics,
	// safety check in code, not in DS
	point := s.items[0]
	s.items = s.items[1:]
	return point
}

func (s *Solution) ContainsPoint(points [][]int, y, x int) bool {
	return slices.ContainsFunc(points, func(checked []int) bool {
		if checked[0] == y && checked[1] == x {
			return true
		}
		return false
	})
}

func (s *Solution) IsInBoxesGroup(area [][]string, points [][]int, y, x, starting_y, starting_x int, dir string) bool {
	if dir == C_DOWN {

		// Init step
		if area[y][x] == BOX_LEFT && area[y-1][x] == ROBOT ||
			area[y][x] == BOX_RIGHT && area[y-1][x] == ROBOT ||
			area[y][x] == BOX_LEFT && area[y-1][x+1] == ROBOT ||
			area[y][x] == BOX_RIGHT && area[y-1][x-1] == ROBOT {
			return true
		}

		// Straight up-down
		if area[y][x] == BOX_LEFT && area[y-1][x] == BOX_LEFT && s.ContainsPoint(points, y-1, x) ||
			area[y][x] == BOX_RIGHT && area[y-1][x] == BOX_RIGHT && s.ContainsPoint(points, y-1, x) {
			return true
		}

		// Halfstacking down-movement
		if area[y][x] == BOX_LEFT && area[y-1][x-1] == BOX_LEFT && s.ContainsPoint(points, y-1, x-1) ||
			area[y][x] == BOX_LEFT && area[y-1][x+1] == BOX_LEFT && s.ContainsPoint(points, y-1, x+1) ||
			area[y][x] == BOX_RIGHT && area[y-1][x+1] == BOX_RIGHT && s.ContainsPoint(points, y-1, x+1) ||
			area[y][x] == BOX_RIGHT && area[y-1][x-1] == BOX_RIGHT && s.ContainsPoint(points, y-1, x-1) ||
			area[y][x] == BOX_RIGHT && area[y-1][x] == BOX_LEFT && s.ContainsPoint(points, y-1, x) ||
			area[y][x] == BOX_LEFT && area[y-1][x] == BOX_RIGHT && s.ContainsPoint(points, y-1, x) {
			return true
		}
	}

	if dir == C_TOP {

		// Init step
		if area[y][x] == BOX_LEFT && area[y+1][x] == ROBOT ||
			area[y][x] == BOX_RIGHT && area[y+1][x] == ROBOT ||
			area[y][x] == BOX_LEFT && area[y+1][x+1] == ROBOT ||
			area[y][x] == BOX_RIGHT && area[y+1][x-1] == ROBOT {
			return true
		}

		// Straight up-down
		if area[y][x] == BOX_LEFT && area[y+1][x] == BOX_LEFT && s.ContainsPoint(points, y+1, x) ||
			area[y][x] == BOX_RIGHT && area[y+1][x] == BOX_RIGHT && s.ContainsPoint(points, y+1, x) {
			return true
		}

		// Halfstacking top-movement
		if area[y][x] == BOX_LEFT && area[y+1][x-1] == BOX_LEFT && s.ContainsPoint(points, y+1, x-1) ||
			area[y][x] == BOX_LEFT && area[y+1][x+1] == BOX_LEFT && s.ContainsPoint(points, y+1, x+1) ||
			area[y][x] == BOX_RIGHT && area[y+1][x+1] == BOX_RIGHT && s.ContainsPoint(points, y+1, x+1) ||
			area[y][x] == BOX_RIGHT && area[y+1][x-1] == BOX_RIGHT && s.ContainsPoint(points, y+1, x-1) ||
			area[y][x] == BOX_RIGHT && area[y+1][x] == BOX_LEFT && s.ContainsPoint(points, y+1, x) ||
			area[y][x] == BOX_LEFT && area[y+1][x] == BOX_RIGHT && s.ContainsPoint(points, y+1, x) {
			return true
		}

	}

	return false
}

func (s *Solution) IsNextStepEligible(area [][]string, points [][]int, y, x, n_y, n_x int) bool {
	// Base case
	if area[n_y][n_x] != BOX_LEFT && area[n_y][n_x] != BOX_RIGHT {
		return false
	}

	// if (area[n_y][n_x] == BOX_RIGHT && area[y][x] == BOX_RIGHT && n_x != x) ||
	// 	area[n_y][n_x] == BOX_LEFT && area[y][x] == BOX_LEFT && n_x != x {
	// 	return false
	// }

	return true
}

func (s *Solution) FloodFillTheYaxis(area [][]string, points *[][]int, stack *Stack[[]int], dir string, s_y int, s_x int) {
	fmt.Println("Stack length: ", len(stack.items), " contents", stack.items, " dir: ", dir)
	if !stack.IsEmpty() {
		p := stack.Pop()
		//fmt.Println("PROBA 1", p, " ", area[p[0]][p[1]])

		if s.IsInBoxesGroup(area, (*points), p[0], p[1], s_y, s_x, dir) {
			if !slices.ContainsFunc((*points), func(checked []int) bool {
				if checked[0] == p[0] && checked[1] == p[1] {
					return true
				}
				return false
			}) {
				(*points) = append((*points), p)
			}
			//fmt.Println("PROBA 2", (*points))
		}

		begin_i := 0
		max_i := 0
		if dir == C_DOWN {
			max_i = 2
		}
		if dir == C_TOP {
			begin_i = 3
			max_i = 5
		}

		for i := begin_i; i <= max_i; i++ {
			// if dir == C_TOP && i == I_DOWN {
			// 	continue
			// }
			// if dir == C_DOWN && i == I_TOP {
			// 	continue
			// }
			n_y := p[0] + movement_y[i]
			n_x := p[1] + movement_x[i]

			if s.IsPointSafe(area, n_y, n_x) {
				fmt.Println("PROBA 3", n_y, " ", n_x)
				if s.IsNextStepEligible(area, (*points), p[0], p[1], n_y, n_x) {
					if !slices.ContainsFunc(stack.items, func(checked []int) bool {
						if checked[0] == n_y && checked[1] == n_x {
							fmt.Println("ASDAWEHEJAHO")
							return true
						}
						return false
					}) {
						stack.Push([]int{n_y, n_x})
					}
				}
			}
		}
		s.FloodFillTheYaxis(area, points, stack, dir, s_y, s_x)
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
		if area[y-1][x] == WALL {
			return area
		}
		if area[y-1][x] != BOX_LEFT && area[y-1][x] != BOX_RIGHT {
			area[y][x] = BLANK
			area[y-1][x] = ROBOT
			return area
		}

		//fmt.Println("TESTING @ZE IDEA")
		points := [][]int{}
		stack := Stack[[]int]{}
		stack.Push([]int{y - 1, x})

		if s.IsPointSafe(area, y-1, x+1) {
			if area[y-1][x] == BOX_LEFT {
				stack.Push([]int{y - 1, x + 1})
			}
		}
		if s.IsPointSafe(area, y-1, x-1) {
			if area[y-1][x] == BOX_RIGHT {
				stack.Push([]int{y - 1, x - 1})
			}
		}

		s.FloodFillTheYaxis(area, &points, &stack, C_TOP, y-1, x)
		//fmt.Println("Len points", len(points))

		if len(points) == 0 {
			if area[y-1][x] != WALL {
				area[y][x] = BLANK
				area[y-1][x] = ROBOT
			}
			return area
		}

		slices.SortFunc(points, func(f []int, s []int) int {
			if f[0] < s[0] {
				return -1
			}
			return 1
		})

		for _, p := range points {
			fmt.Println("Point", "(", p[0], ",", p[1], ")", " ", area[p[0]][p[1]])

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
		// if area[y-1][x] == WALL {
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
		// 	area[y-1][x] = ROBOT
		// 	return area
		// }

		// //fmt.Println("TOPSTART! ", filter)
		// //fmt.Println(filtered_boxes)
		// //fmt.Println("Edge case", s.DataProvider.GetArea()[filter[len(filter)-1][0]-1][x])

		// /* How to piece together the flood fill alg with this logic>? */

		// // If the boxes are further than 1 from the robot, move the robot and fajrant
		// if filter[len(filter)-1][1] != y {
		// 	area[y][x] = BLANK
		// 	area[y-1][x] = ROBOT
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
		if area[y+1][x] == WALL {
			return area
		}
		if area[y+1][x] != BOX_LEFT && area[y+1][x] != BOX_RIGHT {
			area[y][x] = BLANK
			area[y+1][x] = ROBOT
			return area
		}

		//fmt.Println("TESTING @ZE IDEA")
		points := [][]int{}
		stack := Stack[[]int]{}
		stack.Push([]int{y + 1, x})

		if s.IsPointSafe(area, y+1, x+1) {
			if area[y+1][x] == BOX_LEFT {
				stack.Push([]int{y + 1, x + 1})
			}
		}
		if s.IsPointSafe(area, y+1, x-1) {
			if area[y+1][x] == BOX_RIGHT {
				stack.Push([]int{y + 1, x - 1})
			}
		}

		s.FloodFillTheYaxis(area, &points, &stack, C_DOWN, y, x)
		//fmt.Println("Len points", len(points))

		if len(points) == 0 {
			if area[y+1][x] != WALL {
				area[y][x] = BLANK
				area[y+1][x] = ROBOT
			}
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
		// 	area[y+1][x] = ROBOT
		// 	return area
		// }

		// //fmt.Println("DOWNSTART! ", filter)
		// //fmt.Println(filtered_boxes)

		// // If the boxes are further than 1 from the robot, move the robot and fajrant
		// if math.Abs(float64(filter[0][0]-y)) != 1 {
		// 	area[y][x] = BLANK
		// 	area[y+1][x] = ROBOT
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
	fmt.Println("Where is ze robot ", " coords: ", y, x)
	//fmt.Println(y, x, "EHASD")

	switch d {
	case C_RIGHT:
		if area[y][x+1] == WALL {
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
			area[y][x+1] = ROBOT
			//fmt.Println("HELO!")
			return area
		}

		// If the boxes are further than 1 from the robot, move the robot and fajrant
		if math.Abs(float64(filter[0][0]-x)) != 1 {
			//fmt.Println("JELO")
			area[y][x] = BLANK
			area[y][x+1] = ROBOT
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
	case C_LEFT:
		if area[y][x-1] == WALL {
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

		// If all of the boxes are on the opposite side, move and fajrant
		if len(filter) == 0 || len(filtered_boxes) == 0 {
			//fmt.Println("AASDASDSDASDQWW")
			area[y][x] = BLANK
			area[y][x-1] = ROBOT
			return area
		}

		if math.Abs(float64(filter[len(filter)-1][1]-1-x)) != 1 {
			area[y][x] = BLANK
			area[y][x-1] = ROBOT
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
		for _, line := range area {
			fmt.Println(line)
		}
		// if i > 5 {
		// 	break
		// }
		//y, x := s.FindRobotPosition(s.DataProvider.GetArea())

		//fmt.Println(area)
		switch command {
		case C_TOP:
			//
			fmt.Println("Command top")
			// if s.DataProvider.GetArea()[y-1][x] == WALL {
			// 	continue
			// }
			area = s.PushAxisYpart2(area, C_TOP)

		case C_RIGHT:
			//
			fmt.Println("Command right")
			// if s.DataProvider.GetArea()[y][x+1] == WALL {
			// 	continue
			// }
			area = s.PushAxisXpart2(area, C_RIGHT)

		case C_DOWN:
			//
			fmt.Println("Command down")
			// if s.DataProvider.GetArea()[y+1][x] == WALL {
			// 	continue
			// }
			area = s.PushAxisYpart2(area, C_DOWN)

		case C_LEFT:
			//
			fmt.Println("Command left")
			// if s.DataProvider.GetArea()[y][x-1] == WALL {
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
		boxesMatch, _ := regexp.Compile(`[\[][\]]`)
		indexes := boxesMatch.FindAllStringIndex(strings.Join(line, ""), -1)
		for _, index := range indexes {
			fmt.Println(index)
			res += y*100 + index[0]
		}
		fmt.Println("PRzerwa")
	}

	return res
}
