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
	"time"
)

// this was v0.1 to test the api. Still exposed in cmdln because why not?
// I just use it to check tickers. E.g. will USO give me data?

// the global "structType" drives JsonToPostgres() and QueryBuilder() which are
var structType string // global, is specified in query builder,
// then used to marshal json, manage sql inserts...

// unmarshal into correct structs, because of the different nesting forms, can't avoid it
// to marshal into better format, use huge switch statement, which checks the global var structType
// is this (9/func) less loc than the requisite interface satisfying?
func DepricatedReformatJson(resp io.Reader) string {
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

// example querry:
// https://www.alphavantage.co//query?function=TIME_SERIES_DAILY&symbol=EWZ&apikey= examplekey'
// ticker's a string with spaces, check the first word in the switch statement (e.g. overview ewz -> OVERVIEW EWZ)
// this frist word picks the func/querry type
func DepricatedQueryBuilder(ticker string) (url string) {
	// the actual ticker comes after
	tickerFirst := strings.Fields(ticker)[0]
	var dateRegexIsTrue string
	if regexp.MustCompile(`\b\d{4}-\d{2}\b`).MatchString(tickerFirst) {
		dateRegexIsTrue = tickerFirst
	}

	switch tickerFirst {

	// News sentiment - complicated beast - figure out later

	// FX_DAILY
	case "EXCHANGE", "CURRENCY", "RATE":
		from := strings.Fields(ticker)[1]
		to := strings.Fields(ticker)[2]
		url = baseUrl + "FX_DAILY" + "&outputsize=full" + "&from_symbol=" + from + "&to_symbol=" + to
		structType = "APIs.ForexPrices"
		return
	// TOP_GAINERS_LOSERS and most active...
	// fix pls
	case "TGLAT", "TGLATS", "GAINERS", "LOSERS", "TOPGAINERSLOSERS":
		url = baseUrl + "TOP_GAINERS_LOSERS"
		structType = "APIs.TGLATs"
		return
	// overview OVERVIEW
	case "OVERVIEW":
		tickerNext := strings.Fields(ticker)[1]
		url = baseUrl + "OVERVIEW" + "&symbol=" + tickerNext
		structType = "APIs.StockOverview"
		return
	// income INCOME_STATEMENT  // "EARNINGS", "CASHFLOW", "BALANCE", "BALANCESHEET", "INCOME", "INCOMESTATEMENT"
	// balance 	BALANCE_SHEET
	case "INCOME_STATEMENT", "INCOME", "STATEMENT", "INCOMESTATEMENT":
		tickerNext := strings.Fields(ticker)[1]
		url = baseUrl + "INCOME_STATEMENT" + "&symbol=" + tickerNext
		structType = "APIs.IncomeStatements"
		return
	case "BALANCE_SHEET", "BALANCESHEET", "BALANCE":
		tickerNext := strings.Fields(ticker)[1]
		url = baseUrl + "BALANCE_SHEET" + "&symbol=" + tickerNext
		structType = "APIs.BalanceSheets"
		return
	// cashflow CASH_FLOW
	case "CASH_FLOW", "CASHFLOW":
		tickerNext := strings.Fields(ticker)[1]
		url = baseUrl + "CASH_FLOW" + "&symbol=" + tickerNext
		structType = "APIs.CashFlowStatements"
		return
	// earnings	EARNINGS
	case "EARNINGS":
		tickerNext := strings.Fields(ticker)[1]
		url = baseUrl + "EARNINGS" + "&symbol=" + tickerNext
		structType = "APIs.EarningsData"
		return

	// commodities and macro indicators use the same structs, but are funcs instead of ticker
	// WTI, BRENT, NATURAL_GAS, COPPER, ALUMINUM, WHEAT, CORN, COTTON, SUGAR, COFFEE, ALL_COMMODITIES
	// WTI
	case "WTI":
		url = baseUrl + "WTI" + "&interval=daily"
		structType = "APIs.CommodityPrices"
		return
	// BRENT
	case "BRENT":
		url = baseUrl + "BRENT" + "&interval=daily"
		structType = "APIs.CommodityPrices"
		return
	// nat gas
	case "NATURAL_GAS", "GAS": // check if GAS is a stock...
		url = baseUrl + "NATURAL_GAS" + "&interval=daily"
		structType = "APIs.CommodityPrices"
		return
	// COPPER
	case "COPPER":
		url = baseUrl + "COPPER" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	// ALUMINUM
	case "ALUMINUM":
		url = baseUrl + "ALUMINUM" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	// WHEAT
	case "WHEAT":
		url = baseUrl + "WHEAT" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	// CORN
	case "CORN":
		url = baseUrl + "CORN" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	// COTTON
	case "COTTON":
		url = baseUrl + "COTTON" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	// SUGAR
	case "SUGAR":
		url = baseUrl + "SUGAR" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	// COFFEE
	case "COFFEE":
		url = baseUrl + "COFFEE" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	// ALL_COMMODITIES
	case "ALL_COMMODITIES":
		url = baseUrl + "ALL_COMMODITIES" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	// REAL_GDP
	case "GDP", "REAL_GDP":
		url = baseUrl + "REAL_GDP" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	// real gdp per cap
	case "GDPPC", "GDPPERCAP", "REAL_GDP_PER_CAPITA":
		url = baseUrl + "REAL_GDP_PER_CAPITA" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	case "FEDFUNDSRATE", "FEDFUNDS", "FUNDS", "EFFECTIVEFEDERALFUNDSRATE", "EFFR", "FEDERAL_FUNDS_RATE":
		url = baseUrl + "FEDERAL_FUNDS_RATE" + "&interval=daily"
		structType = "APIs.CommodityPrices"
		return
	case "CPI":
		url = baseUrl + "CPI" + "&interval=monthly"
		structType = "APIs.CommodityPrices"
		return
	// inflation
	case "INFLATION":
		url = baseUrl + "INFLATION"
		structType = "APIs.CommodityPrices"
		return
	// retail sales - RETAIL_SALES
	case "RETAILSALES", "RETAIL", "RETAIL_SALES":
		url = baseUrl + "RETAIL_SALES"
		structType = "APIs.CommodityPrices"
		return
		// durable goods orders - DURABLES
	case "DURABLES":
		url = baseUrl + "DURABLES"
		structType = "APIs.CommodityPrices"
		return
	// unemployment - UNEMPLOYMENT
	case "UNEMPLOYMENT":
		url = baseUrl + "UNEMPLOYMENT"
		structType = "APIs.CommodityPrices"
		return
	// nonfarm payroll
	case "NONFARMPAYROLL", "NONFARM", "PAYROLL", "EMPLOYMENT", "NONFARM_PAYROLL":
		url = baseUrl + "NONFARM_PAYROLL"
		structType = "APIs.CommodityPrices"
		return

		// TREASURY_YIELD  &maturity 3month, 2year,5,year,7year, 10 year, 30year
	case "BOND", "YIELD", "TREASURY", "TREASURY_YIELD":
		// second field is maturity
		maturity := strings.Fields(ticker)[1]
		switch maturity {
		case "3", "3m", "3month":
			maturity = "3month"
		case "2", "2y", "2yr", "2year":
			maturity = "2year"
		case "5", "5y", "5yr", "5year":
			maturity = "5year"
		case "7", "7y", "7yr", "7year":
			maturity = "7year"
		case "10", "10y", "10yr", "10year":
			maturity = "10year"
		case "30", "30y", "30yr", "30year":
			maturity = "30year"
		}

		url = baseUrl + "TREASURY_YIELD" + "&interval=daily" + "&maturity=" + maturity
		structType = "APIs.CommodityPrices"
		return

	// time series intraday   	 ?interval=1min  extended true/false?
	// e.g. month=2009-01, since 2000-01
	case dateRegexIsTrue: // looks for 2001-01 format

		// Parse date to check if it's before "2000-01".
		date, err := time.Parse("2006-01", tickerFirst)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}
		refDate, _ := time.Parse("2006-01", "2000-01")
		if date.Before(refDate) {
			fmt.Println("Error: Date is before 2000-01")
		}
		tickerNext := strings.Fields(ticker)[1]
		url = baseUrl + "TIME_SERIES_INTRADAY" + "&month" + tickerFirst + "&interval=1min" + "&symbol=" + tickerNext + "&outputsize=full"
		structType = "APIs.IntradayOHLCVs"
		return

	// daily time series, DailyOHLCVs
	// &outputsize=full gets 20 years of data, remove it when testing defaulting to compact with 100 data points...
	default:
		url = baseUrl + "TIME_SERIES_DAILY" + "&symbol=" + ticker + "&outputsize=full"
		structType = "APIs.DailyOHLCVs"
		return

	}
}
