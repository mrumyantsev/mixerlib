package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mrumyantsev/unguess/pkg/gen"
)

var (
	passLen = 32
)

func init() {
	n, ok := parsePassLen()
	if ok {
		passLen = n
	}
}

func main() {
	pass := gen.Generate(nil, passLen)

	if len(pass) == 0 {
		return
	}

	fmt.Println(string(pass))
}

// parsePassLen parses the second CLI argument as password length and
// returns its value and parse status. If the argument is not specified
// or is not a number then returning status will be false.
func parsePassLen() (length int, ok bool) {
	if len(os.Args) == 1 {
		return
	}

	var err error

	length, err = strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	ok = true

	return
}
