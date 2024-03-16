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

func main() {
	parsePassLen()

	pass := gen.Generate(nil, passLen)

	if len(pass) == 0 {
		return
	}

	fmt.Println(string(pass))
}

// parsePassLen parses the second CLI argument as password length.
//
// If the password length argument is not specified or could not parse
// as a number then password length stays declared defaultly.
func parsePassLen() {
	if len(os.Args) < 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	passLen = n
}
