package main

import (
	"errors"
	"net/http"
	"os"

	"google.golang.org/api/googleapi/transport"
	sheets "google.golang.org/api/sheets/v4"
)

func writeToSheet(flatDesc string) error {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		return errors.New("No Google APi key")
	}

	client := http.Client{Transport: &transport.APIKey{Key: apiKey}}

	srv, err := sheets.New(&client)
	if err != nil {
		return err
	}

	sheetID := os.Getenv("SPREADSHEET_ID")
	if sheetID == "" {
		return errors.New("No Spreadsheet key")
	}

	sheet := srv.Spreadsheets.Get(sheetID).Do()
	sheet.Fields()

	return nil
}
