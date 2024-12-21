package main

import (
	s "go-application/server"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	//In Memory store
	//server := s.NewPlayerServer(&s.InMomoryStore{})

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	store, err := s.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	game := s.NewTexasHoldem(s.BlindAlerterFunc(s.Alerter), store)
	server, err := s.NewPlayerServer(store, game)
	if err != nil {
		log.Fatalf("error while starging server %v", err)
	}
	log.Fatal(http.ListenAndServe(":5000", server))
}
