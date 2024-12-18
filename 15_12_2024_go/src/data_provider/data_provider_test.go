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

	assert.Equal(t, len(provider.GetArea()), 0)
	assert.Equal(t, len(provider.GetCommands()), 0)
}

func TestDataProviderData1Commands15Area8by8(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}

	provider := NewDataProvider(io_provider)

	provider.TakeInput(path)

	assert.Equal(t, len(provider.GetArea()), 8)
	assert.Equal(t, len(provider.GetCommands()), 15)
}

func TestDataProviderData1Commands15Area8by8FirstLine(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}

	provider := NewDataProvider(io_provider)

	provider.TakeInput(path)

	assert.Equal(t, provider.GetArea()[0], []string{"#", "#", "#", "#", "#", "#", "#", "#"})
}

func TestDataProviderData1Commands15Area8by8ThirdLine(t *testing.T) {

	path := "./test_data/data_2.txt"

	io_provider := file_handling.Test{}

	provider := NewDataProvider(io_provider)

	provider.TakeInput(path)

	assert.Equal(t, provider.GetArea()[2], []string{"#", "#", "@", ".", "O", ".", ".", "#"})
}

func TestDataProviderData1Commands700RegexIdeaArea10(t *testing.T) {

	path := "./test_data/data_3.txt"

	io_provider := file_handling.Test{}

	provider := NewDataProvider(io_provider)

	provider.TakeInput(path)

	assert.Equal(t, len(provider.GetArea()), 10)
	assert.Equal(t, len(provider.GetCommands()), 700)
}
