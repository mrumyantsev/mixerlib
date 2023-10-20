package randomizer

import (
	"math/rand"
	"time"

	"github.com/mrumyantsev/passgen/internal/pkg/node"
)

type Randomizer struct {
	seed         int64
	randomSource rand.Source
	random       *rand.Rand
	node         *node.Node
}

func New() *Randomizer {
	return &Randomizer{}
}

func (r *Randomizer) ShuffleNodes(node *node.Node) {
	r.node = node

	nodeLength := r.node.GetLength()

	for i := 0; i < nodeLength; i++ {
		r.resetRandomSeed()
		r.shuffle()
	}
}

func (r *Randomizer) shuffle() {
	var (
		i int = r.node.GetLength() - 1
		j int
	)

	// Performs Fisher-Yates shuffle.

	for ; i > 1<<31-1-1; i-- {
		j = int(r.random.Int63n(int64(i + 1)))
		r.swap(i, j)
	}

	for ; i > 0; i-- {
		j = int(r.random.Int31n(int32(i + 1)))
		r.swap(i, j)
	}
}

func (r *Randomizer) swap(i, j int) {
	var (
		node1 = r.node.GetNode(i)
		node2 = r.node.GetNode(j)
	)

	r.node.SetNode(j, node1)
	r.node.SetNode(i, node2)
}

func (r *Randomizer) GetRandomIndex(nodesLength int) int {
	r.resetRandomSeed()

	return r.random.Intn(nodesLength)
}

func (r *Randomizer) resetRandomSeed() {
	r.seed = time.Now().UnixNano()
	r.randomSource = rand.NewSource(r.seed)
	r.random = rand.New(r.randomSource)
}
