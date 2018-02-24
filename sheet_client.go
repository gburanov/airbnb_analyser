package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"google.golang.org/api/googleapi/transport"
	sheets "google.golang.org/api/sheets/v4"
)

func getHTTPClient() (*http.Client, error) {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		return nil, errors.New("No Google APi key")
	}

	client := http.Client{Transport: &transport.APIKey{Key: apiKey}}
	return &client, nil
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

	sheet, err := srv.Spreadsheets.Get(sheetID).Do()
	if err != nil {
		return err
	}

	newSheet, err := srv.Spreadsheets.Create(sheet).Do()
	if err != nil {
		return err
	}
	fmt.Println("New sheet %v", newSheet)
	return nil
}
