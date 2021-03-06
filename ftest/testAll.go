package ftest

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/freiny/go-util/ffile"
)

// TestAll walks thru all folders in the working directory
// running "go test" in each folder
func TestAll() bool {
	nTests := 0
	nPassed := 0

	f := func(folder string) error {
		b := path.Base(folder)
		if b[0] == byte('.') && len(b) > 1 {
			return nil
		}
		os.Chdir(folder)
		wd, _ := os.Getwd()

		cmd := "go"
		args := []string{"test"}

		out, err := exec.Command(cmd, args...).Output()
		o := strings.Split(string(out), "\n")[0]
		isPass := strings.Contains(o, "PASS")
		isFail := strings.Contains(o, "FAIL")
		if err != nil {
			if isFail {
				nTests++
				fmt.Println("FAIL:", wd)
				fmt.Println(string(out))
			}
			return fmt.Errorf("ERROR: %v", err)
		}

		if isPass {
			nTests++
			nPassed++
			fmt.Println("PASS:", wd)
		}
		return nil
	}
	ffile.WalkFoldersFunc(f)
	return nPassed == nTests
}
