package main

import (
	"os"
	"strconv"

	"github.com/mrumyantsev/passgen"
)

const (
	_DEFAULT_PASSWORD_LENGTH = 32
	_END_OF_LINE             = "\n"
)

var (
	passGen        = passgen.New()
	passwordLength = _DEFAULT_PASSWORD_LENGTH
)

func init() {
	if len(os.Args) == 1 {
		return
	}

	secondArg := os.Args[1]

	passLength, _ := strconv.Atoi(secondArg)

	if passLength > 0 {
		passwordLength = passLength
	}
}

func main() {
	password := passGen.Generate(passwordLength)

	os.Stdout.Write(password)
	os.Stdout.WriteString(_END_OF_LINE)
}
