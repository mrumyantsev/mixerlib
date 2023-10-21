package randomizer

import (
	"math/rand"
	"time"

	"github.com/mrumyantsev/passgen/internal/pkg/node"
)

type Randomizer struct {
	seed   int64
	source rand.Source
	random *rand.Rand
}

func New() *Randomizer {
	return &Randomizer{}
}

func (r *Randomizer) ShuffleNodes(node *node.Node) {
	length := node.GetCount()

	for i := 0; i < length; i++ {
		r.resetRandomSeed()
		r.shuffle(node)
	}
}

func (r *Randomizer) shuffle(node *node.Node) {
	var (
		i int = node.GetCount() - 1
		j int
	)

	// Performs Fisher-Yates shuffle.
	for ; i > 1<<31-1-1; i-- {
		j = int(r.random.Int63n(int64(i + 1)))
		r.swap(node, i, j)
	}

	for ; i > 0; i-- {
		j = int(r.random.Int31n(int32(i + 1)))
		r.swap(node, i, j)
	}
}

func (r *Randomizer) swap(node *node.Node, i, j int) {
	var (
		node1 = node.GetNode(i)
		node2 = node.GetNode(j)
	)

	node.SetNode(j, node1)
	node.SetNode(i, node2)
}

func (r *Randomizer) GetRandomIndex(nodesLength int) int {
	r.resetRandomSeed()

	return r.random.Intn(nodesLength)
}

func (r *Randomizer) resetRandomSeed() {
	r.seed = time.Now().UnixNano()
	r.source = rand.NewSource(r.seed)
	r.random = rand.New(r.source)
}
