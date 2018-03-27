package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
	//log.SetLevel(log.InfoLevel)

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Empty dotenv")
	}
	sheetInst, err := newSheetClient()
	if err != nil {
		fmt.Printf("No output to google. %v\n", err)
	}

	lat := 52.573120
	lng := 13.355920
	guests := 1
	rooms, err := search(lat, lng, guests)
	if err != nil {
		log.Fatal(err)
	}
	rooms, err = selectFromMap(rooms, canInstantBook)
	if err != nil {
		log.Fatal(err)
	}
	rooms, err = selectFromMap(rooms, lowCapacity)
	if err != nil {
		log.Fatal(err)
	}
	if sheetInst != nil {
		err = reportStatistics(sheetInst, rooms)
	} else {
		w := outputWriter{}
		err = reportStatistics(&w, rooms)
	}
	if err != nil {
		log.Fatal(err)
	}

	return
}
