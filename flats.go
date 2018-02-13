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

func selectFromMap(f *flats, chooser func(*flat) (bool, error)) (*flats, error) {
	ret := flats{}
	for k, v := range *f {
		good, err := chooser(&v)
		if err != nil {
			return nil, err
		}
		if good {
			ret[k] = v
		}
	}
	return &ret, nil
}

func getFlat(c *gabs.Container) (flat, error) {
	flatInstance := flat{}
	flatInstance.id = int(c.Search("listing", "id").Data().(float64))
	flatInstance.price = c.Search("pricing_quote", "rate", "amount").Data().(float64)
	flatInstance.c = c
	return flatInstance, nil
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
