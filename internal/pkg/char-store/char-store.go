package charstore

import (
	"github.com/mrumyantsev/passgen/internal/pkg/consts"
	"github.com/mrumyantsev/passgen/internal/pkg/node"
	"github.com/mrumyantsev/passgen/internal/pkg/randomizer"
)

type CharStore struct {
	collection        *node.Node
	count             int
	availableCharSets []int
	index             int
	previousIndex     int
	randomizer        *randomizer.Randomizer
}

func New(randomizer *randomizer.Randomizer) *CharStore {
	store := &CharStore{}

	store.randomizer = randomizer

	return store
}

// Gets new character from the collection.
func (c *CharStore) GetCharacter() byte {
	var charSet *node.Node

	for {
		if c.count == 0 {
			return consts.STUB_FOR_NO_CHARS
		} else if c.count == 1 {
			c.index = 0
		} else {
			c.randomizeCharsIndex()
		}

		charSet = c.collection.GetNode(c.availableCharSets[c.index])

		if charSet.GetCount() == 0 {
			c.remakeAvailable()

			continue
		}

		break
	}

	value := charSet.Pop().GetValue()

	return value
}

func (c *CharStore) randomizeCharsIndex() {
	for {
		c.index = c.randomizer.GetRandomIndex(c.count)

		if c.index != c.previousIndex {
			break
		}
	}

	c.previousIndex = c.index
}

func (c *CharStore) remakeAvailable() {
	c.count--

	if c.count == 0 {
		return
	}

	var (
		newavailableCharSets = make([]int, c.count)
		i                    = 0
	)

	for index, id := range c.availableCharSets {
		if index == c.index {
			continue
		}

		newavailableCharSets[i] = id
		i++
	}

	c.availableCharSets = newavailableCharSets

	c.previousIndex = c.randomizer.GetRandomIndex(c.count)
}
