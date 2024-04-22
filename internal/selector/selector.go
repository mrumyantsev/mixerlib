package selector

import "github.com/mrumyantsev/mixer/internal/core"

type Selector struct {
	core.Shuffler
}

func New(s core.Shuffler) *Selector {
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
