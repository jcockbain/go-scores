package main

// NewInMemoryPlayerStore creates a new InMemoryStore
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// InMemoryPlayerStore maps a player to their score
type InMemoryPlayerStore struct {
	store map[string]int
}

// RecordWin records a win to an InMemoryStore
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

// GetPlayerScore gets a players score from an
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}
