package data_provider

import (
	"bufio"
	"fmt"

	icrud "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/file"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/logging"
)

type DataProvider struct {
	IO   icrud.CRUD
	Data [][]string
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
	for scanner.Scan() {
		var row []string

		// Utilized @github mnml code for splitting the read line
		var word string
		fmt.Sscanf(scanner.Text(), "%s", &word)

		// Appending the array
		for _, char := range word {
			row = append(row, string(char))
		}

		d.Data = append(d.Data, row)
	}

	if e := scanner.Err(); e != nil {
		logging.Error("Errors occured while scanning", "errors", e.Error())
		return false
	}

	logging.Info("Successfully saved the input to DataProvider{}", "path", path)
	return true
}

func (d *DataProvider) Get2DArray() [][]string {
	return d.Data
}
