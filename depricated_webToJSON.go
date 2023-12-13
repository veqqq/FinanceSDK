package main

import (
	"FinanceSDK/APIs"
	"FinanceSDK/e"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// this was v0.1 to test the api. Still exposed in cmdln because why not?
// I just use it to check tickers. E.g. will USO give me data?

// unmarshal into correct structs, because of the different nesting forms, can't avoid it
// to marshal into better format, use huge switch statement, which checks the global var structType
// is this (9/func) less loc than the requisite interface satisfying?
func ReformatJson(resp io.Reader) string {
	// declare encoder
	decoder := json.NewDecoder(resp)
	// the switch checks global var structType, then uses it as the marshaling struct type
	switch structType {
	case "APIs.ForexPrices":
		var m APIs.ForexPrices
		err := decoder.Decode(&m)
		e.Check(err)
		output, err := json.Marshal(m.TimeSeriesFX) // perhaps change, but 3 maps
		e.Check(err)
		return string(output)
	case "APIs.TGLATs":
		var m APIs.TGLATs
		err := decoder.Decode(&m)
		e.Check(err)
		output, err := json.Marshal(m) // perhaps change, but 3 maps
		e.Check(err)
		return string(output)
	case "APIs.StockOverview":
		var m APIs.StockOverview
		err := decoder.Decode(&m)
		e.Check(err)
		output, err := json.Marshal(m)
		e.Check(err)
		return string(output)
	case "APIs.IncomeStatements":
		var m APIs.IncomeStatements
		err := decoder.Decode(&m)
		e.Check(err)
		output, err := json.Marshal(m.QuarterlyReports)
		e.Check(err)
		return string(output)
	case "APIs.BalanceSheets":
		var m APIs.BalanceSheets
		err := decoder.Decode(&m)
		e.Check(err)
		output, err := json.Marshal(m.QuarterlyReports)
		e.Check(err)
		return string(output)
	case "APIs.CashFlowStatements":
		var m APIs.CashFlowStatements
		err := decoder.Decode(&m)
		e.Check(err)
		output, err := json.Marshal(m.QuarterlyReports)
		e.Check(err)
		return string(output)
	case "APIs.EarningsData":
		var m APIs.EarningsData
		err := decoder.Decode(&m)
		e.Check(err)
		output, err := json.Marshal(m.QuarterlyEarnings)
		e.Check(err)
		return string(output)
	// Commodities and Economic Indicators - use same structure
	// WTI, BRENT, nat gas, COPPER, ALUMINUM, WHEAT, CORN, COTTON, SUGAR, COFFEE
	case "APIs.CommodityPrices":
		var m APIs.CommodityPrices
		err := decoder.Decode(&m)
		e.Check(err)
		output, err := json.Marshal(m.Data)
		e.Check(err)
		return string(output)
	case "APIs.IntradayOHLCVs":
		var m APIs.IntradayOHLCVs
		err := decoder.Decode(&m)
		e.Check(err)
		output, err := json.Marshal(m.TimeSeries1min)
		e.Check(err)
		return string(output)
	case "APIs.DailyOHLCVs":
		var m APIs.DailyOHLCVs
		err := decoder.Decode(&m)
		e.Check(err)
		output, err := json.Marshal(m.TimeSeries)
		e.Check(err)
		return string(output)
	default: // why do i need this? wont trigger, hm
		panic("confident I don't need this")
	}
}

func WriteToFile(filename, data string) {
	// if the file doesn't exist, make it, using ticker name (incl day/full querry)
	// n.b. will duplicate data if file exists - remove append or
	// have the sql builder have aditional logic?
	// Move the date to end
	words := strings.Fields(filename)
	for _, word := range words {
		if regexp.MustCompile(`\b\d{4}-\d{2}\b`).MatchString(word) { // date in 2003-01 format
			filename = strings.ReplaceAll(filename, word, "")
			filename = filename[1:] + " " + word // [1:] to remove a space, or: word = word + " "
		}
	}

	filename = strings.ReplaceAll(filename, " ", "_") // remove spaces from file name
	err := os.MkdirAll("data", 0755)                  // make folder if it doesn't exist
	e.Check(err)
	f, err := os.OpenFile("data/"+filename+".json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	e.Check(err)
	defer f.Close()
	fmt.Fprint(f, data)
}
