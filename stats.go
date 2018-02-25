package main

import (
	"fmt"
	"time"
)

func reportStatistics(s *sheetClient, f *flats) error {
	date := time.Now()
	fmt.Println(date.String())
	s.write(date.String())

	prices := []float64{}
	for _, flat := range *f {
		isActive, err := flat.isActive()
		if err != nil {
			return err
		}

		if isActive {
			str, err := flat.toStr()
			if err != nil {
				return err
			}
			fmt.Println(str)
			s.write(str)

			prices = append(prices, flat.price)
		}
	}

	// get av price
	var total float64
	var min float64
	var max float64 = 120000
	for _, value := range prices {
		total += value
		if value > min {
			min = value
		}
		if value < max {
			max = value
		}
	}
	maxPrice := fmt.Sprintf("Maximum price %v", min)
	minPrice := fmt.Sprintf("Minimum price %v", max)
	avPrice := fmt.Sprintf("Average price %v", total/float64(len(prices)))
	fmt.Println(avPrice)
	s.write(avPrice)
	fmt.Println(maxPrice)
	s.write(maxPrice)
	fmt.Println(minPrice)
	s.write(minPrice)
	return nil
}
