package data_provider

import (
	"bufio"
	"regexp"

	icrud "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/file"
	"github.com/M4KIF/advent_of_code_2024/middleware/go/logging"
)

type DataProvider struct {
	IO       icrud.CRUD
	Area     [][]string
	Commands []string
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

	area_done := false
	for scanner.Scan() {

		found := dec.FindAll([]byte(scanner.Text()), -1)
		if found != nil && !area_done {
			temp := []string{}
			for _, c := range found {
				temp = append(temp, string(c))
			}
			d.Area = append(d.Area, temp)
		} else if found != nil && area_done {
			for _, c := range found {
				d.Commands = append(d.Commands, string(c))
			}
		} else {
			// Loading the area completed, now to the commands part
			area_done = true
		}
	}

	if e := scanner.Err(); e != nil {
		logging.Error("Errors occured while scanning", "errors", e.Error())
		return false
	}

	// logging.Info("Successfully saved the input to DataProvider{}", "path", path,
	// 	"area", d.Area, "commands", d.Commands)
	return true
}

func (d *DataProvider) GetArea() [][]string {
	return d.Area
}

func (d *DataProvider) GetCommands() []string {
	return d.Commands
}
