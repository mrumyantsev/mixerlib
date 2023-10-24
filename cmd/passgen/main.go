package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mrumyantsev/passgen/internal/app/passgen"
)

const (
	DEFAULT_PASSWORD_LENGTH = 32
	VERSION                 = "0.9.1"
)

var (
	passGen           = passgen.New()
	passwordLength    = DEFAULT_PASSWORD_LENGTH
	isUserWantVersion = false
)

func init() {
	if len(os.Args) == 1 {
		return
	}

	secondArg := os.Args[1]

	if (secondArg == "-v") || (secondArg == "--version") {
		isUserWantVersion = true
		return
	}

	passLength, _ := strconv.Atoi(secondArg)

	if passLength > 0 {
		passwordLength = passLength
	}
}

func main() {
	if isUserWantVersion {
		fmt.Println("Version:", VERSION)
	} else {
		password := passGen.Generate(passwordLength)
		fmt.Println(string(password))
	}
}
