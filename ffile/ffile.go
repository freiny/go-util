package ffile

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

// WalkFoldersFunc walks thru all folders in working directory and
// runs a callback function for each folder
func WalkFoldersFunc(f func(string) error) {
	fn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		if info.IsDir() {
			f(path)
		}
		return nil
	}

	path, _ := syscall.Getwd()
	filepath.Walk(path, fn)
}
