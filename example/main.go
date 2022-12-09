package main

import (
	"flag"
	"fmt"

	gtkflags "github.com/ted537/gtk-flags"
)

var inputFlag = flag.String("input path", "", "Should be a file")
var outputFlag = flag.String("output path", "", "Should be a folder")
var strictFlag = flag.Bool("strict", false, "Check data for strict correctness")

func main() {
	gtkflags.RunGui("My custom application", doStuff)
}

func doStuff() error {
	if *strictFlag {
		return fmt.Errorf("strict not supported yet")
	}
	fmt.Println(*inputFlag, *outputFlag)
	return nil
}
