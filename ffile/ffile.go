package ffile

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

// WalkFoldersFunc ...
func WalkFoldersFunc(f func(string)) {
	// f(path)
	// base := "/Users/name/dev/golang/src/github.com/freiny/ztmp"
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
