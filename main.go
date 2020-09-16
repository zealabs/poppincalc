package main

import (
	"./pkg/modules"
)

func main() {
	// test
	modules.StartJsAnalysis("./examples/vulned.js")
	//modules.StartPyAnalysis()
}
