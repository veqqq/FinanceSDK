package main

import (
	"FinanceSDK/e"
	"net/http"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// tests to add:

// query to each used API
// // for breakage
// bond rate date, stock date, comparing to saved result
//    	also compare to somewhat cleaned json
// specific stock date - in DB - full integration test
// 			- use https://github.com/ory/dockertest to spin up DB!
//			- use actual prod DB? lol
//			- if DB logic allows it, in memory data store to replicate actions
//			similar to testing writers with bytes.Buffer
// // fundementals

// unit tests
// example info to DB
// query builder - TestGetTickerFromUser

func TestGetTickerFromUser(t *testing.T) {
	// Define a test table with input values and expected results
	testCases := []struct {
		input    string
		expected string
	}{
		{"pbr.a", "PBR-A"},
		{"ewz", "EWZ"},
		{"GOOG", "GOOG"},
		{"bDorY", "BDORY"},
		{"brent", "BRENT"},
		{"overview clf", "OVERVIEW CLF"},
		{"clf overview", "OVERVIEW CLF"},
		{"5 bond", "BOND 5"},
		{"5 yield", "YIELD 5"},
		{"yield 5", "YIELD 5"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			// Replace os.Args with the input value for testing, then reset
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()
			os.Args = []string{"test_prog", tc.input}

			result := GetTickerFromUser()
			if result != tc.expected {
				t.Errorf("Input: %s, Expected: %s, Got: %s", tc.input, tc.expected, result)
			}
		})
	}
}
func TestJsonToPostgres(t *testing.T) {
	testCases := []struct {
		ticker   string
		url      string
		testtype string
		expected string
	}{
		{"IBM",
			"https://www.alphavantage.co/query?function=OVERVIEW&symbol=IBM&apikey=demo",
			"APIs.StockOverview",
			"nope", // #todo fix this my god, make DB fake...
		}}

	for _, tc := range testCases {
		t.Run(tc.ticker, func(t *testing.T) {
			structType = tc.testtype
			resp, err := http.Get(tc.url)
			e.Check(err)

			// Create a mock database connection
			db, mock, err := sqlmock.New()
			e.Check(err)
			defer db.Close()

			// Set up expectations for the mock database
			mock.ExpectExec("INSERT INTO stock_overviews").WillReturnResult(sqlmock.NewResult(1, 1))

			JsonToPostgres(db, tc.ticker, resp.Body)

			// Check if all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("unfulfilled expectations: %s", err)
			}
		})
	}
}

// --------------------
