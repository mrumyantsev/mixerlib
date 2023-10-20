package charstore

import (
	"github.com/mrumyantsev/password-generator/internal/pkg/node"
)

const (
	DEFAULT_CHAR_GROUPS_CAPACITY int = 4
)

func (c *CharStore) Init() {
	c.charGroups = node.Make(DEFAULT_CHAR_GROUPS_CAPACITY)

	c.charGroups.Push(makeCharsWithPattern("0-9"))
	c.charGroups.Push(makeCharsWithPattern("@#$%&"))
	c.charGroups.Push(makeCharsWithPattern("a-z"))
	c.charGroups.Push(makeCharsWithPattern("A-Z"))
}

func makeCharsWithPattern(pattern string) *node.Node {
	if len(pattern) == 3 {
		if pattern[1] == '-' {
			return makeCharsWithRange(pattern[0], pattern[2])
		}
	}

	return makeCharsWithSpecifiedChars(pattern)
}

func makeCharsWithRange(startChar, endChar byte) *node.Node {
	var (
		charGroup = node.Make(int(endChar-startChar) + 1)
	)

	for ch := startChar; ch <= endChar; ch++ {
		charItem := node.New()
		charItem.SetValue(ch)
		charGroup.Push(charItem)
	}

	return charGroup
}

func makeCharsWithSpecifiedChars(specifiedChars string) *node.Node {
	var (
		charGroup = node.Make(len(specifiedChars))
	)

	for _, ch := range specifiedChars {
		charItem := node.New()
		charItem.SetValue(byte(ch))
		charGroup.Push(charItem)
	}

	return charGroup
}
