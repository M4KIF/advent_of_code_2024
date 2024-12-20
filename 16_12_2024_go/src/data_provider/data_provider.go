package data_provider

import (
	"bufio"
	"fmt"
	"regexp"

	icrud "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/file"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/logging"
)

type DataProvider struct {
	IO    icrud.CRUD
	Area  [][]string
	Start []int
	End   []int
}

func NewDataProvider(io icrud.CRUD) *DataProvider {
	return &DataProvider{IO: io}
}

func (d *DataProvider) TakeInput(path string) bool {
	// Opening the file
	file, e := d.IO.Open(path)
	if e != nil {
		logging.Error("Oribl!", "error", e.Error())
	}

	defer file.Close()

	// Resolving the input
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	dec, _ := regexp.Compile(`.`)

	y := 0
	for scanner.Scan() {

		found := dec.FindAll([]byte(scanner.Text()), -1)
		if found != nil {
			temp := []string{}
			x := 0
			for _, c := range found {
				// Finding the starting and ending point in the mean-time
				if string(c) == "S" {
					fmt.Println("SZTARTA")
					d.Start = []int{y, x}
				} else if string(c) == "E" {
					d.End = []int{y, x}
				}
				temp = append(temp, string(c))
				x++
			}
			d.Area = append(d.Area, temp)
		}
		y++
	}

	if e := scanner.Err(); e != nil {
		logging.Error("Errors occured while scanning", "errors", e.Error())
		return false
	}

	// logging.Info("Successfully saved the input to DataProvider{}", "path", path,
	// 	"area", d.Area)
	return true
}

func (d *DataProvider) GetArea() [][]string {
	return d.Area
}

func (d *DataProvider) GetStartPoint() []int {
	return d.Start
}

func (d *DataProvider) GetEndPoint() []int {
	return d.End
}
