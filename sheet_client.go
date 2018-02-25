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

type sheetClient struct {
	position int
	client   *sheets.Service
	sheetID  string
}

func newSheetClient() (*sheetClient, error) {
	client, err := getHTTPClient()
	if err != nil {
		return nil, err
	}

	srv, err := sheets.New(client)
	if err != nil {
		return nil, err
	}

	sheetID := os.Getenv("SPREADSHEET_ID")
	if sheetID == "" {
		return nil, errors.New("No Spreadsheet key")
	}

	sheetClientInst := sheetClient{
		position: 2,
		client:   srv,
		sheetID:  sheetID,
	}

	return &sheetClientInst, nil
}

func (c *sheetClient) pos() string {
	return fmt.Sprintf("A%d", c.position)
}

func (c *sheetClient) write(str string) error {
	s := []interface{}{str}
	values := sheets.ValueRange{Values: [][]interface{}{s}}

	_, err := c.client.Spreadsheets.Values.Append(c.sheetID, c.pos(), &values).
		ValueInputOption("USER_ENTERED").
		Do()
	if err != nil {
		return err
	}
	c.position = c.position + 1

	return nil
}
