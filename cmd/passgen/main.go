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
	passwordLength = _DEFAULT_PASSWORD_LENGTH
)

func init() {
	if len(os.Args) > 1 {
		param, err := strconv.Atoi(os.Args[1])
		if err != nil {
			return
		}

		if param > 0 {
			passwordLength = param
		}
	}
}

func main() {
	passGen := passgen.New()

	password := passGen.Generate(passwordLength)

	os.Stdout.Write(password)
	os.Stdout.WriteString(_END_OF_LINE)
}
