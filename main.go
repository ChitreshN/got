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
        lib.RunStatus()
    case "diff":
        if argLength < 3 {
            fmt.Println("usage: ./got diff file1")
        }
        lib.RunDiff()
    case "commit":
        lib.Commit("some message")
    case "prev_commit":
        if argLength < 3 {
            fmt.Println("usage: ./got prev_commit file1")
        }
        lib.ConstLatestCommit(os.Args[2])
	}
    //how do this
    //i am shtupid, i am shtupid
}
