package store

import "sync"

type Store struct {
	source   []byte
	data     []byte
	capacity int
	index    int
	mu       *sync.Mutex
}

func New(src []byte) *Store {
	return &Store{
		source:   src,
		capacity: len(src),
		mu:       &sync.Mutex{},
	}
}

func (s *Store) Data() []byte {
	return s.data
}

func (s *Store) Cap() int {
	return s.capacity
}

func (s *Store) DecCap() {
	s.capacity--
}

func (s *Store) ResetData() {
	s.data = make([]byte, s.capacity)
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
