package main

import (
	"os"
	"strconv"

	"github.com/mrumyantsev/passgen/internal/app/passgen"
)

const (
	DEFAULT_PASSWORD_LENGTH = 32
	END_OF_LINE             = "\n"
	VERSION                 = "0.9.2"
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
		os.Stdout.WriteString("Version: " + VERSION + END_OF_LINE)
	} else {
		password := passGen.Generate(passwordLength)

		os.Stdout.Write(password)
		os.Stdout.WriteString(END_OF_LINE)
	}
}
