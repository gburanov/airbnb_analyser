package main

import (
	"github.com/Jeffail/gabs"
)

type flat struct {
	id    int
	price float64
	c     *gabs.Container
}

type flats map[int]flat

func (f flats) add(flatInstance flat) {
	f[flatInstance.id] = flatInstance
	return
}

func getFlat(c *gabs.Container) (flat, error) {
	flatInstance := flat{}
	flatInstance.id = int(c.Search("listing", "id").Data().(float64))
	flatInstance.price = c.Search("pricing_quote", "rate", "amount").Data().(float64)
	flatInstance.c = c
	return flatInstance, nil
}
