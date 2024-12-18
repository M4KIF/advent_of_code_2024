package data_provider

import (
	"testing"

	"github.com/M4KIF/advent_of_code_2024/middleware/go/file_handling"
	"github.com/stretchr/testify/assert"
)

func TestDataProviderEmptyFile(t *testing.T) {

	path := "./test_data/data_1.txt"

	io_provider := file_handling.Test{}

	provider := NewDataProvider(io_provider)

	provider.TakeInput(path)

	assert.Equal(t, len(provider.Get2DArray()), 0)
}
