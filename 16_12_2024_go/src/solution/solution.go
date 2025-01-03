package solution

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"math/big"
	"time"

	interfaces "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/data"
)

var (
	BOUNDARY = "#"

	UP     = 0
	RIGHT1 = 1
	DOWN1  = 2
	LEFT1  = 3

	MOVE1 = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
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

type Point struct {
	y int
	x int
}

type Neighbour struct {
	Vertice   *Vertice
	Direction int
}

type Vertice struct {
	// Y, X
	Coordinates Point

	Direction int

	// As pointers to existing points
	Neighbours []Neighbour
}

type AStarHelper struct {
	// Non heap related data - business
	Vertice       *Vertice
	ParentVertice *AStarHelper
	G             uint64
	H             uint64
	Direction     int

	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueueAStar []*AStarHelper

func (pq PriorityQueueAStar) Len() int { return len(pq) }

func (pq PriorityQueueAStar) Less(i, j int) bool {
	// if (pq[i].G + pq[i].H) == (pq[j].G + pq[j].H) {
	// 	return pq[i].H < pq[j].H
	// }
	return (pq[i].G + pq[i].H) < (pq[j].G + pq[j].H)
}

func (pq PriorityQueueAStar) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueueAStar) Push(x any) {
	n := len(*pq)
	item := x.(*AStarHelper)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueueAStar) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueueAStar) Get(point [2]int) *AStarHelper {
	if pq.Len() == 0 {
		return nil
	}

	for i := 0; i < pq.Len(); i++ {
		if (*pq)[i].Vertice.Coordinates.y == point[0] &&
			(*pq)[i].Vertice.Coordinates.x == point[1] {
			return (*pq)[i]
		}
	}
	return nil
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueueAStar) update(helper *AStarHelper, y, x int, g uint64, h uint64) {
	helper.Vertice.Coordinates.y = y
	helper.Vertice.Coordinates.x = x
	helper.G = g
	helper.H = h
	heap.Fix(pq, helper.index)
}

func (pq *PriorityQueueAStar) delete(to_delete *AStarHelper) {
	// Saving the old index and swapping the wanted item with the last one
	old_index := to_delete.index
	pq.Swap(to_delete.index, pq.Len()-1)

	// Deleting the last item
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	item = nil
	heap.Fix(pq, old_index)
}

func MakeWalkableVerticesMap(grid [][]string) *map[[2]int]*Vertice {

	walkable_vertices := map[[2]int]*Vertice{}

	// Creating the graph representation out of 2D grid
	for y, points_row := range grid {
		for x, point := range points_row {

			// Walkable meaning, the blockades won't be added to the list
			if point != BOUNDARY {

				// Contents of the newly created and non-asociated(yet) vertice
				coords := Point{y: y, x: x}
				neighbours := []Neighbour{}

				// Appending to the *graph*
				walkable_vertices[[2]int{y, x}] = &Vertice{Coordinates: coords, Neighbours: neighbours}
			}
		}
	}

	return &walkable_vertices
}

func AssociateWithNeighbours(vertices *map[[2]int]*Vertice) *map[[2]int]*Vertice {

	// Iterating over the vertices and asociating each and one of them
	// With potential neighbours
	for point, vertice := range *vertices {

		/*
			Adding the UP/RIGHT/DOWN/LEFT neighbours to the vertices
			if possible
		*/
		for i, move := range MOVE1 {
			y := point[0] + move[0]
			x := point[1] + move[1]

			if neighbour := (*vertices)[[2]int{y, x}]; neighbour != nil {
				(*vertice).Neighbours = append((*vertice).Neighbours, Neighbour{Vertice: neighbour, Direction: i})
			}
		}
	}

	return vertices
}

func (s *Solution) AStarAlternate(vertices *map[[2]int]*Vertice, start, end [2]int) int {
	open := PriorityQueueAStar{}
	heap.Init(&open)

	StartVertice := (*vertices)[start]
	EndVertice := (*vertices)[end]
	if StartVertice == nil || EndVertice == nil {
		return 0
	}

	first := AStarHelper{Vertice: StartVertice, ParentVertice: nil, H: 0, G: 0, Direction: LEFT1}
	heap.Push(&open, &first)

	cost_so_far := map[[2]int]*AStarHelper{}
	cost_so_far[[2]int{StartVertice.Coordinates.y, StartVertice.Coordinates.x}] = &first
	// cost_so_far := map[[2]int]*AStarHelper{}
	// cost_so_far[[2]int{StartVertice.Coordinates.y, StartVertice.Coordinates.x}] = &first

	visited := map[[2]int]bool{}
	visited[[2]int{StartVertice.Coordinates.y, StartVertice.Coordinates.x}] = true
	// visited := map[[2]int]bool{}
	// visited[[2]int{StartVertice.Coordinates.y, StartVertice.Coordinates.x}] = true

	for open.Len() > 0 {
		checked := heap.Pop(&open).(*AStarHelper)

		fmt.Println("", checked.Vertice.Coordinates.y, " ", checked.Vertice.Coordinates.x, " ", checked.Direction, " ", int(checked.G))
		//time.Sleep(time.Second)

		// if checked.Vertice.Coordinates.y == end[0] &&
		// 	checked.Vertice.Coordinates.x == end[1] {
		// 	break
		// }

		// if visited[[3]int{checked.Vertice.Coordinates.y, checked.Vertice.Coordinates.x, checked.Direction}] {
		// 	continue
		// }

		// if visited[[2]int{checked.Vertice.Coordinates.y, checked.Vertice.Coordinates.x}] {
		// 	continue
		// }

		for _, neighbour := range checked.Vertice.Neighbours {
			inSearch := cost_so_far[[2]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x}]
			//inSearch := cost_so_far[[2]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x}]

			g := uint64(math.MaxInt)
			if inSearch != nil {
				g = inSearch.G
			}

			n_g := checked.G + s.CostOfMovement(checked.Direction, neighbour.Direction)

			// if neighbour_in_open := open.Get([2]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x}); neighbour_in_open != nil {
			// 	if neighbour_in_open.G > n_g {
			// 		open.update(neighbour_in_open, neighbour_in_open.Vertice.Coordinates.y, neighbour_in_open.Vertice.Coordinates.x, n_g, 0)
			// 	}
			// 	continue
			// }

			// if inSearch := cost_so_far[[2]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x}]; inSearch != nil {
			// 	if inSearch.G > n_g {
			// 		inSearch.G = n_g
			// 		open.Push(inSearch)
			// 		cost_so_far[[2]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x}] = nil
			// 	}
			// 	continue
			// }

			if inSearch == nil {
				n_h := 5 * s.Manhattan(neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x, end)
				first := AStarHelper{Vertice: neighbour.Vertice, ParentVertice: checked, H: n_h, G: n_g, Direction: neighbour.Direction}
				cost_so_far[[2]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x}] = &first
				visited[[2]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x}] = true
				//cost_so_far[[2]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x}] = &first
				heap.Push(&open, &first)
			} else if g != math.MaxInt {
				n_h := uint64(0)

				inSearch.ParentVertice = checked
				inSearch.G = n_g
				inSearch.H = n_h
				inSearch.Direction = neighbour.Direction

				cost_so_far[[2]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x}] = inSearch
				visited[[2]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x}] = true
				//cost_so_far[[2]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x}] = &first
				heap.Push(&open, inSearch)
			}
			// }
			// first := AStarHelper{Vertice: neighbour.Vertice, ParentVertice: checked.Vertice, H: 0, G: n_g, Direction: neighbour.Direction}
			// cost_so_far[[2]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x}] = &first
			// //cost_so_far[[2]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x}] = &first
			// heap.Push(&open, &first)
		}
	}

	// Path reconstruction
	helper := cost_so_far[[2]int{end[0], end[1]}]
	//helper := cost_so_far[[2]int{end[0], end[1]}]
	//count := 0
	score := helper.G

	// for helper != nil {

	// 	fmt.Println("", helper.Vertice.Coordinates.y, " ", helper.Vertice.Coordinates.x, " ", helper.Direction, " ", int(helper.G), " ", count)
	// 	if helper.ParentVertice != nil {
	// 		count += len(helper.ParentVertice)
	// 		helper = cost_so_far[[2]int{helper.ParentVertice[len(helper.ParentVertice)-1].Coordinates.y, helper.ParentVertice[len(helper.ParentVertice)-1].Coordinates.x}]
	// 	} else {
	// 		break
	// 	}
	// }

	return int(score)
}

// To be used only on success condition, otherwise It will loop forever. It doesn't have any reasonable safety.
func (s *Solution) ReconstructPath(vertex *AStarHelper, vertices map[[2]int]*Vertice, start, end [2]int) {

	if vertex.Vertice.Coordinates.y == end[0] && vertex.Vertice.Coordinates.x == end[1] {
		if vertices[[2]int{vertex.Vertice.Coordinates.y, vertex.Vertice.Coordinates.x}] == nil {
			vertices[[2]int{vertex.Vertice.Coordinates.y, vertex.Vertice.Coordinates.x}] = vertex.Vertice
		}
		return
	}

	if vertices[[2]int{vertex.Vertice.Coordinates.y, vertex.Vertice.Coordinates.x}] == nil {
		vertices[[2]int{vertex.Vertice.Coordinates.y, vertex.Vertice.Coordinates.x}] = vertex.Vertice
	}

	s.ReconstructPath(vertex.ParentVertice, vertices, start, end)
}

func (s *Solution) Dijkstra(vertices *map[[2]int]*Vertice, start, end [2]int) int {
	open := PriorityQueueAStar{}
	heap.Init(&open)

	StartVertice := (*vertices)[start]
	EndVertice := (*vertices)[end]
	if StartVertice == nil || EndVertice == nil {
		return 0
	}

	first := AStarHelper{Vertice: StartVertice, ParentVertice: nil, H: 0, G: 0, Direction: LEFT1}
	heap.Push(&open, &first)

	cost := map[[3]int]uint64{}
	cost[[3]int{StartVertice.Coordinates.y, StartVertice.Coordinates.x, LEFT1}] = 0

	visited := map[[4]int]bool{}
	visited[[4]int{StartVertice.Coordinates.y, StartVertice.Coordinates.x, LEFT1, 0}] = true

	for open.Len() > 0 {
		checked := heap.Pop(&open).(*AStarHelper)

		// If the end vertex has been found, returns the score
		if checked.Vertice.Coordinates.y == end[0] &&
			checked.Vertice.Coordinates.x == end[1] {
			return int(checked.G)
		}

		// Dijkstra neighbour checking
		for _, neighbour := range checked.Vertice.Neighbours {
			g := cost[[3]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x, neighbour.Direction}]

			n_g := checked.G + s.CostOfMovement(checked.Direction, neighbour.Direction)

			if g == 0 || n_g <= g && !visited[[4]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x, neighbour.Direction, int(n_g)}] {
				visited[[4]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x, neighbour.Direction, int(n_g)}] = true
				frontier_element := AStarHelper{Vertice: neighbour.Vertice, ParentVertice: checked, H: 0, G: n_g, Direction: neighbour.Direction}
				cost[[3]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x, neighbour.Direction}] = n_g
				heap.Push(&open, &frontier_element)
			}
		}
	}

	// Error occured
	return 0
}

func (s *Solution) DijkstraAllPaths(vertices *map[[2]int]*Vertice, start, end [2]int) (int, int) {
	open := PriorityQueueAStar{}
	heap.Init(&open)

	StartVertice := (*vertices)[start]
	EndVertice := (*vertices)[end]
	if StartVertice == nil || EndVertice == nil {
		return 0, 0
	}

	first := AStarHelper{Vertice: StartVertice, ParentVertice: nil, H: 0, G: 0, Direction: LEFT1}
	heap.Push(&open, &first)

	cost := map[[3]int]uint64{}
	cost[[3]int{StartVertice.Coordinates.y, StartVertice.Coordinates.x, LEFT1}] = 0

	best_path_result := math.MaxInt
	vert := map[[2]int]*Vertice{}

	visited := map[[3]int]bool{}
	visited[[3]int{StartVertice.Coordinates.y, StartVertice.Coordinates.x, LEFT1}] = true

	for open.Len() > 0 {
		checked := heap.Pop(&open).(*AStarHelper)

		// If the score gets bigger than the "optimal", the
		// area of our interest has been exhausted - hence returns
		if checked.G > uint64(best_path_result) {
			break
		}

		// Reconstructing the path on found target,
		// including the vertices to a set
		if checked.Vertice.Coordinates.y == end[0] &&
			checked.Vertice.Coordinates.x == end[1] {
			s.ReconstructPath(checked, vert, end, start)
			best_path_result = int(checked.G)
		}

		// Dijkstra
		for _, neighbour := range checked.Vertice.Neighbours {
			g := cost[[3]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x, neighbour.Direction}]

			n_g := checked.G + s.CostOfMovement(checked.Direction, neighbour.Direction)

			if g == 0 || n_g <= g && !visited[[3]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x, neighbour.Direction}] {
				//visited[[3]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x, neighbour.Direction}] = true
				frontier_element := AStarHelper{Vertice: neighbour.Vertice, ParentVertice: checked, H: 0, G: n_g, Direction: neighbour.Direction}
				cost[[3]int{neighbour.Vertice.Coordinates.y, neighbour.Vertice.Coordinates.x, neighbour.Direction}] = n_g
				heap.Push(&open, &frontier_element)
			}
		}
	}

	return len(vert), best_path_result
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

func (s *Solution) Chebyshev(y, x int, destination [2]int) uint64 {
	return uint64(
		math.Round(
			math.Max(
				math.Abs(float64(destination[0]-y)),
				math.Abs(float64(destination[1]-x)))))
}

func (s *Solution) Manhattan(y, x int, destination [2]int) uint64 {
	return 2 * uint64(math.Abs(float64(destination[0]-y))+math.Abs(float64(destination[1]-x)))
}

func (s *Solution) H(y, x int, end []int) uint64 {
	return uint64(math.Round(math.Sqrt(math.Exp2(float64(end[0]-y))+math.Exp2(float64(end[1]-x)))) * float64(10))
}

func (s *Solution) H2(y, x int, end [2]int) uint64 {
	return uint64(math.Abs(float64(end[0]-y)) + 10*math.Abs(float64(end[1]-x)))
}

func (s *Solution) CostOfMovement(oldDir, newDir int) uint64 {
	if oldDir == LEFT1 && newDir == RIGHT1 {
		return 1
	} else if oldDir == UP && newDir == DOWN1 {
		return 1
	} else if oldDir == RIGHT1 && newDir == LEFT1 {
		return 1
	} else if oldDir == DOWN1 && newDir == UP {
		return 1
	} else {
		diff := math.Round(math.Abs(float64(oldDir - newDir)))
		if diff > 0 && diff < 3 {
			return uint64(math.Round(math.Abs(float64(oldDir-newDir))))*1000 + 1
		} else if diff == 3 {
			return 1001
		}
	}

	// diff := math.Round(math.Abs(float64(oldDir - newDir)))
	// // if diff > 0 && diff < 3 {
	// // 	return uint64(math.Round(math.Abs(float64(oldDir-newDir))))*1000 + 1
	// // } else if diff == 3 {
	// // 	return 1001
	// // }

	// if diff != 0 {
	// 	return 1001
	// }

	return 1
}

func (s *Solution) Astar(start, end [2]int) int {

	visited, _ := s.PrepareHelpers()

	fmt.Println(s.DataProvider.GetStartPoint())
	fmt.Println(s.DataProvider.GetEndPoint())

	visited[start[0]][start[1]] = true

	open := PriorityQueue{}
	heap.Init(&open)

	start_item := &Item{priority: 0, g: uint64(0), h: 0, y: start[0], x: start[1], dir: LEFT, p_y: start[0], p_x: start[1]}
	heap.Push(&open, start_item)

	cost_so_far := make(map[Point](*Item))
	cost_so_far[Point{y: start[0], x: start[1]}] = &Item{priority: 0, g: uint64(0), h: 0, y: start[0], x: start[1], dir: LEFT, p_y: start[0], p_x: start[1]}
	// 99416 to troche high

	for open.Len() > 0 {
		checked := heap.Pop(&open).(*Item)
		//closed = append(closed, checked)

		if checked.y == end[0] && checked.x == end[1] {
			break
		}

		for i, move := range MOVE {
			neighbor := Point{y: checked.y + move[0], x: checked.x + move[1]}

			if !s.IsInBounds(neighbor.y, neighbor.x) ||
				s.DataProvider.GetArea()[neighbor.y][neighbor.x] == "#" {
				continue
			}

			inSearch := cost_so_far[neighbor]

			// If We are talking about the neighbour currently
			//inSearch := open.Get([]int{n_y, n_x})

			g := uint64(math.MaxInt)
			if inSearch != nil {
				g = inSearch.g
			}

			n_g := checked.g + s.CostOfMovement(checked.dir, i)
			if inSearch == nil || (n_g < g && g != math.MaxInt) {
				n_h := s.H2(neighbor.y, neighbor.x, end)

				item := Item{priority: n_g + n_h, y: neighbor.y, x: neighbor.x, g: n_g, h: n_h, dir: i, p_y: checked.y, p_x: checked.x}

				cost_so_far[neighbor] = &item
				//fmt.Println("New ", n_h, " NIUEJCZ ", n_g, " ENDZI ", n_y, " ", n_x)
				heap.Push(&open, &item)
				visited[neighbor.y][neighbor.x] = true
			}
		}
	}

	// Path reconstruction
	endItem := cost_so_far[Point{y: end[0], x: end[1]}]

	result := uint64(0)
	if endItem != nil {
		result = endItem.g
	}

	// for endItem != nil {
	// 	for i := 0; i < len(closed); i++ {
	// 		if closed[i].y == endItem.p_y && closed[i].x == endItem.p_x {
	// 			if closed[i].y == start[0] && closed[i].x == start[1] {
	// 				endItem = closed[i]
	// 				fmt.Println(endItem)
	// 				endItem = nil
	// 				break
	// 			}
	// 			fmt.Println(endItem)
	// 			endItem = closed[i]
	// 			break

	// 		}
	// 	}
	// }

	return int(result)
}

/*
My idea for this solutioin would be to use
A* algorithm for least cost path finding.
h(x) as Chebyshev distance
f(x) as maybe the current move cost, ie. Front-Right-Down-Left
No need for path reconstruction if used recursively(?)
*/
// func (s *Solution) Part1() int {

//		return s.Astar(s.DataProvider.GetStartPoint(), s.DataProvider.GetEndPoint())
//	}

/*
Both of those colleagues have been solved with one dijkstra custom
It searches for all paths with "best" cost and exists if all
options are exhausted
*/

func (s *Solution) Part1() int {

	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	vertices := MakeWalkableVerticesMap(s.DataProvider.GetArea())

	vertices = AssociateWithNeighbours(vertices)

	// Pseudo memo
	result := s.Dijkstra(vertices, s.DataProvider.GetStartPoint(), s.DataProvider.GetEndPoint())

	elapsed := time.Since(start)
	log.Printf("Calculations for Part1 took %s", elapsed)

	return result
}

func (s *Solution) Part2() int {

	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	vertices := MakeWalkableVerticesMap(s.DataProvider.GetArea())

	vertices = AssociateWithNeighbours(vertices)

	result, _ := s.DijkstraAllPaths(vertices, s.DataProvider.GetStartPoint(), s.DataProvider.GetEndPoint())

	elapsed := time.Since(start)
	log.Printf("Calculations for Part2 took %s", elapsed)

	return result
}
