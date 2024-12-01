package file_handling

import (
	"os"
	"path/filepath"

	"github.com/M4KIF/advent_of_code_2024/middleware/go/logging"
)

type Default struct{}

func (d Default) Open(relativePath string) (*os.File, error) {
	// Getting the path to the executable
	executablePath, e := os.Executable()
	if e != nil {
		return nil, e
	}

	// Getting the path to parent from executable
	parentPath := filepath.Dir(executablePath)

	// Concating the relative path to parent
	finalPath := parentPath + relativePath
	logging.Info(
		"Paths used",
		"relative", relativePath,
		"parent", parentPath,
		"final", finalPath)

	// Opening the file

	if file, e := os.Open(finalPath); e != nil {
		return nil, e
	} else {
		return file, nil
	}
}
