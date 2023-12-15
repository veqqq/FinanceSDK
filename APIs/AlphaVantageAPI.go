package APIs

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// earnings statements, tglats, sentiment are not used/implemented in sql

// n.b. alphavantage dates are already in sql format, no reformating needed

///////////////////////
// Stock Ticker Data:

// TIME_SERIES_DAILY
type DailyOHLCVs struct {
	MetaData   DailyOHLCVMetaData    `json:"Meta Data"`
	TimeSeries map[string]DailyOHLCV `json:"Time Series (Daily)"`
}

// reformat stuff for this?
type DailyOHLCVMetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

type DailyOHLCV struct {
	Open   float64 `json:"1. open,string"`
	High   float64 `json:"2. high,string"`
	Low    float64 `json:"3. low,string"`
	Close  float64 `json:"4. close,string"`
	Volume int64   `json:"5. volume,string"`
}

// TIME_SERIES_INTRADAY
type IntradayOHLCVs struct {
	MetaData       IntradayOHLCVMetaData    `json:"Meta Data"`
	TimeSeries1min map[string]IntradayOHLCV `json:"Time Series (1min)"`
}

type IntradayOHLCVMetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	Interval      string `json:"4. Interval"`
	OutputSize    string `json:"5. Output Size"`
	TimeZone      string `json:"6. Time Zone"`
}

type IntradayOHLCV struct {
	Open   float64 `json:"1. open,string"`
	High   float64 `json:"2. high,string"`
	Low    float64 `json:"3. low,string"`
	Close  float64 `json:"4. close,string"`
	Volume int64   `json:"5. volume,string"`
}

//////////////////
// Stock Fundementals:
// The 4 financial statements + "overview"

// income statement
// json: invalid use of ,string struct tag, trying to unmarshal "None" into float64
// unsure why, so unmarshaling into map any

type IncomeStatements struct {
	Symbol           string            `json:"symbol"`
	AnnualReports    []IncomeStatement `json:"annualReports"`
	QuarterlyReports []IncomeStatement `json:"quarterlyReports"`
	FinancialStatements
}

type IncomeStatement struct {
	FiscalDateEnding                  string     `json:"fiscalDateEnding"`
	ReportedCurrency                  string     `json:"reportedCurrency"`
	GrossProfit                       ownFloat64 `json:"grossProfit"`
	TotalRevenue                      ownFloat64 `json:"totalRevenue"`
	CostOfRevenue                     ownFloat64 `json:"costOfRevenue"`
	CostofGoodsAndServicesSold        ownFloat64 `json:"costofGoodsAndServicesSold"`
	OperatingIncome                   ownFloat64 `json:"operatingIncome"`
	SellingGeneralAndAdministrative   ownFloat64 `json:"sellingGeneralAndAdministrative"`
	ResearchAndDevelopment            ownFloat64 `json:"researchAndDevelopment"`
	OperatingExpenses                 ownFloat64 `json:"operatingExpenses"`
	InvestmentIncomeNet               ownFloat64 `json:"investmentIncomeNet"`
	NetInterestIncome                 ownFloat64 `json:"netInterestIncome"`
	InterestIncome                    ownFloat64 `json:"interestIncome"`
	InterestExpense                   ownFloat64 `json:"interestExpense"`
	NonInterestIncome                 ownFloat64 `json:"nonInterestIncome"`
	OtherNonOperatingIncome           ownFloat64 `json:"otherNonOperatingIncome"`
	Depreciation                      ownFloat64 `json:"depreciation"`
	DepreciationAndAmortization       ownFloat64 `json:"depreciationAndAmortization"`
	IncomeBeforeTax                   ownFloat64 `json:"incomeBeforeTax"`
	IncomeTaxExpense                  ownFloat64 `json:"incomeTaxExpense"`
	InterestAndDebtExpense            ownFloat64 `json:"interestAndDebtExpense"`
	NetIncomeFromContinuingOperations ownFloat64 `json:"netIncomeFromContinuingOperations"`
	ComprehensiveIncomeNetOfTax       ownFloat64 `json:"comprehensiveIncomeNetOfTax"`
	EBIT                              ownFloat64 `json:"ebit"`
	EBITDA                            ownFloat64 `json:"ebitda"`
	NetIncome                         ownFloat64 `json:"netIncome"`
}

// https://github.com/veqqq/StockDataSDK/issues/6
// implement UnmarshalJSON to avoid this error:
// json: invalid use of ,string struct tag, trying to unmarshal "None" into float64
// should add every "None" possible field to aux and give it a similar if statement below
//

// balance sheet

type BalanceSheets struct {
	Symbol           string         `json:"symbol"`
	AnnualReports    []BalanceSheet `json:"annualReports"`
	QuarterlyReports []BalanceSheet `json:"quarterlyReports"`
	FinancialStatements
}

type BalanceSheet struct {
	FiscalDateEnding                       string     `json:"fiscalDateEnding"`
	ReportedCurrency                       string     `json:"reportedCurrency"`
	TotalAssets                            ownFloat64 `json:"totalAssets"`
	TotalCurrentAssets                     ownFloat64 `json:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue  ownFloat64 `json:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments            ownFloat64 `json:"cashAndShortTermInvestments"`
	Inventory                              ownFloat64 `json:"inventory"`
	CurrentNetReceivables                  ownFloat64 `json:"currentNetReceivables"`
	TotalNonCurrentAssets                  ownFloat64 `json:"totalNonCurrentAssets"`
	PropertyPlantEquipment                 ownFloat64 `json:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPPE ownFloat64 `json:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                       ownFloat64 `json:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill      ownFloat64 `json:"intangibleAssetsExcludingGoodwill"`
	Goodwill                               ownFloat64 `json:"goodwill"`
	Investments                            ownFloat64 `json:"investments"`
	LongTermInvestments                    ownFloat64 `json:"longTermInvestments"`
	ShortTermInvestments                   ownFloat64 `json:"shortTermInvestments"`
	OtherCurrentAssets                     ownFloat64 `json:"otherCurrentAssets"`
	OtherNonCurrentAssets                  ownFloat64 `json:"otherNonCurrentAssets"`
	TotalLiabilities                       ownFloat64 `json:"totalLiabilities"`
	TotalCurrentLiabilities                ownFloat64 `json:"totalCurrentLiabilities"`
	CurrentAccountsPayable                 ownFloat64 `json:"currentAccountsPayable"`
	DeferredRevenue                        ownFloat64 `json:"deferredRevenue"`
	CurrentDebt                            ownFloat64 `json:"currentDebt"`
	ShortTermDebt                          ownFloat64 `json:"shortTermDebt"`
	TotalNonCurrentLiabilities             ownFloat64 `json:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations                ownFloat64 `json:"capitalLeaseObligations"`
	LongTermDebt                           ownFloat64 `json:"longTermDebt"`
	CurrentLongTermDebt                    ownFloat64 `json:"currentLongTermDebt"`
	LongTermDebtNoncurrent                 ownFloat64 `json:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                 ownFloat64 `json:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities                ownFloat64 `json:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities             ownFloat64 `json:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                 ownFloat64 `json:"totalShareholderEquity"`
	TreasuryStock                          ownFloat64 `json:"treasuryStock"`
	RetainedEarnings                       ownFloat64 `json:"retainedEarnings"`
	CommonStock                            ownFloat64 `json:"commonStock"`
	CommonStockSharesOutstanding           ownFloat64 `json:"commonStockSharesOutstanding"`
}

// Cash flow

type CashFlowStatements struct {
	Symbol           string              `json:"symbol"`
	AnnualReports    []CashFlowStatement `json:"annualReports"`
	QuarterlyReports []CashFlowStatement `json:"quarterlyReports"`
	FinancialStatements
}

type CashFlowStatement struct {
	FiscalDateEnding                                          string     `json:"fiscalDateEnding"`
	ReportedCurrency                                          string     `json:"reportedCurrency"`
	OperatingCashflow                                         ownFloat64 `json:"operatingCashflow"`
	PaymentsForOperatingActivities                            ownFloat64 `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities                           ownFloat64 `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities                              ownFloat64 `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets                                   ownFloat64 `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization                      ownFloat64 `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                                       ownFloat64 `json:"capitalExpenditures"`
	ChangeInReceivables                                       ownFloat64 `json:"changeInReceivables"`
	ChangeInInventory                                         ownFloat64 `json:"changeInInventory"`
	ProfitLoss                                                ownFloat64 `json:"profitLoss"`
	CashflowFromInvestment                                    ownFloat64 `json:"cashflowFromInvestment"`
	CashflowFromFinancing                                     ownFloat64 `json:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt                     ownFloat64 `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock                        ownFloat64 `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity                             ownFloat64 `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock                     ownFloat64 `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                                            ownFloat64 `json:"dividendPayout"`
	DividendPayoutCommonStock                                 ownFloat64 `json:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock                              ownFloat64 `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock                         ownFloat64 `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet ownFloat64 `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock                      ownFloat64 `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity                            ownFloat64 `json:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock                           ownFloat64 `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents                            ownFloat64 `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                                      ownFloat64 `json:"changeInExchangeRate"`
	NetIncome                                                 ownFloat64 `json:"netIncome"`
}

// Earnings
type EarningsData struct {
	Symbol            string              `json:"symbol"`
	AnnualEarnings    []AnnualEarnings    `json:"annualEarnings"`
	QuarterlyEarnings []QuarterlyEarnings `json:"quarterlyEarnings"`
	FinancialStatements
}

type AnnualEarnings struct { // lol pointless
	FiscalDateEnding string     `json:"fiscalDateEnding"`
	ReportedEPS      ownFloat64 `json:"reportedEPS"`
}

type QuarterlyEarnings struct {
	FiscalDateEnding   string     `json:"fiscalDateEnding"`
	ReportedDate       string     `json:"reportedDate"`
	ReportedEPS        ownFloat64 `json:"reportedEPS"`
	EstimatedEPS       ownFloat64 `json:"estimatedEPS"`
	Surprise           ownFloat64 `json:"surprise"`
	SurprisePercentage ownFloat64 `json:"surprisePercentage"`
}

// overview
type StockOverview struct {
	Symbol                     string     `json:"Symbol"`
	AssetType                  string     `json:"AssetType"`
	Name                       string     `json:"Name"`
	Description                string     `json:"Description"`
	CIK                        string     `json:"CIK"`
	Exchange                   string     `json:"Exchange"`
	Currency                   string     `json:"Currency"`
	Country                    string     `json:"Country"`
	Sector                     string     `json:"Sector"`
	Industry                   string     `json:"Industry"`
	Address                    string     `json:"Address"`
	FiscalYearEnd              string     `json:"FiscalYearEnd"`
	LatestQuarter              string     `json:"LatestQuarter"`
	MarketCapitalization       ownFloat64 `json:"MarketCapitalization"`
	EBITDA                     ownFloat64 `json:"EBITDA"`
	PERatio                    ownFloat64 `json:"PERatio"`
	PEGRatio                   ownFloat64 `json:"PEGRatio"`
	BookValue                  ownFloat64 `json:"BookValue"`
	DividendPerShare           ownFloat64 `json:"DividendPerShare"`
	DividendYield              ownFloat64 `json:"DividendYield"`
	EPS                        ownFloat64 `json:"EPS"`
	RevenuePerShareTTM         ownFloat64 `json:"RevenuePerShareTTM"`
	ProfitMargin               ownFloat64 `json:"ProfitMargin"`
	OperatingMarginTTM         ownFloat64 `json:"OperatingMarginTTM"`
	ReturnOnAssetsTTM          ownFloat64 `json:"ReturnOnAssetsTTM"`
	ReturnOnEquityTTM          ownFloat64 `json:"ReturnOnEquityTTM"`
	RevenueTTM                 ownFloat64 `json:"RevenueTTM"`
	GrossProfitTTM             ownFloat64 `json:"GrossProfitTTM"`
	DilutedEPSTTM              ownFloat64 `json:"DilutedEPSTTM"`
	QuarterlyEarningsGrowthYOY ownFloat64 `json:"QuarterlyEarningsGrowthYOY"`
	QuarterlyRevenueGrowthYOY  ownFloat64 `json:"QuarterlyRevenueGrowthYOY"`
	AnalystTargetPrice         ownFloat64 `json:"AnalystTargetPrice"`
	TrailingPE                 ownFloat64 `json:"TrailingPE"`
	ForwardPE                  ownFloat64 `json:"ForwardPE"`
	PriceToSalesRatioTTM       ownFloat64 `json:"PriceToSalesRatioTTM"`
	PriceToBookRatio           ownFloat64 `json:"PriceToBookRatio"`
	EVToRevenue                ownFloat64 `json:"EVToRevenue"`
	EVToEBITDA                 ownFloat64 `json:"EVToEBITDA"`
	Beta                       ownFloat64 `json:"Beta"`
	WeekHigh                   ownFloat64 `json:"52WeekHigh"`
	WeekLow                    ownFloat64 `json:"52WeekLow"`
	DayMovingAverage50         ownFloat64 `json:"50DayMovingAverage"`
	DayMovingAverage200        ownFloat64 `json:"200DayMovingAverage"`
	SharesOutstanding          ownFloat64 `json:"SharesOutstanding"`
	DividendDate               string     `json:"DividendDate"`
	ExDividendDate             string     `json:"ExDividendDate"`
}

///////

///////
// Commodities and Economic Indicators - use same structure

// WTI, BRENT, nat gas, COPPER, ALUMINUM, WHEAT, CORN, COTTON, SUGAR, COFFEE

// REAL_GDP in billions of dollars - same structure as commodities...
// real gdp per cap - in "chained 2012 dollars"
// fed funds rate - in percent - daily? monthly?
// cpi - "index 1982-1984=100" - monthly
// inflation - in percent, only annual
// retail sales - in millions, only monthly
// durable goods orders - in millions, only monthly
// unemployment - in percent, only monthly
// nonfarm payroll - in thousands of people, only monthly
// treasury yield - in percent, monthly or daily?
//          // maturities: 3month, 2year, 5year, 7year, 10year

type CommodityPrices struct { // rename
	Name string `json:"name"` // e.g. global price of copper or henry hub...
	// Interval string `json:"interval"` // should be daily always
	// Unit     string           `json:"unit"` // implicit in the commodity type
	Data []CommodityPrice `json:"data"`
}

type CommodityPrice struct { // rename
	Date  string  `json:"date,omitempty"`
	Value float64 `json:"value,string,omitempty"`
}

// CommodityPrice satisfies Unmarshaler
func (c *CommodityPrice) UnmarshalJSON(data []byte) error {
	var aux struct {
		Date  string `json:"date"`
		Value string `json:"value"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	// If val or date are useless, ignore struct
	// some dates have "." val, which normal unmarshaler doesn't handle, hence writing this
	if aux.Value == "null" || aux.Value == "" || aux.Value == "." || aux.Value == "0" || aux.Value == "0.0" || aux.Date == "" || aux.Date == "null" || aux.Date == "." {
		return nil
	}
	value, err := strconv.ParseFloat(aux.Value, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	c.Date = aux.Date
	c.Value = value
	return nil
}

/////////
//// Other

// exchange rates, daily

type ForexPrices struct {
	MetaData     ForexMetaData         `json:"Meta Data"`
	TimeSeriesFX map[string]ForexPrice `json:"Time Series FX (Daily)"` // monthly is the same, just uses (Monthly here)
}

type ForexMetaData struct {
	Information string `json:"1. Information"` // "1. Information": "Forex Daily Prices (open, high, low, close)",
	FromSymbol  string `json:"2. From Symbol"` // "2. From Symbol": "EUR",
	ToSymbol    string `json:"3. To Symbol"`   // "3. To Symbol": "USD",
	//	OutputSize    string `json:"4. Output Size"`
	// LastRefreshed string `json:"5. Last Refreshed"`
	// TimeZone string `json:"6. Time Zone"` #todo ignoring this can be dangerous when more vendors are used
}

type ForexPrice struct {
	Open  float64 `json:"1. open,string"` // "1. open": "1.10020",
	High  float64 `json:"2. high,string"`
	Low   float64 `json:"3. low,string"`
	Close float64 `json:"4. close,string"`
}

// 20 TOP_GAINERS_LOSERS  and most actively traded
// #todo not implemented in SQL

type TGLATs struct {
	Metadata     string  `json:"metadata"`
	LastUpdated  string  `json:"last_updated"`
	TopGainers   []TGLAT `json:"top_gainers"`
	TopLosers    []TGLAT `json:"top_losers"`
	MostActively []TGLAT `json:"most_actively_traded"`
}

type TGLAT struct {
	Ticker           string  `json:"ticker"`
	Price            float64 `json:"price,string"`
	ChangeAmount     float64 `json:"change_amount,string"`
	ChangePercentage string  `json:"change_percentage"` // has %, later implement json unmarshaler for this
	Volume           int64   `json:"volume,string"`
}

// News sentiment - complicated beast
// #todo - what to even do with this?

type SentimentData struct {
	Items                    string     `json:"items"`
	SentimentScoreDefinition string     `json:"sentiment_score_definition"`
	RelevanceScoreDefinition string     `json:"relevance_score_definition"`
	Feed                     []FeedData `json:"feed"`
}

type FeedData struct {
	Title                 string                `json:"title"`
	URL                   string                `json:"url"`
	TimePublished         string                `json:"time_published"` // "20230805T122000",
	Authors               []string              `json:"authors"`
	Summary               string                `json:"summary"`
	BannerImage           string                `json:"banner_image"`           // lol, for a spam farm?
	Source                string                `json:"source"`                 // "CNBC",
	CategoryWithinSource  string                `json:"category_within_source"` // "Top News",
	SourceDomain          string                `json:"source_domain"`          // "www.cnbc.com",
	Topics                []TopicData           `json:"topics"`
	OverallSentimentScore float64               `json:"overall_sentiment_score"` //	0.072673,
	OverallSentimentLabel string                `json:"overall_sentiment_label"` //	"Neutral",
	TickerSentiment       []TickerSentimentData `json:"ticker_sentiment"`
}

type TopicData struct {
	Topic          string  `json:"topic"`                  // "Technology",
	RelevanceScore float64 `json:"relevance_score,string"` // "1.0"
}

type TickerSentimentData struct {
	Ticker               string  `json:"ticker"`
	RelevanceScore       float64 `json:"relevance_score,string"`        // "0.699089",
	TickerSentimentScore float64 `json:"ticker_sentiment_score,string"` // "0.116531",
	TickerSentimentLabel string  `json:"ticker_sentiment_label"`        //  "Neutral"
}

// var sentimentData SentimentData
// if err == json.Unmarshal([]byte(jsonData), &sentimentData); err != nil {}
// 	fmt.Println("Error unmarshaling JSON:", err)
// 	return
// }
