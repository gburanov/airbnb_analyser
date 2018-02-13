package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Jeffail/gabs"
	log "github.com/sirupsen/logrus"
)

func searchPage(lat float64, lng float64, guests int, offset int) (*flats, error) {
	client := http.Client{}
	url := fmt.Sprintf("https://api.airbnb.com/v2/search_results?client_id=3092nxybyb0otqw18e8nh5nty&locale=en-US&currency=EUR&_format=for_search_results_with_minimal_pricing&_limit=10&_offset=%d&fetch_facets=false&guests=%d&ib=false&ib_add_photo_flow=false&location=Reinichendort&min_bathrooms=1&min_bedrooms=1&min_beds=1&min_num_pic_urls=10&sort=1&user_lat=%v&user_lng=%v", offset, guests, lat, lng)

	log.WithFields(log.Fields{"url": url}).Info("Calling url")

	response, err := client.Get(url)
	if err != nil {
		return nil, smartError(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, smartError(err)
	}

	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		return nil, smartError(err)
	}

	// first of all, check for error
	error := jsonParsed.Search("error_message").Data()
	if error != nil {
		return nil, smartError(errors.New(error.(string)))
	}

	ret := flats{}
	flats, err := jsonParsed.Search("search_results").Children()
	if err != nil {
		return nil, smartError(err)
	}
	for _, flatInstance := range flats {
		var singleFlat flat
		singleFlat, err = getFlat(flatInstance)
		if err != nil {
			return nil, smartError(err)
		}
		ret.add(singleFlat)
	}

	return &ret, nil
}

func search(lat float64, lng float64, guests int) (*flats, error) {
	ret := flats{}
	for offset := 0; offset <= 130; offset += 10 {
		inter, err := searchPage(lat, lng, guests, offset)
		if err != nil {
			return nil, err
		}
		for k, v := range *inter {
			ret[k] = v
		}
	}
	return &ret, nil
}
