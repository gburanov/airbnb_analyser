package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	sheets "google.golang.org/api/sheets/v4"
)

func getHTTPClient() (*http.Client, error) {
	data, err := ioutil.ReadFile("token.json")
	if err != nil {
		return nil, err
	}
	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, err
	}
	client := conf.Client(oauth2.NoContext)
	return client, nil
}

func writeToSheet(flatDesc string) error {
	client, err := getHTTPClient()
	if err != nil {
		return err
	}

	srv, err := sheets.New(client)
	if err != nil {
		return err
	}

	sheetID := os.Getenv("SPREADSHEET_ID")
	if sheetID == "" {
		return errors.New("No Spreadsheet key")
	}

	s := []interface{}{"Hi there"}
	values := sheets.ValueRange{Values: [][]interface{}{s}}

	_, err = srv.Spreadsheets.Values.Append(sheetID, "A2:E", &values).
		ValueInputOption("USER_ENTERED").
		Do()
	if err != nil {
		return err
	}

	fmt.Println("DONE")
	return nil
}
