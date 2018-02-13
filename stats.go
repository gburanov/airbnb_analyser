package main

import "fmt"

func reportStatistics(f *flats) error {
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

			prices = append(prices, flat.price)
		}
	}

	// get av price
	var total float64
	for _, value := range prices {
		total += value
	}
	fmt.Printf("Average price %v", total/float64(len(prices)))
	return nil
}
