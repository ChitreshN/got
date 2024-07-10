package main

import (
	"fmt"
	"github.com/got/lib"
	"os"
)

func main() {
	argLength := len(os.Args)
	if argLength < 2 {
		fmt.Println("usage: ./got command [options]")
		os.Exit(1)
	}
	command := os.Args[1]
	switch command {
	case "init":
		lib.Init()
	case "add":
		lib.Add(os.Args[2])
	case "status":
        tracked,untracked := lib.Status()
        println("The following lines are tracked:")
        for _, fileName := range tracked {
           println(fileName) 
        }
        println()
        println("The following lines are untracked:")
        for _, fileName := range untracked {
           println(fileName) 
        }
    case "diff":
        if argLength < 3 {
            fmt.Println("usage: ./got diff file1")
        }
	}
}
