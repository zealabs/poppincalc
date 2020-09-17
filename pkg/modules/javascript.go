package modules

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func StartJsAnalysis(path string) {
	// analysis starts here
	filePath, err := filepath.Abs(path)
	check(err)

	//read file
	codeFile, err := ioutil.ReadFile(filePath)
	check(err)

	fmt.Print(string(codeFile))
}
