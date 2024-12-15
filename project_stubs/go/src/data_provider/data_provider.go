package data_provider

import (
	"bufio"
	"regexp"
	"strconv"

	icrud "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/file"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/logging"
)

type DataProvider struct {
	IO   icrud.CRUD
	Data [][]int
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

	counter := 0
	temp := []int{}
	for scanner.Scan() {
		// If the presumed data set finishes,
		// assign a new int array
		if counter == 3 {
			counter = 0
			d.Data = append(d.Data, temp)
			temp = nil
			temp = []int{}
		}

		dec, _ := regexp.Compile(`\d+`)
		found := dec.FindAll([]byte(scanner.Text()), -1)
		if found != nil {
			for _, d := range found {
				int_value, _ := strconv.Atoi(string(d))
				temp = append(temp, int_value)
			}

			counter += 1
		}
	}
	if len(temp) > 0 {
		d.Data = append(d.Data, temp)
	}

	if e := scanner.Err(); e != nil {
		logging.Error("Errors occured while scanning", "errors", e.Error())
		return false
	}

	logging.Info("Successfully saved the input to DataProvider{}", "path", path)
	return true
}

func (d *DataProvider) Get2DArray() [][]int {
	return d.Data
}
