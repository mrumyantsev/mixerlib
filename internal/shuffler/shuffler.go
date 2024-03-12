package shuffler

import (
	"sync"

	"github.com/mrumyantsev/unguess/internal/types"
)

type Shuffler struct {
	types.Storer
	*sync.WaitGroup
}

func New(s types.Storer) *Shuffler {
	return &Shuffler{s, &sync.WaitGroup{}}
}

func (s *Shuffler) ShuffleMany(times int) {
	for i := 0; i < times; i++ {
		s.ResetData()
		s.Shuffle()
		s.MoveSource()
	}
}

func (s *Shuffler) Shuffle() {
	gCount := s.Cap()

	s.Add(gCount)

	for i := 0; i < gCount; i++ {
		go func(i int) {
			defer s.Done()

			s.Push(i)
		}(i)
	}

	s.Wait()
}
