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

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Empty dotenv")
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

	fmt.Printf("Rooms %v", rooms)
}
