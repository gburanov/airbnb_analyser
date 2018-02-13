package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Empty dotenv")
	}
	lat := 52.573120
	lng := 13.355920
	guests := 1

	rooms, err := search(lat, lng, guests)
	if err != nil {
		fmt.Println("Empty dotenv")
	}
	fmt.Printf("Rooms %v", rooms)
}
