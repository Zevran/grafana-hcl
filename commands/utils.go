package commands

import (
	"os"
	"path/filepath"
)

func Getwd() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	return dir, err
}
