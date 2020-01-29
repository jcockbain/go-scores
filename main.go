package main

import (
	"log"
	"net/http"
)

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

// GetPlayerScore gets a players score from an InMemoryStore
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
