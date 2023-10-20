package charstore

import (
	"github.com/mrumyantsev/password-generator/internal/pkg/node"
	"github.com/mrumyantsev/password-generator/internal/pkg/randomizer"
)

type CharStore struct {
	charGroups             *node.Node
	randomizer             *randomizer.Randomizer
	currentCharGroupIndex  int
	previousCharGroupIndex int
}

func New(randomizer *randomizer.Randomizer) *CharStore {
	store := &CharStore{}

	store.randomizer = randomizer

	return store
}

func (c *CharStore) Shuffle() {
	var (
		charGroupsLength = c.charGroups.GetLength()
		charGroup        *node.Node
	)

	for i := 0; i < charGroupsLength; i++ {
		charGroup = c.charGroups.GetNode(i)
		c.randomizer.ShuffleNodes(charGroup)
		c.charGroups.SetNode(i, charGroup)
	}

	c.randomizer.ShuffleNodes(c.charGroups)
}

func (c *CharStore) GetCharacter() byte {
	var charGroup *node.Node

	for {
		if c.charGroups.GetLength() == 0 {
			return '0'
		}

		charGroup = c.GetCharGroup()

		if charGroup.GetLength() == 0 {
			c.remakeCharGroups()
			continue
		}

		break
	}

	value := charGroup.Pop().GetValue()

	return value
}

func (c *CharStore) remakeCharGroups() {
	var (
		newCapacity = c.charGroups.GetLength() - 1
	)

	if newCapacity == 0 {
		c.charGroups = node.New()
		return
	}

	var (
		charGroups = node.Make(newCapacity)
		charGroup  *node.Node
	)

	for i := 0; i <= newCapacity; i++ {
		if i == c.currentCharGroupIndex {
			continue
		}

		charGroup = c.charGroups.GetNode(i)

		charGroups.Push(charGroup)
	}

	c.charGroups = charGroups

	c.previousCharGroupIndex = c.randomizer.GetRandomIndex(newCapacity)
}

func (c *CharStore) GetCharGroup() *node.Node {
	for {
		c.currentCharGroupIndex = c.randomizer.GetRandomIndex(c.charGroups.GetLength())

		if c.charGroups.GetLength() == 1 {
			break
		}

		if c.currentCharGroupIndex != c.previousCharGroupIndex {
			break
		}
	}

	c.previousCharGroupIndex = c.currentCharGroupIndex

	return c.charGroups.GetNode(c.currentCharGroupIndex)
}
