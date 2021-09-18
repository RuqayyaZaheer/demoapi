package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/gitgo/api/nhlApi"
)

func main() {
	now := time.Now()
	rstFile, err := os.OpenFile("listOfTeams.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error to open file %v", err)
	}
	defer rstFile.Close()

	wrt := io.MultiWriter(os.Stdout, rstFile)

	log.SetOutput(wrt)

	teams, err := nhlApi.GetAllTeams()

	if err != nil {
		log.Fatalf("Error while getting all teams: %v", err)

	}

	for _, team := range teams {
		log.Println("___________")
		log.Printf("Team Name: %s", team.Name)
		log.Printf("Team Name: %s", team.venue)

		log.Println("___________")
	}

	log.Printf("Took time to Complete %v", time.Now().Sub(now).String())

}
