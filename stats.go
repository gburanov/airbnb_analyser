package main

import "fmt"

func reportStatistics(f *flats) {
	for _, flat := range *f {
		fmt.Println(flat.toStr())
	}
}
