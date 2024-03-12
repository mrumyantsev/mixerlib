package store

import "sync"

type Store struct {
	source   []rune
	data     []rune
	capacity int
	index    int
	mu       *sync.Mutex
}

func New(source []rune) *Store {
	return &Store{
		source:   source,
		capacity: len(source),
		mu:       &sync.Mutex{},
	}
}

func (s *Store) Data() []rune {
	return s.data
}

func (s *Store) Cap() int {
	return s.capacity
}

func (s *Store) DecCap() {
	s.capacity--
}

func (s *Store) ResetData() {
	s.data = make([]rune, s.capacity)
	s.index = 0
}

func (s *Store) Push(index int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[s.index] = s.source[index]
	s.index++
}

func (s *Store) MoveSource() {
	s.source = s.data
}

func (s *Store) MoveData() {
	s.data = s.source
}
