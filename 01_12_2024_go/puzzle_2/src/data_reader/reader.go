package data_reader

import (
	"bufio"
	"strconv"
	"strings"

	icrud "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/file"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/logging"
)

type PuzzleTwoData struct {
	LeftArray  []int
	RightArray []int
	IO         icrud.CRUD
}

func Create(io icrud.CRUD) *PuzzleTwoData {
	return &PuzzleTwoData{IO: io}
}

func (d *PuzzleTwoData) TakeInput(path string) bool {
	// Creating the absolute path and opening the file
	file, e := d.IO.Open(path)

	if e != nil {
		logging.Error("Error occured while opening the file", "errors", e.Error())
		return false
	}

	// Async file close
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		// Dividing the line in order to extract the left and right entry
		l, r, f := strings.Cut(scanner.Text(), " ")
		if !f {
			continue
		}

		// Getting rid of whitespaces
		l = strings.ReplaceAll(l, " ", "")
		r = strings.ReplaceAll(r, " ", "")

		// Converting to integers
		left, e := strconv.Atoi(l)
		if e != nil {
			logging.Error("Error occured on converting the left string to int", "errors", e.Error())
			return false
		}

		right, e := strconv.Atoi(r)
		if e != nil {
			logging.Error("Error occured on converting the right string to int", "errors", e.Error())
			return false
		}

		// // Utilized @github mnml code for splitting the read line
		// var n1, n2 int
		// fmt.Sscanf(scanner.Text(), "%d   %d", &n1, &n2)
		// // Appending the Arrays
		// d.LeftArray = append(d.LeftArray, n1)
		// d.RightArray = append(d.RightArray, n2)

		// Appending the Arrays
		d.LeftArray = append(d.LeftArray, left)
		d.RightArray = append(d.RightArray, right)
	}

	if e := scanner.Err(); e != nil {
		logging.Error("Errors occured while scanning", "errors", e.Error())
		return false
	}

	logging.Info("Successfully saved the input to PuzzleTwoData{}", "path", path)
	return true
}

func (d *PuzzleTwoData) GetFirstArray() []int {
	return d.LeftArray
}

func (d *PuzzleTwoData) GetSecondArray() []int {
	return d.RightArray
}
