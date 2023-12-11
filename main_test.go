package main

import (
	"os"
	"testing"
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

// how to actually test for these cases?
// add more test cases... bonds etc. use curl and then the actual program output?
// don't test the json response, trust stdlib

// func TestWriteToFileAndQueryBuilder(t *testing.T) {
// 	testCases := []struct {
// 		ticker   string // WriteToFile requires a structTyle, which is determined in QueryBuilder
// 		input    string
// 		expected string
// 	}{
// 		{
// 			input: `{
// 				"Meta Data": {
// 					"1. Information": "Daily Prices (open, high, low, close) and Volumes",
// 					"2. Symbol": "FAKESTOCK",
// 					"3. Last Refreshed": "2023-08-04",
// 					"4. Output Size": "Compact",
// 					"5. Time Zone": "US/Eastern"
// 				},
// 				"Time Series (Daily)": {
// 					"2023-08-04": {
// 						"1. open": "32.5300",
// 						"2. high": "32.8550",
// 						"3. low": "32.0425",
// 						"4. close": "32.0600",
// 						"5. volume": "30396398"
// 					},
// 					"2023-08-03": {
// 						"1. open": "32.8000",
// 						"2. high": "33.0300",
// 						"3. low": "32.2700",
// 						"4. close": "32.2900",
// 						"5. volume": "27908785"
// 					}
// 				}
// 			}`,
// 			ticker:   "EWZ",
// 			expected: `{"2023-08-03":{"1. open":"32.8","2. high":"33.03","3. low":"32.27","4. close":"32.29","5. volume":"27908785"},"2023-08-04":{"1. open":"32.53","2. high":"32.855","3. low":"32.0425","4. close":"32.06","5. volume":"30396398"}}`,
// 		},
// 		{
// 			input: `CURRENCY EUR USD`,
// 			ticker: `{
// 					"Meta Data": {
// 						"1. Information": "Forex Daily Prices (open, high, low, close)",
// 						"2. From Symbol": "EUR",
// 						"3. To Symbol": "USD",
// 						"4. Output Size": "Full size",
// 						"5. Last Refreshed": "2023-08-15 03:00:00",
// 						"6. Time Zone": "UTC"
// 					},
// 					"Time Series FX (Daily)": {
// 						"2004-06-21": {
// 							"1. open": "1.21370",
// 							"2. high": "1.21440",
// 							"3. low": "1.20720",
// 							"4. close": "1.21080"
// 						},
// 						"2004-06-18": {
// 							"1. open": "1.20510",
// 							"2. high": "1.21440",
// 							"3. low": "1.19710",
// 							"4. close": "1.21370"
// 						},
// 						"2004-06-17": {
// 							"1. open": "1.20010",
// 							"2. high": "1.20760",
// 							"3. low": "1.19890",
// 							"4. close": "1.20610"
// 						},
// 						"2004-06-16": {
// 							"1. open": "1.21580",
// 							"2. high": "1.21650",
// 							"3. low": "1.19760",
// 							"4. close": "1.20060"
// 						},
// 						"2004-06-15": {
// 							"1. open": "1.20650",
// 							"2. high": "1.21650",
// 							"3. low": "1.20220",
// 							"4. close": "1.21580"
// 						}
// 					}
// 				}`,
// 			expected: `{"2004-06-15":{"1. open":"1.2065","2. high":"1.2165","3. low":"1.2022","4. close":"1.2158"},"2004-06-16":{"1. open":"1.2158","2. high":"1.2165","3. low":"1.1976","4. close":"1.2006"},"2004-06-17":{"1. open":"1.2001","2. high":"1.2076","3. low":"1.1989","4. close":"1.2061"},"2004-06-18":{"1. open":"1.2051","2. high":"1.2144","3. low":"1.1971","4. close":"1.2137"},"2004-06-21":{"1. open":"1.2137","2. high":"1.2144","3. low":"1.2072","4. close":"1.2108"}}`,
// 		},
// 		// {
// 		// 	input:    ``,
// 		// 	ticker:   "",
// 		// 	expected: ``,
// 	}

// 	for _, tc := range testCases {
// 		t.Run("Test case", func(t *testing.T) {
// 			// Convert the JSON string to an io.Reader
// 			reader := strings.NewReader(tc.input)

// 			_ = QueryBuilder(tc.ticker) // returns unneeded url

// 			// Call the WriteToFile function with the formatted JSON data
// 			// Delete test file after
// 			filename := "test_data"
// 			WriteToFile(filename, ReformatJson(reader))
// 			defer os.Remove("data/test_data.txt")

// 			// Read the written data from the file
// 			writtenData, err := os.ReadFile("data/" + filename + ".txt")
// 			if err != nil {
// 				t.Fatalf("Error reading written data from file: %v", err)
// 			}

// 			// Compare the written data with the expected output
// 			actual := string(writtenData)
// 			expected := tc.expected

// 			if actual != expected {
// 				t.Errorf("Actual output does not match expected output for test case.\nActual:\n%s\nExpected:\n%s", actual, expected)
// 			}
// 		})
// 	}
// }
