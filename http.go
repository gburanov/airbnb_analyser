package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/Jeffail/gabs"
	log "github.com/sirupsen/logrus"
)

func getURL(url string) (*gabs.Container, *smartError) {
	client := http.Client{}
	log.WithFields(log.Fields{"url": url}).Info("Calling url")

	response, err := client.Get(url)
	if err != nil {
		return nil, newSmartError(err, "")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, newSmartError(err, "")
	}

	log.WithFields(log.Fields{"body": string(body)}).Info("Response")

	if strings.Contains(string(body), "503 Service Temporarily Unavailable") {
		return nil, newSmartError(errors.New("service_unavailable"), "service_unavailable")
	}

	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		return nil, newSmartError(err, "")
	}

	// first of all, check for error
	error := jsonParsed.Search("error").Data()
	if error != nil {
		return nil, newSmartError(errors.New(error.(string)), "service_unavailable")
	}

	return jsonParsed, nil
}

func getURLWithRetries(url string) (*gabs.Container, error) {
	for {
		json, err := getURL(url)
		if err == nil {
			return json, nil
		}
		if err != nil {
			if err.getType() == "service_unavailable" {
				time.Sleep(time.Second)
				continue
			}
			return nil, err
		}
	}
}
