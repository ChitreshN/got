package main

import (
	"fmt"
	"os"
)

func main (){
    argLength := len(os.Args)
    if argLength < 2 {
        fmt.Println("usage: ./got command [options]")
    }
    command := os.Args[1]
    switch command {
    case "init":
        Init()
    case "add":
        Add(os.Args[2])
    }
}
