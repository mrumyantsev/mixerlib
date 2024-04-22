package selector

import "github.com/mrumyantsev/mixer/internal/types"

type Selector struct {
	types.Shuffler
}

func New(s types.Shuffler) *Selector {
	return &Selector{s}
}

func (s *Selector) Select(count int) {
	if count > s.Cap() {
		s.MoveData()

		return
	}

	if count < 0 {
		count = 0
	}

	count = s.Cap() - count

	for i := 0; i < count; i++ {
		s.DecCap()
		s.ResetData()
		s.Shuffle()
		s.MoveSource()
	}
}
