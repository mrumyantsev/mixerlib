package charstore

import (
	"github.com/mrumyantsev/passgen/internal/pkg/node"
)

// Shuffles chars in the char sets and char sets in the collection.
func (c *CharStore) Shuffle() {
	var (
		collectionCount = c.collection.GetCount()
		charSet         *node.Node
	)

	for i := 0; i < collectionCount; i++ {
		charSet = c.collection.GetNode(i)
		c.randomizer.ShuffleNodes(charSet)
		c.collection.SetNode(i, charSet)
	}

	c.randomizer.ShuffleNodes(c.collection)
}
