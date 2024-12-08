package data_provider

import (
	icrud "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/file"
)

type DataProvider struct {
	IO   icrud.CRUD
	Data [][]string
}

func NewDataProvider(io icrud.CRUD) *DataProvider {
	return &DataProvider{IO: io}
}

func (d *DataProvider) TakeInput(path string) bool {
	return false
}

func (d *DataProvider) Get2DArray() [][]string {
	return d.Data
}
