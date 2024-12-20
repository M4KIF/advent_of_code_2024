package solution

import (
	"container/heap"
	"fmt"
	"math"

	interfaces "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/data"
)

type Solution struct {
	DataProvider interfaces.StringArea2DStartEnd
	Path         string
}

var (
	FRONT = 1
	RIGHT = 2
	DOWN  = 3
	LEFT  = 0

	MOVE = [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
)

func NewSolution(dataProvider interfaces.StringArea2DStartEnd, path string) *Solution {
	sol := &Solution{DataProvider: dataProvider, Path: path}
	sol.DataProvider.TakeInput(path)
	return sol
}

// An Item is something we manage in a priority queue.
type Item struct {
	priority uint64 // The priority of the item in the queue.

	g uint64
	h uint64

	y int
	x int

	dir int

	p_y int
	p_x int

	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].priority == pq[j].priority {
		return pq[i].h < pq[j].h
	}
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Get(point []int) *Item {
	if pq.Len() == 0 {
		return nil
	}

	for i := 0; i < pq.Len(); i++ {
		if (*pq)[i].y == point[0] && (*pq)[i].x == point[1] {
			return (*pq)[i]
		}
	}
	return nil
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, y, x int, priority uint64) {
	item.y = y
	item.x = x
	item.priority = priority
	heap.Fix(pq, item.index)
}

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
	point := s.items[0]
	s.items = s.items[1:]
	return point
}

/*
PART 1
*/

// Preparing the visited array, and the f(x), g(x), h(x), parrent_y, parrent_x
func (s *Solution) PrepareHelpers() ([][]bool, [][][]int) {
	visited := [][]bool{}
	a_star_helper := [][][]int{}

	for _, line := range s.DataProvider.GetArea() {
		temp_visited := []bool{}
		temp_a_star_helper := [][]int{}
		for i := 0; i < len(line); i++ {
			temp_visited = append(temp_visited, false)
			temp_a_star_helper = append(temp_a_star_helper, []int{math.MaxInt, 0, 0, 0, 0, 0})
		}
		visited = append(visited, temp_visited)
		a_star_helper = append(a_star_helper, temp_a_star_helper)
	}
	return visited, a_star_helper
}

func (s *Solution) IsInBounds(y, x int) bool {
	return y >= 0 && y < len(s.DataProvider.GetArea()) &&
		x >= 0 && x < len(s.DataProvider.GetArea()[0])
}

func (s *Solution) Chebyshev(y, x int, destination []int) uint64 {
	return uint64(
		math.Round(
			math.Max(
				math.Abs(float64(destination[0]-y)),
				math.Abs(float64(destination[1]-x)))))
}

func (s *Solution) H(y, x int, end []int) uint64 {
	return uint64(math.Round(math.Sqrt(math.Exp2(float64(end[0]-y))+math.Exp2(float64(end[1]-x)))) * float64(100))
}

func (s *Solution) H2(y, x int, end []int) uint64 {
	return uint64(2 * math.Round(math.Abs(float64(end[0]-y))+math.Abs(float64(end[1]-x))))
}

func (s *Solution) CostOfMovement(oldDir, newDir int) uint64 {
	if math.Round(math.Abs(float64(oldDir-newDir))) != 0 {
		return uint64(math.Round(math.Abs(float64(oldDir-newDir))))*1000 + 1
	}
	return 1
}

func (s *Solution) Astar(start, end []int) int {

	visited, _ := s.PrepareHelpers()

	fmt.Println(s.DataProvider.GetStartPoint())
	fmt.Println(s.DataProvider.GetEndPoint())

	visited[start[0]][start[1]] = true

	open := PriorityQueue{}
	heap.Init(&open)

	closed := []*Item{}

	heap.Push(&open, &Item{priority: 0, g: uint64(0), h: 0, y: start[0], x: start[1], dir: LEFT, p_y: start[0], p_x: start[1]})

	// 99416 to troche high

	for open.Len() > 0 {
		checked := heap.Pop(&open).(*Item)
		closed = append(closed, checked)

		if checked.y == end[0] && checked.x == end[1] {
			break
		}

		for i, move := range MOVE {
			n_y := checked.y + move[0]
			n_x := checked.x + move[1]

			if !s.IsInBounds(n_y, n_x) ||
				s.DataProvider.GetArea()[n_y][n_x] == "#" {
				continue
			}

			// If We are talking about the neighbour currently
			inSearch := open.Get([]int{n_y, n_x})

			n_g := checked.g + s.CostOfMovement(checked.dir, i)
			n_h := s.Chebyshev(n_y, n_x, end)
			//
			if inSearch == nil {
				//fmt.Println("New ", n_h, " NIUEJCZ ", n_g, " ENDZI ", n_y, " ", n_x)
				newItem := Item{priority: n_g + n_h, y: n_y, x: n_x, g: n_g, h: n_h, dir: i, p_y: checked.y, p_x: checked.x}
				heap.Push(&open, &newItem)
				visited[n_y][n_x] = true
			} else {
				if n_g < checked.g {
					fmt.Println("Old ", n_h, " NIUEJCZ ", n_g, " ENDZI ", n_y, " ", n_x)
					inSearch.priority = n_g + n_h
					inSearch.g = n_g
					inSearch.h = n_h
					inSearch.p_y = checked.y
					inSearch.p_x = checked.x
					inSearch.dir = i
					heap.Fix(&open, inSearch.index)
				}
			}
		}
	}

	// Path reconstruction
	endItem := closed[len(closed)-1]
	result := endItem.g

	for endItem != nil {
		for i := 0; i < len(closed); i++ {
			if closed[i].y == endItem.p_y && closed[i].x == endItem.p_x {
				if closed[i].y == start[0] && closed[i].x == start[1] {
					endItem = closed[i]
					fmt.Println(endItem)
					endItem = nil
					break
				}
				fmt.Println(endItem)
				endItem = closed[i]
				break

			}
		}
	}

	return int(result)
}

/*
My idea for this solutioin would be to use
A* algorithm for least cost path finding.
h(x) as Chebyshev distance
f(x) as maybe the current move cost, ie. Front-Right-Down-Left
No need for path reconstruction if used recursively(?)
*/
func (s *Solution) Part1() int {

	return s.Astar(s.DataProvider.GetStartPoint(), s.DataProvider.GetEndPoint())
}

/*
PART 2
*/

func (s *Solution) Part2() int {

	return 1
}
