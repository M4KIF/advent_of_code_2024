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

type Int2DArray interface {
	TakeInput(string) bool
	Get2DArray() [][]int
}

type String2Dand1Darray interface {
	TakeInput(string) bool
	GetArea() [][]string
	GetCommands() []string
}

type StringArea2DStartEnd interface {
	TakeInput(string) bool
	GetArea() [][]string
	GetStartPoint() [2]int
	GetEndPoint() [2]int
}
