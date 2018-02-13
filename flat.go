package main

import (
	"fmt"
	"strings"

	"github.com/Jeffail/gabs"
)

type flat struct {
	id        int
	price     float64
	name      string
	indexJSON *gabs.Container
	mainJSON  *gabs.Container
}

func (f *flat) isActive() (bool, error) {
	if f.mainJSON == nil {
		url := fmt.Sprintf("https://api.airbnb.com/v2/listings/%d?client_id=3092nxybyb0otqw18e8nh5nty&locale=en-US&currency=USD&_format=v1_legacy_for_p3&_source=mobile_p3&number_of_guests=1", f.id)
		jsonParsed, err := getURLWithRetries(url)
		if err != nil {
			return false, newSmartError(err, "")
		}
		f.mainJSON = jsonParsed
	}
	hasAvail := f.mainJSON.Search("listing", "has_availability").Data().(bool)
	reviewsCount := f.mainJSON.Search("listing", "reviews_count").Data().(float64)
	calendarUpdated := f.mainJSON.Search("listing", "calendar_updated_at").Data().(string)

	isActive := hasAvail && (reviewsCount > 2) && !strings.Contains(calendarUpdated, "month")
	return isActive, nil
}

func (f *flat) toStr() (string, error) {
	isActiv, err := f.isActive()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("URL: %v, %v, price %v EUR. Active status %v", f.url(), f.name, f.price, isActiv), nil
}

func (f *flat) url() string {
	return fmt.Sprintf("https://www.airbnb.com/rooms/%d", f.id)
}

func canInstantBook(f *flat) (bool, error) {
	ret := f.indexJSON.Search("pricing_quote", "can_instant_book").Data().(bool)
	return ret, nil
}

func lowCapacity(f *flat) (bool, error) {
	capacity := f.indexJSON.Search("listing", "person_capacity").Data().(float64)
	ret := capacity <= 2.5
	return ret, nil
}

func getFlat(c *gabs.Container) (flat, error) {
	flatInstance := flat{}
	flatInstance.id = int(c.Search("listing", "id").Data().(float64))
	flatInstance.price = c.Search("pricing_quote", "rate", "amount").Data().(float64)
	flatInstance.name = c.Search("listing", "name").Data().(string)
	flatInstance.indexJSON = c
	return flatInstance, nil
}
