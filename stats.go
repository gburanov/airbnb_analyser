package main

import (
	"fmt"
	"time"
)

func reportStatistics(w writer, f *flats) error {
	date := time.Now()
	w.Write(date.String())

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
			w.Write(str)

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
	w.Write(avPrice)
	w.Write(maxPrice)
	w.Write(minPrice)
	return nil
}
