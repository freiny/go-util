package ftest

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/freiny/go-util/ffile"
)

// TestAll walks thru all folders in the working directory
// running "go test" in each folder
func TestAll() {
	f := func(path string) error {
		os.Chdir(path)
		wd, _ := os.Getwd()

		cmd := "go"
		args := []string{"test"}

		out, err := exec.Command(cmd, args...).Output()
		o := strings.Split(string(out), "\n")[0]
		isPass := strings.Contains(o, "PASS")
		isFail := strings.Contains(o, "FAIL")
		if err != nil {
			if isFail {
				fmt.Println("FAIL:", wd)
				fmt.Println(string(out))
			}
			return fmt.Errorf("ERROR: %v", err)
		}

		if isPass {
			fmt.Println("PASS:", wd)
		}
		return nil
	}
	ffile.WalkFoldersFunc(f)
}
