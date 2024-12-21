package main

import (
	"fmt"
	s "go-application/server"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := s.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	alerter := s.BlindAlerterFunc(s.Alerter)
	game := s.NewTexasHoldem(alerter, store)
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	s.NewCLI(os.Stdin, os.Stdout, game).PlayPoker()
}
