package lib

import (
	"bufio"
	"os"
)

func Add(fileName string) {
	staged, err := os.OpenFile(".got/staged", os.O_RDONLY|os.O_CREATE, 0666)

	stgScanner := bufio.NewScanner(staged)

	for stgScanner.Scan() {
		if stgScanner.Text() == fileName {
			println("File already staged")
			return
		}
	}
	staged.Close()

	index, err := os.OpenFile(".got/index", os.O_RDONLY|os.O_CREATE, 0666)
	Check(err)
	idxScanner := bufio.NewScanner(index)

	for idxScanner.Scan() {
		if idxScanner.Text() == fileName {
			staged, err = os.OpenFile(".got/staged", os.O_WRONLY|os.O_APPEND, 0666)
			fileName = fileName + "\n"
			staged.Write([]byte(fileName))
			staged.Close()
			return
		}
	}
	index.Close()

	index, err = os.OpenFile(".got/index", os.O_APPEND|os.O_WRONLY, 0666)
	Check(err)
	staged, err = os.OpenFile(".got/staged", os.O_WRONLY|os.O_APPEND, 0666)
	Check(err)

	fileName = fileName + "\n"
	_, err = index.Write([]byte(fileName))
	Check(err)
	_, err = staged.Write([]byte(fileName))
	Check(err)
	staged.Close()
	index.Close()
}
