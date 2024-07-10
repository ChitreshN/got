package lib

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
	file, err := os.OpenFile(".got/index", os.O_RDONLY|os.O_CREATE, 0666)
	Check(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == fileName {
			println("File already being tracked")
			return
		}
	}

	file.Close()

	file, err = os.OpenFile(".got/index", os.O_APPEND|os.O_WRONLY, 0666)
	Check(err)

	fileName = fileName + "\n"
	_, err = file.Write([]byte(fileName))
	file.Close()
}

func Status() (trackedFiles ,untrackedFiles []string){
	directory, err := os.Getwd()
	Check(err)

	fileList, err := os.ReadDir(directory)
	Check(err)

	untrackedFiles = make([]string, 0)
    trackedFiles = make([]string,0)

	for _, fileInfo := range fileList {

		tracked := false

		if fileInfo.IsDir() {
			continue
		}

		file, err := os.OpenFile(".got/index", os.O_RDONLY|os.O_CREATE, 0666)
		Check(err)

		fileName := fileInfo.Name()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if scanner.Text() == fileName {
				tracked = true
                trackedFiles = append(trackedFiles, fileName)
				break
			}
		}
		if !tracked {
			untrackedFiles = append(untrackedFiles, fileName)
		}
	}
    
    return
}
