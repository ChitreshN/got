package main

import (
	"os"
)


func main (){
    _, err := os.ReadFile("todo.md")
    check(err)
    Init()
    //fmt.Printf("%s", dat)
}
