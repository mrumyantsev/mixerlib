package main

import (
	"os"
	"strconv"

	"github.com/mrumyantsev/passgen"
)

const (
	_DEFAULT_PASSWORD_LENGTH = 32
	_END_OF_LINE             = "\n"
	_VERSION                 = "0.10.0"
)

var (
	passGen           = passgen.New()
	passwordLength    = _DEFAULT_PASSWORD_LENGTH
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
		os.Stdout.WriteString("Version: " + _VERSION + _END_OF_LINE)
	} else {
		password := passGen.Generate(passwordLength)

		os.Stdout.Write(password)
		os.Stdout.WriteString(_END_OF_LINE)
	}
}
