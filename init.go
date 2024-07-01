package main

import (
	"bufio"
	"os"
)

func Init() {
    err := os.Mkdir(".got", 0775)
    if err != nil {
        println("Already initialized")
        os.Exit(1)
    }
    println("Initialized empty got directory")
}

func Add(fileName string) {
    file, err := os.OpenFile(".got/index",os.O_RDONLY|os.O_CREATE,0666)
    check(err)

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        if scanner.Text() == fileName {
            println("File already being tracked")
            return
        }
    }

    file.Close()

    file, err = os.OpenFile(".got/index",os.O_APPEND|os.O_WRONLY,0666)
    check(err)

    fileName = fileName + "\n"
    _, err = file.Write([]byte(fileName))
    file.Close()
    return
}


