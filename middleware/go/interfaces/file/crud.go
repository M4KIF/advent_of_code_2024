package icrud

import "os"

type CRUD interface {
	Open(string) (*os.File, error)
}
