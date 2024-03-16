package types

type Shuffler interface {
	Storer

	ShuffleMany(times int)
	Shuffle()
}

type Storer interface {
	Data() []byte
	Cap() int
	DecCap()
	ResetData()
	Push(index int)
	MoveSource()
	MoveData()
}
