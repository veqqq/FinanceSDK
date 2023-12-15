package main

import (
	"FinanceSDK/APIs"
	"FinanceSDK/e"
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func GetSanatizedInput(msg string, args ...string) string {
	for {
		fmt.Println(msg)
		var userInput string
		fmt.Scan(&userInput)
		for _, arg := range args {
			if userInput == arg {
				return userInput
			}
		}
		fmt.Printf("Invalid input")
	}
}

// main logic/loop
func WhatDoesUserWantToDo() {
	var choice int

	// open DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo) // first arg is driver name (pq!)
	e.Check(err)

	fmt.Print("What do you want to do?\n1. Add a ticker\n2. Let DB Update\n3. Save ticker as local json (v0.1)\n4. Exit\n")
	fmt.Scan(&choice)

mainchoice:
	switch choice {
	case 1:
		var newtickers []Job
		for {
			ticker := GetTickerFromUser() // sanatize inputs, nothing wrong besides forex should get through
			fmt.Println(ticker)

			// n = only ohcvls, a = accounting docs and ohcvls, i = intraday ohcvls
			msg := "\nHow often is this needed?\n" +
				"q. Quarterly\nm. monthly\n"
			importance := GetSanatizedInput(msg, "m", "q")

			newticker := Job{
				TickerSymbol: ticker,
				Importance:   importance,
			}
			// todo #add "type"
			newtickers = append(newtickers, newticker)

			willUserContinue := GetSanatizedInput("Do you want to add more tickers? y or n\n", "y", "n")
			if willUserContinue == "n" {
				fmt.Println(newtickers)
				AddTickersToDB(db, newtickers)
				break mainchoice
			}
		}
	case 2:
		UpdateJobQueue(db)
		tickers, err := GetJobQueue(db)
		e.Check(err)

		for _, ticker := range tickers {
			thing := QueryBuilder(ticker.TickerSymbol) // thing is Container w Uploader
			req, _ := http.NewRequest("GET", thing.url, nil)
			req.Header.Add("X-RapidAPI-Key", apiKey)
			req.Header.Add("X-RapidAPI-Host", "alpha-vantage.p.rapidapi.com")

			resp, err := http.DefaultClient.Do(req)
			e.Check(err)
			defer resp.Body.Close()
			fmt.Println(thing.url+" %T", thing.giventype) // will it say uploader? Should be more precise #todo

			readResp, err := io.ReadAll(resp.Body)
			e.Check(err)
			thing.giventype.Upload(db, ticker.TickerSymbol, ticker.TickerID, readResp) // should wrap this, to pass thing

			UpdateLastUpdated(db, ticker.TickerSymbol)
			RemoveFromJobQueue(db, ticker.TickerSymbol)
			ReportApiCall(db)

			time.Sleep(6 * time.Second) // timer to not overload the api
		}
	case 3:
		// v0.1 functionality implemented in depricated.go
		ticker := GetTickerFromUser()
		url := DepricatedQueryBuilder(ticker)
		fmt.Println(url+" %T", url)

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("X-RapidAPI-Key", apiKey)
		req.Header.Add("X-RapidAPI-Host", "alpha-vantage.p.rapidapi.com")

		resp, err := http.DefaultClient.Do(req)
		e.Check(err)
		defer resp.Body.Close()
		finalData := DepricatedReformatJson(resp.Body)
		WriteToFile(ticker, finalData)
	case 4:
		fmt.Println("Exiting program.")
		os.Exit(0)
	default:
		fmt.Println("Input not allowed")
	}
}

// accepts ewz overview, overview ewz etx.
// if not given with script invocation, will ask
// accepts multiple args and turns them into one string
// makes all caps, puts relevant thingslike bond, overview etc. first
func GetTickerFromUser() string {
	fmt.Println("Enter Ticker. N.b. prefereds use a hyphen (PBR-A):")

	var userInput, arg1, arg2, arg3 string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	args := strings.Fields(line)
	if len(args) > 0 {
		arg1 = args[0]
	}
	if len(args) > 1 {
		arg2 = args[1]
	}
	if len(args) > 2 {
		arg3 = args[2]
	}
	userInput = arg1 + " " + arg2 + " " + arg3
	userInput = strings.ToUpper(userInput)
	// put overview, bond etc. as the first word before returning, so it can accept ewz overview, also.
	if strings.Contains(userInput, ".") {
		userInput = strings.ReplaceAll(userInput, ".", "-")
	}
	// looking for these, e.g. bond..
	words := []string{"NONFARMPAYROLL", "NONFARM", "PAYROLL", "EMPLOYMENT", "TGLAT", "GAINERS", "LOSERS", "TOPGAINERSLOSERS",
		"OVERVIEW", "RETAIL", "INFLATION", "CPI", "FEDFUNDSRATE", "FUNDS", "EFFECTIVEFEDERALFUNDSRATE", "EFFR",
		"GDPPC", "GDPPERCAP", "GDP", "EARNINGS", "CASHFLOW", "BALANCE_SHEET", "BALANCE", "BALANCESHEET",
		"INCOME", "INCOMESTATEMENT", "STATEMENT", "INCOME_STATEMENT", "RETAILSALES",
		"BOND", "YIELD", "TREASURY", "TREASURY_YIELD", "EXCHANGE", "CURRENCY", "RATE", "FX_DAILY", "FOREX",
	}
	for _, word := range words {
		if strings.Contains(userInput, word) {
			userInput = strings.ReplaceAll(userInput, word, "")
			userInput = word + " " + userInput
		}
	}
	// Move the date to beginning
	words = strings.Fields(userInput)
	for _, word := range words {
		if regexp.MustCompile(`\b\d{4}-\d{2}\b`).MatchString(word) { // date in 2003-01 format
			userInput = strings.ReplaceAll(userInput, word, "")
			userInput = word + " " + userInput
		}
	}
	userInput = strings.ReplaceAll(userInput, "  ", " ")
	userInput = strings.TrimSuffix(userInput, " ")
	return userInput
}

type Uploader interface {
	Upload(db *sql.DB, ticker string, ID int, body []byte)
}

type TypeContainer struct {
	url       string
	giventype Uploader
}

// this sanatizes input in the cli, but also chooses types for moving data around
// nearly a god function
// example querry:
// https://alpha-vantage.p.rapidapi.com/query?function=TIME_SERIES_DAILY&symbol=EWZ - the apikey goes in a header
// ticker's a string with spaces, check the first word in the switch statement (e.g. overview ewz -> OVERVIEW EWZ)
// this frist word picks the func/querry type
func QueryBuilder(ticker string) (result TypeContainer) {
	// the actual ticker comes after
	tickerFirst := strings.Fields(ticker)[0]
	var dateRegexIsTrue string
	if regexp.MustCompile(`\b\d{4}-\d{2}\b`).MatchString(tickerFirst) {
		dateRegexIsTrue = tickerFirst
	}

	switch tickerFirst {
	// News sentiment - complicated beast - figure out later
	// FX_DAILY
	case "EXCHANGE", "CURRENCY", "RATE", "FX_DAILY", "FOREX": // #todo need to implement ForexPrices uploader...
		from := strings.Fields(ticker)[1]
		to := strings.Fields(ticker)[2]
		result.url = baseUrl + "FX_DAILY" + "&outputsize=full" + "&from_symbol=" + from + "&to_symbol=" + to
		result.giventype = APIs.ForexPrices{}
		return
	// // TOP_GAINERS_LOSERS and most active... // removed
	// case "TGLAT", "TGLATS", "GAINERS", "LOSERS", "TOPGAINERSLOSERS":
	// 	url = baseUrl + "TOP_GAINERS_LOSERS"
	// 	result.giventype = "APIs.TGLATs"
	// 	return
	case "OVERVIEW":
		tickerNext := strings.Fields(ticker)[1]
		result.url = baseUrl + "OVERVIEW" + "&symbol=" + tickerNext
		result.giventype = APIs.StockOverview{}
		return
	// income INCOME_STATEMENT  // "EARNINGS", "CASHFLOW", "BALANCE", "BALANCESHEET", "INCOME", "INCOMESTATEMENT"
	case "INCOME_STATEMENT", "INCOME", "STATEMENT", "INCOMESTATEMENT":
		tickerNext := strings.Fields(ticker)[1]
		result.url = baseUrl + "INCOME_STATEMENT" + "&symbol=" + tickerNext
		result.giventype = APIs.IncomeStatements{}
		return
	case "BALANCE_SHEET", "BALANCESHEET", "BALANCE":
		tickerNext := strings.Fields(ticker)[1]
		result.url = baseUrl + "BALANCE_SHEET" + "&symbol=" + tickerNext
		result.giventype = APIs.BalanceSheets{}
		return
	case "CASH_FLOW", "CASHFLOW":
		tickerNext := strings.Fields(ticker)[1]
		result.url = baseUrl + "CASH_FLOW" + "&symbol=" + tickerNext
		result.giventype = APIs.CashFlowStatements{}
		return
	// case "EARNINGS": #did not put earnings data in sql, so no uploader implemented
	// 	tickerNext := strings.Fields(ticker)[1]
	// 	result.url = baseUrl + "EARNINGS" + "&symbol=" + tickerNext
	// 	result.giventype = APIs.EarningsData{}
	// 	return
	// commodities and macro indicators use the same structs, but are funcs instead of ticker
	// WTI, BRENT, NATURAL_GAS, COPPER, ALUMINUM, WHEAT, CORN, COTTON, SUGAR, COFFEE, ALL_COMMODITIES
	case "WTI":
		result.url = baseUrl + "WTI" + "&interval=daily"
		result.giventype = APIs.CommodityPrices{}
		return
	case "BRENT":
		result.url = baseUrl + "BRENT" + "&interval=daily"
		result.giventype = APIs.CommodityPrices{}
		return
	case "NATURAL_GAS", "GAS":
		result.url = baseUrl + "NATURAL_GAS" + "&interval=daily"
		result.giventype = APIs.CommodityPrices{}
		return
	case "COPPER":
		result.url = baseUrl + "COPPER" + "&interval=quarterly"
		result.giventype = APIs.CommodityPrices{}
		return
	case "ALUMINUM":
		result.url = baseUrl + "ALUMINUM" + "&interval=quarterly"
		result.giventype = APIs.CommodityPrices{}
		return
	case "WHEAT":
		result.url = baseUrl + "WHEAT" + "&interval=quarterly"
		result.giventype = APIs.CommodityPrices{}
		return
	case "CORN":
		result.url = baseUrl + "CORN" + "&interval=quarterly"
		result.giventype = APIs.CommodityPrices{}
		return
	case "COTTON":
		result.url = baseUrl + "COTTON" + "&interval=quarterly"
		result.giventype = APIs.CommodityPrices{}
		return
	case "SUGAR":
		result.url = baseUrl + "SUGAR" + "&interval=quarterly"
		result.giventype = APIs.CommodityPrices{}
		return
	case "COFFEE":
		result.url = baseUrl + "COFFEE" + "&interval=quarterly"
		result.giventype = APIs.CommodityPrices{}
		return
	case "ALL_COMMODITIES":
		result.url = baseUrl + "ALL_COMMODITIES" + "&interval=quarterly"
		result.giventype = APIs.CommodityPrices{}
		return
	case "GDP", "REAL_GDP":
		result.url = baseUrl + "REAL_GDP" + "&interval=quarterly"
		result.giventype = APIs.CommodityPrices{}
		return
	case "GDPPC", "GDPPERCAP", "REAL_GDP_PER_CAPITA":
		result.url = baseUrl + "REAL_GDP_PER_CAPITA" + "&interval=quarterly"
		result.giventype = APIs.CommodityPrices{}
		return
	case "FEDFUNDSRATE", "FEDFUNDS", "FUNDS", "EFFECTIVEFEDERALFUNDSRATE", "EFFR", "FEDERAL_FUNDS_RATE":
		result.url = baseUrl + "FEDERAL_FUNDS_RATE" + "&interval=daily"
		result.giventype = APIs.CommodityPrices{}
		return
	case "CPI":
		result.url = baseUrl + "CPI" + "&interval=monthly"
		result.giventype = APIs.CommodityPrices{}
		return
	case "INFLATION":
		result.url = baseUrl + "INFLATION"
		result.giventype = APIs.CommodityPrices{}
		return
	case "RETAILSALES", "RETAIL", "RETAIL_SALES":
		result.url = baseUrl + "RETAIL_SALES"
		result.giventype = APIs.CommodityPrices{}
		return
	case "DURABLES":
		result.url = baseUrl + "DURABLES"
		result.giventype = APIs.CommodityPrices{}
		return
	case "UNEMPLOYMENT":
		result.url = baseUrl + "UNEMPLOYMENT"
		result.giventype = APIs.CommodityPrices{}
		return
	case "NONFARMPAYROLL", "NONFARM", "PAYROLL", "EMPLOYMENT", "NONFARM_PAYROLL":
		result.url = baseUrl + "NONFARM_PAYROLL"
		result.giventype = APIs.CommodityPrices{}
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
		result.url = baseUrl + "TREASURY_YIELD" + "&interval=daily" + "&maturity=" + maturity
		result.giventype = APIs.CommodityPrices{}
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
		refDate, _ := time.Parse("2006-01", "2000-01") // #todo i suspect this will add itself as a new ticker like "CLF month"...
		if date.Before(refDate) {
			fmt.Println("Error: Date is before 2000-01")
		}
		tickerNext := strings.Fields(ticker)[1]
		result.url = baseUrl + "TIME_SERIES_INTRADAY" + "&month" + tickerFirst + "&interval=1min" + "&symbol=" + tickerNext + "&outputsize=full"
		result.giventype = APIs.IntradayOHLCVs{}
		return
	// daily time series, DailyOHLCVs
	// &outputsize=full gets 20 years of data, remove it when testing defaulting to compact with 100 data points...
	default:
		result.url = baseUrl + "TIME_SERIES_DAILY" + "&symbol=" + ticker + "&outputsize=full"
		result.giventype = APIs.DailyOHLCVs{}
		return
	}
}
