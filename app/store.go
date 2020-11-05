package app

type Storage interface {
	Insert(value int)
	Search(value int) bool
	Remove(value int)
}
