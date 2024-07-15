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
	Check(err)
	err = os.Mkdir(".got/com", 0755)
	Check(err)
    err = os.WriteFile(".got/com/cf",[]byte("0\n0\n"),0666)
	Check(err)
	println("Initialized empty got directory")
}
