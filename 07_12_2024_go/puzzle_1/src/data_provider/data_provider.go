package data_provider

import (
	"bufio"
	"strconv"
	"strings"

	icrud "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/file"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/logging"
)

type DataProvider struct {
	IO              icrud.CRUD
	ExpectedResults []int
	Elements        [][]int
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
	scanner.Split(bufio.ScanWords)
	temp := []int{}
	for scanner.Scan() {
		// Checking if the line contains the colon
		s := strings.Split(scanner.Text(), ":")

		// If there indeed was a colon
		logging.Info("Read", "el", s)
		if len(s) > 1 {
			conv_result, _ := strconv.Atoi(s[0])
			d.ExpectedResults = append(d.ExpectedResults, conv_result)
			if len(temp) > 0 {
				d.Elements = append(d.Elements, temp)
				temp = []int{}
			}
		} else {
			conv_result, _ := strconv.Atoi(s[0])
			temp = append(temp, conv_result)
		}

	}

	d.Elements = append(d.Elements, temp)

	logging.Info("Used data", "expected_results", d.ExpectedResults,
		"elements", d.Elements)

	if e := scanner.Err(); e != nil {
		logging.Error("Errors occured while scanning", "errors", e.Error())
		return false
	}

	logging.Info("Successfully saved the input to DataProvider{}", "path", path)
	return true
}

func (d *DataProvider) GetFirstArray() []int {
	return d.ExpectedResults
}

func (d *DataProvider) GetSecondArray() [][]int {
	return d.Elements
}
