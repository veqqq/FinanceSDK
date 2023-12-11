package main

import (
	"FinanceSDK/APIs"
	"FinanceSDK/e"
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

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
		"BOND", "YIELD", "TREASURY", "TREASURY_YIELD", "EXCHANGE", "CURRENCY", "RATE",
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

// build a base url, fetching the apikey from .env file
// lazy implementation from godotenv to reduce dependencies
func buildBaseURL() string {
	f, err := (os.Open("e/.env")) // e for err, but hide from docker
	e.Check(err)
	defer f.Close()

	var envMap map[string]string
	err = json.NewDecoder(f).Decode(&envMap)
	e.Check(err)

	currentEnv := map[string]bool{}
	rawEnv := os.Environ()
	for _, rawEnvLine := range rawEnv {
		key := strings.Split(rawEnvLine, "=")[0]
		currentEnv[key] = true
	}
	for key, value := range envMap {
		if !currentEnv[key] {
			_ = os.Setenv(key, value)
		}
	}

	apiKey := os.Getenv("APIKEY")
	// apiKey, ok := os.LookupEnv("APIKEY")
	// if !ok {
	// 	log.Fatalf("Add API Key to .env")
	// }
	apiKey = "?apikey=" + apiKey
	return "https://www.alphavantage.co//query" + apiKey + "&function="
	// #todo currently coupled to alphavantage

}

// example querry:
// https://www.alphavantage.co//query?function=TIME_SERIES_DAILY&symbol=EWZ&apikey= examplekey'
// ticker's a string with spaces, check the first word in the switch statement (e.g. overview ewz -> OVERVIEW EWZ)
// this frist word picks the func/querry type
func QueryBuilder(ticker string) (url string) {
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
		structType = "CommodityPrices"
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
	case "GDP":
		url = baseUrl + "REAL_GDP" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	// real gdp per cap
	case "GDPPC", "GDPPERCAP":
		url = baseUrl + "REAL_GDP_PER_CAPITA" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	case "FEDFUNDSRATE", "FEDFUNDS", "FUNDS", "EFFECTIVEFEDERALFUNDSRATE", "EFFR":
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
	case "RETAILSALES", "RETAIL":
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
	case "NONFARMPAYROLL", "NONFARM", "PAYROLL", "EMPLOYMENT":
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

// unmarshal into correct structs, because of the different nesting forms, can't avoid it
// to marshal into better format, use huge switch statement, which checks the global var structType
// is this (9/func) less loc than the requisite interface satisfying?

// what if I don't remarshal?

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

// original
// func ReformatJson(body io.Reader, structType string) string {
// trying without structs, since I don't know how to declare the structType in accordance to the querry
// 	// use structType to pick structs
// 	// might need to make funcs to marshall and unmarshal for other programs later

// 	// structType will be the var instead of DailyOHLCVs, picking it...
// 	// declare m in switch cases!
// 	// but there will be scope issues, so predeclare them globally
// 	// or more elegantly, use interfaces...
// 	err := json.NewDecoder(body).Decode(&m)
// 	e.Check(err)

// 	// isolate just the time data
// 	jsondata, err := json.Marshal(m.TimeSeries) // TimeSeries must be modular too...
// 	e.Check(err)
// 	return string(jsondata) // return a pointer? these are big items...
// }

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

var baseUrl string    // global
var structType string // global, is specified in query builder, then used it reformat Json

func init() {
	baseUrl = buildBaseURL()
}

func WhatDoesUserWantToDo() {
	var choice int

	// open DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo) // first arg is driver name (pq!)
	e.Check(err)

	fmt.Print("What do you want to do?\n1. Add a ticker\n2. Let DB Update\n3. Save ticker as local json (v0.1) - Requires /data folder\n4. Exit\n")
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

			newtickers = append(newtickers, newticker)

			willUserContinue := GetSanatizedInput("Do you want to add more tickers? y or n\n", "y", "n")
			if willUserContinue == "n" {
				fmt.Println(newtickers)
				AddTickersToDB(db, newtickers)
				break mainchoice
			}
		}

	case 2:
		// build jobqueue
		// check if last updated is null
		// check if last updated is further back than the importance period

		tickers, err := GetJobQueue(db) // implement
		e.Check(err)
		fmt.Print(tickers)
		for _, ticker := range tickers {
			url := QueryBuilder(ticker.TickerSymbol)
			resp, err := http.Get(url)
			e.Check(err)
			defer resp.Body.Close()

			err = db.Ping()
			e.Check(err)
			JsonToPostgres(db, ticker.TickerSymbol, resp.Body)
		}
	case 3:
		// v0.1 functionality
		ticker := GetTickerFromUser()
		url := QueryBuilder(ticker)
		fmt.Println(url)

		resp, err := http.Get(url)
		e.Check(err)
		defer resp.Body.Close()
		finalData := ReformatJson(resp.Body)
		WriteToFile(ticker, finalData)
	case 4:
		fmt.Println("Exiting program.")
		os.Exit(0)
	default:
		fmt.Println("Input not allowed")
	}
}

func main() {

	// open DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo) // first arg is driver name (pq!)
	e.Check(err)

	structType = "APIs.StockOverview"
	url := "https://www.alphavantage.co/query?function=OVERVIEW&symbol=IBM&apikey=demo"
	fmt.Print(url)
	resp, _ := http.Get(url)

	JsonToPostgres(db, "IBM", resp.Body)

	// for {
	// 	WhatDoesUserWantToDo()
	// }
}

// DB stuff
const (
	host     = "127.0.0.1" // "jsontosql_db_1" // "127.0.0.1" // use docker name if from docker, ip if not in container!
	port     = 5432
	user     = "postgres"
	password = "password2"
	dbname   = "financial_markets"
)

type Job struct {
	TickerSymbol string
	Importance   string
}

func GetJobQueue(db *sql.DB) ([]Job, error) {
	rows, err := db.Query("SELECT TickerSymbol, importance FROM tickers")
	e.Check(err)
	defer rows.Close()

	var jobQueue []Job

	for rows.Next() {
		var job Job
		if err := rows.Scan(&job.TickerSymbol, &job.Importance); err != nil {
			return nil, err
		}
		jobQueue = append(jobQueue, job)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return jobQueue, nil
}

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

func AddTickersToDB(db *sql.DB, jobs []Job) {
	err := db.Ping()
	e.Check(err)

	for _, job := range jobs {
		_, err = db.Exec(`INSERT INTO tickers (TickerSymbol, Importance)
	VALUES ($1, $2)`,
			job.TickerSymbol, job.Importance)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				fmt.Println("Duplicate value not saved")
			} else {
				panic(err)
			}
		}
	}
}

func JsonToPostgres(db *sql.DB, ticker string, resp io.Reader) { // globalvar structtype will hold the type
	decoder := json.NewDecoder(resp)

	// _, err = db.Exec(`INSERT INTO tickers (TickerSymbol, Importance)
	// VALUES ($1, $2)`,
	// 		job.TickerSymbol, job.Importance)
	// 	if err != nil {
	// 		if strings.Contains(err.Error(), "duplicate") {
	// 			fmt.Println("Duplicate value not saved")
	// 		} else { panic(err)

	// the switch checks global var structType
	switch structType {
	// case "APIs.ForexPrices":
	// 	var m APIs.ForexPrices
	// 	err := decoder.Decode(&m)
	// 	e.Check(err) #todo how to know which currency pair? Add it to from and to
	// _, err = db.Exec(`INSERT INTO forex (fromCurrency, toCurrency, open, high, low, close, datasource)
	// VALUES ($1, $2, $3, $4, $5, $6, $7)`, from?, to?, APIs.ForexPrice.Open, APIs.ForexPrice.High, APIs.ForexPrice.Low, APIs.ForexPrice.Close, "Alphavantage")
	case "APIs.StockOverview":
		var m APIs.StockOverview
		err := decoder.Decode(&m)
		e.Check(err)

		tickerID := GetTickerID(db, ticker)

		_, err = db.Exec(`INSERT INTO stock_overviews (id,
        symbol, asset_type, name, cik, exchange, currency, country, sector, industry,
        address, fiscal_year_end, latest_quarter, market_capitalization, ebitda, pe_ratio, peg_ratio,
        book_value, dividend_per_share, dividend_yield, eps, revenue_per_share_ttm, profit_margin,
        operating_margin_ttm, return_on_assets_ttm, return_on_equity_ttm, revenue_ttm, gross_profit_ttm,
        diluted_eps_ttm, quarterly_earnings_growth_yoy, quarterly_revenue_growth_yoy,
        analyst_target_price, trailing_pe, forward_pe, price_to_sales_ratio_ttm, price_to_book_ratio,
        ev_to_revenue, ev_to_ebitda, beta, day_moving_average_50,
        day_moving_average_200, shares_outstanding, dividend_date, ex_dividend_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
        $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38,
        $39, $40, $41, $42, $43, $44)`,
			tickerID, m.Symbol, m.AssetType, m.Name,
			m.CIK, m.Exchange, m.Currency, m.Country,
			m.Sector, m.Industry, m.Address, m.FiscalYearEnd,
			m.LatestQuarter, m.MarketCapitalization, m.EBITDA,
			m.PERatio, m.PEGRatio, m.BookValue, m.DividendPerShare,
			m.DividendYield, m.EPS, m.RevenuePerShareTTM,
			m.ProfitMargin, m.OperatingMarginTTM, m.ReturnOnAssetsTTM,
			m.ReturnOnEquityTTM, m.RevenueTTM, m.GrossProfitTTM,
			m.DilutedEPSTTM, m.QuarterlyEarningsGrowthYOY,
			m.QuarterlyRevenueGrowthYOY, m.AnalystTargetPrice,
			m.TrailingPE, m.ForwardPE, m.PriceToSalesRatioTTM,
			m.PriceToBookRatio, m.EVToRevenue, m.EVToEBITDA,
			m.Beta, m.DayMovingAverage50, m.DayMovingAverage200,
			m.SharesOutstanding, m.DividendDate, m.ExDividendDate,
		)
		if err != nil {
			// handle the error
			log.Fatal(err)
		}

	case "APIs.IncomeStatements":
		var m APIs.IncomeStatements
		err := decoder.Decode(&m)
		e.Check(err)
		// return m.QuarterlyReports
	case "APIs.BalanceSheets":
		var m APIs.BalanceSheets
		err := decoder.Decode(&m)
		e.Check(err)
		// return m.QuarterlyReports
	case "APIs.CashFlowStatements":
		var m APIs.CashFlowStatements
		err := decoder.Decode(&m)
		e.Check(err)
		// return m.QuarterlyReports
	case "APIs.EarningsData":
		var m APIs.EarningsData
		err := decoder.Decode(&m)
		e.Check(err)
		// return m.QuarterlyEarnings
	// Commodities and Economic Indicators - use same structure
	// WTI, BRENT, nat gas, COPPER, ALUMINUM, WHEAT, CORN, COTTON, SUGAR, COFFEE
	case "APIs.CommodityPrices":
		var m APIs.CommodityPrices
		err := decoder.Decode(&m)
		e.Check(err)
		// return m.Data
	case "APIs.IntradayOHLCVs": // #todo totally doesnt work
		// timeseries1min comes out empty, idk why
		var m APIs.IntradayOHLCVs
		fmt.Printf("%v", m)
		err := decoder.Decode(&m)
		e.Check(err)

		if len(m.TimeSeries1min) == 0 {
			fmt.Println("m.TimeSeries1min is empty")
			return
		}
		valueStrings := make([]string, 0, len(m.TimeSeries1min))
		valueArgs := make([]interface{}, 0, len(m.TimeSeries1min)*6)
		i := 0
		fmt.Printf("Number of elements in m.TimeSeries1min: %d\n", len(m.TimeSeries1min))

		tickerID := GetTickerID(db, ticker)

		for _, b := range m.TimeSeries1min {
			fmt.Printf("Adding values: %v %v %v %v %v %v\n", tickerID, b.Open, b.High, b.Low, b.Close, b.Volume)
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)", i*6+1, i*6+2, i*6+3, i*6+4, i*6+5, i*6+6))
			valueArgs = append(valueArgs, tickerID)
			valueArgs = append(valueArgs, b.Open)
			valueArgs = append(valueArgs, b.High)
			valueArgs = append(valueArgs, b.Low)
			valueArgs = append(valueArgs, b.Close)
			valueArgs = append(valueArgs, b.Volume)
			i++
		}

		stmt := fmt.Sprintf("INSERT INTO dailyOHLCVs (tickerID, open, high, low, close, volume, datasource) VALUES %s", strings.Join(valueStrings, ","))
		fmt.Print(stmt)
		fmt.Printf("%v", valueArgs)

		_, err = db.Exec(stmt, valueArgs...)
		e.Check(err)
		// bulk insert strategy https://stackoverflow.com/questions/12486436/how-do-i-batch-sql-statements-with-package-database-sql
		// func BulkInsert(unsavedRows []*ExampleRowStruct) error {
		// 	valueStrings := make([]string, 0, len(unsavedRows))
		// 	valueArgs := make([]interface{}, 0, len(unsavedRows) * 3)
		// 	i := 0
		// 	for _, post := range unsavedRows {
		// 		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3))
		// 		valueArgs = append(valueArgs, post.Column1)
		// 		valueArgs = append(valueArgs, post.Column2)
		// 		valueArgs = append(valueArgs, post.Column3)
		// 		i++
		// 	}
		// 	stmt := fmt.Sprintf("INSERT INTO my_sample_table (column1, column2, column3) VALUES %s", strings.Join(valueStrings, ","))
		// 	_, err := db.Exec(stmt, valueArgs...)
		// 	return err
		// }
	case "APIs.DailyOHLCVs":
		var m APIs.DailyOHLCVs
		err := decoder.Decode(&m)
		e.Check(err)

		var tickerID int
		err = db.QueryRow("SELECT TickerID FROM tickers WHERE TickerSymbol = $1", "IBM").Scan(&tickerID)
		e.Check(err)
		fmt.Print("ticker id is")
		fmt.Print(tickerID)

		for date, m := range m.TimeSeries {
			_, err = db.Exec(`
    INSERT INTO dailyOHLCVs (tickerID, date, open, high, low, close, volume, datasource)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`, tickerID, date, m.Open, m.High,
				m.Low, m.Close, m.Volume, 1) // 1= alphavantage
			e.Check(err)
		}
	// case "APIs.TGLATs": // did not add to postgres, suspect unneeded
	// 	var m APIs.TGLATs // #todo suspect not properly implemented
	// 	err := decoder.Decode(&m)
	// 	e.Check(err)
	// 	output, err := json.Marshal(m) // perhaps change, but 3 maps
	// 	e.Check(err)
	// 	return string(output)
	default: // why do i need this? wont trigger, hm
		panic("confident I don't need this")

	}
}

func GetTickerID(db *sql.DB, ticker string) int {
	var tickerID int
	err := db.QueryRow("SELECT TickerID FROM tickers WHERE TickerSymbol = $1", ticker).Scan(&tickerID)
	e.Check(err)
	return tickerID
}
