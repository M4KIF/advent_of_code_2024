package idata

type TwoIntArrays interface {
	TakeInput(string) bool
	GetFirstArray() []int
	GetSecondArray() []int
}

type String2DArray interface {
	TakeInput(string) bool
	Get2DArray() [][]string
}

type SingleDimIntTwoDimIntArrays interface {
	TakeInput(string) bool
	GetFirstArray() []int
	GetSecondArray() [][]int
}
