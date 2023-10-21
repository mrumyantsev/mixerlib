package charstore

import (
	"github.com/mrumyantsev/passgen/internal/consts"
	"github.com/mrumyantsev/passgen/internal/pkg/node"
)

func (c *CharStore) Init() {
	c.charGroups = node.Make(consts.DEFAULT_CHAR_GROUPS_CAPACITY)

	charGroupsData := []struct {
		pattern     string
		charGroupId byte
	}{
		{"0-9", consts.NUMBERS},
		{"@#$%&", consts.SPEC_CHARS},
		{"a-z", consts.LOW_LETTERS},
		{"A-Z", consts.HIGH_LETTERS},
	}

	var node *node.Node

	for _, chars := range charGroupsData {
		node = makeCharsWithPattern(chars.pattern)
		node.SetValue(chars.charGroupId)
		c.charGroups.Push(node)
	}

	c.availableGroups = []int{
		consts.NUMBERS,
		consts.SPEC_CHARS,
		consts.LOW_LETTERS,
		consts.HIGH_LETTERS,
	}

	c.availableGroupsCount = consts.DEFAULT_CHAR_GROUPS_CAPACITY
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
		chars = node.Make(int(endChar-startChar) + 1)
	)

	for ch := startChar; ch <= endChar; ch++ {
		charItem := node.New()
		charItem.SetValue(ch)
		chars.Push(charItem)
	}

	return chars
}

func makeCharsWithSpecifiedChars(specifiedChars string) *node.Node {
	var (
		chars = node.Make(len(specifiedChars))
	)

	for _, ch := range specifiedChars {
		charItem := node.New()
		charItem.SetValue(byte(ch))
		chars.Push(charItem)
	}

	return chars
}
