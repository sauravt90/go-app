package server

func NewInMemoryPlayerStore() *InMomoryStore {
	return &InMomoryStore{
		scores: map[string]int{},
	}
}

type InMomoryStore struct {
	scores map[string]int
}

func (i *InMomoryStore) GetPlayerScore(name string) int {
	return i.scores[name]
}

func (i *InMomoryStore) RecordWin(name string) {
	i.scores[name]++
}

func (i *InMomoryStore) GetLeague() League {
	var players []Player
	for key, val := range i.scores {
		players = append(players, Player{key, val})
	}
	return players
}
