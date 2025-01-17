package main

import (
	"fmt"
	"github.com/ChitreshN/got/lib"
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
        lib.RunStatus()
    case "diff":
        if argLength < 3 {
            fmt.Println("usage: ./got diff file1")
        }
        lib.RunDiff()
    case "commit":
        lib.Commit("some message")
    case "prev_commit":
        lib.LatestCommit()
    case "revert":
        lib.Revert()
    case "checkout":
        err := lib.Checkout(os.Args[2])
        lib.Check(err)
	}
    //how do this
    //i am shtupid, i am shtupid
}
