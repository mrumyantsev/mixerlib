package gen

import (
	"github.com/mrumyantsev/unguess/internal/selector"
	"github.com/mrumyantsev/unguess/internal/shuffler"
	"github.com/mrumyantsev/unguess/internal/store"
)

const (
	// Default shuffle times value that used by characters generator.
	ShufflesDefault = 3000
)

var (
	shuffles = ShufflesDefault
)

// Generate generates shuffled sequence of characters from the source.
//
// The shuffle algorithm is based on Go programming language
// concurrency uncertainties like calling the goroutines in
// inconsistent order.
//
// If nil is passed as a src, then the source will be set by default
// using the SourceDefault function return value. If 0 or lower, or
// higher than source length number is passed as len, the length of
// output slice will be the same as the source length.
//
// Returns a slice of runes (the same runes that passed in src, but
// very well shuffled, and cut by len).
func Generate(src []rune, len int) []rune {
	if src == nil {
		src = SourceDefault()
	}

	store := store.New(src)

	shuf := shuffler.New(store)

	sel := selector.New(shuf)

	shuf.ShuffleMany(shuffles)

	sel.Select(len)

	return sel.Data()
}

// SourceDefault gets the default characters set used by characters
// generator if nil is passed as source parameter of Generate function.
//
// Returning slice consists of 67 runes:
//   - 5 special characters (@, #, $, %, &)
//   - 10 Western Arabic numerals (0-9)
//   - 26 uppercase latin alphabet letters (A-Z)
//   - 26 lowercase latin alphabet letters (a-z)
func SourceDefault() []rune {
	return []rune{
		'@', '#', '$', '%', '&',

		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',

		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
		'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y', 'Z',

		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y', 'z',
	}
}

// Shuffles gets current number of shuffles that happen inside Generate
// function.
func Shuffles() int {
	return shuffles
}

// SetShuffles sets how much times the source character sequence will
// be shuffled inside Generate function.
func SetShuffles(num int) {
	shuffles = num
}
