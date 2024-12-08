package data_provider

// import (
// 	"os"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// type IoMock struct {
// }

// func (i IoMock) Open(path string) (*os.File, error) {
// 	if file, e := os.Open(path); e != nil {
// 		return nil, e
// 	} else {
// 		return file, nil
// 	}
// }

// func NewMockProvider() *DataProvider {
// 	return NewDataProvider(IoMock{})
// }

// func TestDataProviderCustom(t *testing.T) {
// 	dp := NewMockProvider()
// 	dp.TakeInput("test_input.txt")

// 	read := dp.Get2DArray()

// 	assert.Equal(t, len(read), 5)

// 	for row := range read {
// 		assert.Equal(t, len(read[row]), 6)
// 	}

// }

// func TestDataProviderAocData(t *testing.T) {
// 	// Just Kidding XD
// }
