package main

import (
	"flag"
	"fmt"
)

var nameFlag = flag.String("name", "example name", "choose a name!!!")
var reverseFlag = flag.Bool("reverse", false, "should we reverse?")
var x = flag.CommandLine.Int64("number", 23, "set a number :)")

func main() {
	RunGui("My custom application", doStuff)
}

func doStuff() error {
	if *reverseFlag {
		return fmt.Errorf("Reverse not supported")
	}
	fmt.Println(*nameFlag, *reverseFlag, *x)
	return nil
}
