package main

import (
	"os"
	"strconv"

	"github.com/mrumyantsev/password-generator/internal/app/passgen"
)

const (
	DEFAULT_PASSWORD_LENGTH int    = 32
	END_OF_LINE             string = "\n"
)

func main() {
	var (
		passGen        = passgen.New()
		passwordLength = getPasswordLength()
		password       = passGen.Generate(passwordLength)
	)

	os.Stdout.Write(password)
	os.Stdout.WriteString(END_OF_LINE)
}

func getPasswordLength() int {
	var (
		args   = os.Args
		length int
	)

	if len(args) > 1 {
		length, _ = strconv.Atoi(args[1])

		if length > 0 {
			return length
		}
	}

	return DEFAULT_PASSWORD_LENGTH
}
