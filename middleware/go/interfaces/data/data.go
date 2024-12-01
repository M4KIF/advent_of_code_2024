package idata

import icrud "github.com/M4KIF/advent_of_code_2024/middleware/go/interfaces/file"

type TwoIntArrays interface {
	TakeInput(string, icrud.CRUD) bool
	GetFirstArray() []int
	GetSecondArray() []int
}
