package data_reader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Data struct {
	LeftArray  []int
	RightArray []int
}

func open_file(relativePath string) (*os.File, error) {
	// Getting the path to the executable
	executablePath, e := os.Executable()
	if e != nil {
		return nil, e
	}

	// Getting the path to parent from executable
	parentPath := filepath.Dir(executablePath)

	// Concating the relative path to parent
	fmt.Print(
		fmt.Sprintf("Parent path: %s\nRelative path: %s\n", parentPath, relativePath))
	finalPath := parentPath + relativePath

	// Opening the file

	if file, e := os.Open(finalPath); e != nil {
		return nil, e
	} else {
		return file, nil
	}
}

func (d *Data) Read(path string) bool {
	// Creating the absolute path and opening the file
	file, e := open_file(path)

	// Async file close
	defer file.Close()

	if e != nil {
		fmt.Printf("Error occured while opening the file: %s", e.Error())
		return false
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		// Dividing the line in order to extract the left and right entry
		l, r, f := strings.Cut(scanner.Text(), " ")
		if !f {
			continue
		}

		// Getting rid of whitespaces
		l = strings.ReplaceAll(l, " ", "")
		r = strings.ReplaceAll(r, " ", "")

		// Converting to integers
		left, e := strconv.Atoi(l)
		if e != nil {
			fmt.Printf("Error occured on converting the left string to int: %s\n", e.Error())
			return false
		}

		right, e := strconv.Atoi(r)
		if e != nil {
			fmt.Printf("Error occured on converting the right string to int: %s\n", e.Error())
			return false
		}

		// Appending the Arrays
		d.LeftArray = append(d.LeftArray, left)
		d.RightArray = append(d.RightArray, right)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (d Data) GetLeftArray() []int {
	return d.LeftArray
}

func (d Data) GetRightArray() []int {
	return d.RightArray
}
