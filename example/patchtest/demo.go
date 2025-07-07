package patchtest

import (
	"os"
)

func ChangeMode(path string) error {
	err := os.Chmod(path, os.ModePerm)
	return err
}
