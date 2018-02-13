package main

import (
	"fmt"
)

func getMaxOffet(lat float64, lng float64, guests int) (int, *smartError) {
	url := fmt.Sprintf("https://api.airbnb.com/v2/search_results?client_id=3092nxybyb0otqw18e8nh5nty&locale=en-US&currency=EUR&_format=for_search_results_with_minimal_pricing&_limit=10&_offset=0&fetch_facets=false&guests=%d&ib=false&ib_add_photo_flow=false&location=Reinichendort&min_bathrooms=1&min_bedrooms=1&min_beds=1&min_num_pic_urls=10&sort=1&user_lat=%v&user_lng=%v", guests, lat, lng)
	jsonParsed, err := getURLWithRetries(url)
	if err != nil {
		return 0, newSmartError(err, "")
	}
	return int(jsonParsed.Search("metadata", "listings_count").Data().(float64)), nil
}

func searchPage(lat float64, lng float64, guests int, offset int) (*flats, *smartError) {
	url := fmt.Sprintf("https://api.airbnb.com/v2/search_results?client_id=3092nxybyb0otqw18e8nh5nty&locale=en-US&currency=EUR&_format=for_search_results_with_minimal_pricing&_limit=10&_offset=%d&fetch_facets=false&guests=%d&ib=false&ib_add_photo_flow=false&location=Reinichendort&min_bathrooms=1&min_bedrooms=1&min_beds=1&min_num_pic_urls=10&sort=1&user_lat=%v&user_lng=%v", offset, guests, lat, lng)
	jsonParsed, err := getURLWithRetries(url)
	if err != nil {
		return nil, newSmartError(err, "")
	}

	ret := flats{}
	flats, err := jsonParsed.Search("search_results").Children()
	if err != nil {
		return nil, newSmartError(err, "")
	}
	for _, flatInstance := range flats {
		var singleFlat flat
		singleFlat, err = getFlat(flatInstance)
		if err != nil {
			return nil, newSmartError(err, "")
		}
		ret.add(singleFlat)
	}

	return &ret, nil
}

func search(lat float64, lng float64, guests int) (*flats, error) {
	maxOffset, err := getMaxOffet(lat, lng, guests)
	if err != nil {
		return nil, err
	}

	ret := flats{}
	for offset := 0; offset <= maxOffset; offset += 10 {
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
