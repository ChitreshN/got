package lib

import (
	"os"
)

func Init() {
	err := os.Mkdir(".got", 0775)
	if err != nil {
		println("Already initialized")
		os.Exit(1)
	}
	err = os.Mkdir(".got/obj", 0755)
	err = os.Mkdir(".got/com", 0755)
	Check(err)
	println("Initialized empty got directory")
}
