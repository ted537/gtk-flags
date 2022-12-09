package main

import (
	"flag"
	"fmt"
)

var inputFlag = flag.String("input path", "", "Should be a file")
var outputFlag = flag.String("output path", "", "Should be a folder")
var strictFlag = flag.Bool("strict", false, "Check data for strict correctness")

func main() {
	RunGui("My custom application", doStuff)
}

func doStuff() error {
	if *strictFlag {
		return fmt.Errorf("strict not supported yet")
	}
	fmt.Println(*inputFlag, *outputFlag)
	return nil
}
