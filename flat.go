package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

type flat struct {
	id    int
	price float64
	name  string
	c     *gabs.Container
}

func (f *flat) toStr() string {
	return fmt.Sprintf("URL: %v, %v, price %v EUR", f.url(), f.name, f.price)
}

func (f *flat) url() string {
	return fmt.Sprintf("https://www.airbnb.com/rooms/%d", f.id)
}

func canInstantBook(f *flat) (bool, error) {
	ret := f.c.Search("pricing_quote", "can_instant_book").Data().(bool)
	return ret, nil
}

func lowCapacity(f *flat) (bool, error) {
	capacity := f.c.Search("listing", "person_capacity").Data().(float64)
	ret := capacity <= 2.5
	return ret, nil
}

func getFlat(c *gabs.Container) (flat, error) {
	flatInstance := flat{}
	flatInstance.id = int(c.Search("listing", "id").Data().(float64))
	flatInstance.price = c.Search("pricing_quote", "rate", "amount").Data().(float64)
	flatInstance.name = c.Search("listing", "name").Data().(string)
	flatInstance.c = c
	return flatInstance, nil
}
