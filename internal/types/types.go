package types

type Shuffler interface {
	Storer

	ShuffleMany(times int)
	Shuffle()
}

type Storer interface {
	Data() []rune
	Cap() int
	DecCap()
	ResetData()
	Push(index int)
	MoveSource()
	MoveData()
}
