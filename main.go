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
        if argLength < 4 {
            fmt.Println("usage: ./got diff file1 file2")
        }
        file1, err := os.OpenFile(os.Args[2],os.O_RDONLY,0666)
        lib.Check(err)
        file2, err := os.OpenFile(os.Args[3],os.O_RDONLY,0666)
        lib.Check(err)
        diff := lib.Diff(file1,file2)
        for _, val := range diff {
            switch val.EditType{
            case lib.Append :
                fmt.Printf("\033[32m+%s\033[0m\n",val.Append)
            case lib.Delete :
                fmt.Printf("\033[31m-%s\033[0m\n",val.Delete)
            case lib.Identical :
                println(val.Identical)
            }
        }
	}
}
