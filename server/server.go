package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) (score int)
	RecordWin(name string)
}

type PlayerServer struct {
	Store PlayerStore
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (ps *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		ps.processWin(w, player)
	case http.MethodGet:
		ps.showScore(w, player)
	}
}

func (ps *PlayerServer) processWin(w http.ResponseWriter, player string) {
	ps.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
	w.WriteHeader(http.StatusAccepted)
}

func (ps *PlayerServer) showScore(w http.ResponseWriter, player string) {

	score := ps.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	score := s.scores[player]
	return score
}
func (s *StubPlayerStore) RecordWin(player string) {
	s.winCalls = append(s.winCalls, player)
}
