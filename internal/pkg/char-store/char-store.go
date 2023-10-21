package charstore

import (
	"github.com/mrumyantsev/passgen/internal/pkg/node"
	"github.com/mrumyantsev/passgen/internal/pkg/randomizer"
)

type CharStore struct {
	charGroups           *node.Node
	randomizer           *randomizer.Randomizer
	availableGroups      []int
	availableGroupsCount int
	currentCharsIndex    int
	previousCharsIndex   int
}

func New(randomizer *randomizer.Randomizer) *CharStore {
	store := &CharStore{}

	store.randomizer = randomizer

	return store
}

func (c *CharStore) Shuffle() {
	var (
		charGroupsLength = c.charGroups.GetLength()
		chars            *node.Node
	)

	for i := 0; i < charGroupsLength; i++ {
		chars = c.charGroups.GetNode(i)
		c.randomizer.ShuffleNodes(chars)
		c.charGroups.SetNode(i, chars)
	}

	c.randomizer.ShuffleNodes(c.charGroups)
}

func (c *CharStore) GetCharacter() byte {
	var chars *node.Node

	for {
		if c.availableGroupsCount == 0 {
			return '0'
		} else if c.availableGroupsCount == 1 {
			c.currentCharsIndex = 0
		} else {
			c.randomizeCharsIndex()
		}

		chars = c.charGroups.GetNode(c.availableGroups[c.currentCharsIndex])

		if chars.GetLength() == 0 {
			c.remakeAvailableGroups()

			continue
		}

		break
	}

	value := chars.Pop().GetValue()

	return value
}

func (c *CharStore) randomizeCharsIndex() {
	for {
		c.currentCharsIndex = c.randomizer.GetRandomIndex(c.availableGroupsCount)

		if c.currentCharsIndex != c.previousCharsIndex {
			break
		}
	}

	c.previousCharsIndex = c.currentCharsIndex
}

func (c *CharStore) remakeAvailableGroups() {
	c.availableGroupsCount--

	if c.availableGroupsCount == 0 {
		return
	}

	var (
		newAvailableGroups = make([]int, c.availableGroupsCount)
		j                  = 0
	)

	for index, id := range c.availableGroups {
		if index == c.currentCharsIndex {
			continue
		}

		newAvailableGroups[j] = id
		j++
	}

	c.availableGroups = newAvailableGroups

	c.previousCharsIndex = c.randomizer.GetRandomIndex(c.availableGroupsCount)
}
