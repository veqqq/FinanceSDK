package APIs

import (
	"FinanceSDK/e"
	"database/sql"
	"encoding/json"
	"time"
)

// querry builder returns this type, implemented
// by the formats into which Json in marshalled
// from which they are uploaded into SQL

func (SO StockOverview) Upload(db *sql.DB, ticker string, tickerID int, body []byte) {
	var m StockOverview
	err := json.Unmarshal(body, &m)
	CheckJSON(err, m)
	// tickerID := GetTickerID(db, ticker) <- should be contained in its surrounding struct
	result, err := db.Exec(`INSERT INTO stock_overviews (TickerID,
        TickerSymbol, datasource, asset_type, name, cik, exchange, currency, country, sector, industry,
        address, fiscal_year_end, latest_quarter, market_capitalization, ebitda, pe_ratio, peg_ratio,
        book_value, dividend_per_share, dividend_yield, eps, revenue_per_share_ttm, profit_margin,
        operating_margin_ttm, return_on_assets_ttm, return_on_equity_ttm, revenue_ttm, gross_profit_ttm,
        diluted_eps_ttm, quarterly_earnings_growth_yoy, quarterly_revenue_growth_yoy,
        analyst_target_price, trailing_pe, forward_pe, price_to_sales_ratio_ttm, price_to_book_ratio,
        ev_to_revenue, ev_to_ebitda, beta, day_moving_average_50,
        day_moving_average_200, shares_outstanding, dividend_date, ex_dividend_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
        $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38,
        $39, $40, $41, $42, $43, $44, $45) ON CONFLICT DO NOTHING`,
		tickerID, ticker, 1, m.AssetType, m.Name,
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
		m.SharesOutstanding, m.DividendDate, m.ExDividendDate)
	e.CheckDBInsert(err, result, ticker, "Stock_Overviews", m)
}

func (SO IncomeStatements) Upload(db *sql.DB, ticker string, tickerID int, body []byte) {
	var m IncomeStatements
	err := json.Unmarshal(body, &m)
	CheckJSON(err, m)
	for _, m := range m.QuarterlyReports {
		result, err := db.Exec(`INSERT INTO income_statements (TickerID, TickerSymbol, datasource, fiscal_date_ending, reported_currency,
		gross_profit, total_revenue, cost_of_revenue, cost_of_goods_and_services_sold,
		operating_income, selling_general_and_administrative, research_and_development,
		operating_expenses, investment_income_net, net_interest_income, interest_income,
		interest_expense, non_interest_income, other_non_operating_income, depreciation,
		depreciation_and_amortization, income_before_tax, income_tax_expense,
		interest_and_debt_expense, net_income_from_continuing_operations,
		comprehensive_income_net_of_tax, ebit, ebitda, net_income)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18,
		$19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29) ON CONFLICT DO NOTHING`,
			tickerID, ticker, 1, m.FiscalDateEnding, m.ReportedCurrency, m.GrossProfit,
			m.TotalRevenue, m.CostOfRevenue, m.CostofGoodsAndServicesSold,
			m.OperatingIncome, m.SellingGeneralAndAdministrative, m.ResearchAndDevelopment,
			m.OperatingExpenses, m.InvestmentIncomeNet, m.NetInterestIncome,
			m.InterestIncome, m.InterestExpense, m.NonInterestIncome,
			m.OtherNonOperatingIncome, m.Depreciation, m.DepreciationAndAmortization,
			m.IncomeBeforeTax, m.IncomeTaxExpense, m.InterestAndDebtExpense,
			m.NetIncomeFromContinuingOperations, m.ComprehensiveIncomeNetOfTax,
			m.EBIT, m.EBITDA, m.NetIncome)
		e.CheckDBInsert(err, result, ticker, "Income_Statements", m)
	}
}

func (SO BalanceSheets) Upload(db *sql.DB, ticker string, tickerID int, body []byte) {
	var m BalanceSheets
	err := json.Unmarshal(body, &m)
	CheckJSON(err, m)
	for _, m := range m.QuarterlyReports {
		result, err := db.Exec(`INSERT INTO balance_sheets (TickerID, TickerSymbol, datasource, fiscal_date_ending, reported_currency,
			total_assets, total_current_assets, cash_and_cash_equivalents_at_carrying_value,
			cash_and_short_term_investments, inventory, current_net_receivables,
			total_non_current_assets, property_plant_equipment,
			accumulated_depreciation_amortization_ppe, intangible_assets,
			intangible_assets_excluding_goodwill, goodwill, investments,
			long_term_investments, short_term_investments, other_current_assets,
			other_non_current_assets, total_liabilities, total_current_liabilities,
			current_accounts_payable, deferred_revenue, current_debt, short_term_debt,
			total_non_current_liabilities, capital_lease_obligations, long_term_debt,
			current_long_term_debt, long_term_debt_noncurrent, short_long_term_debt_total,
			other_current_liabilities, other_non_current_liabilities, total_shareholder_equity,
			treasury_stock, retained_earnings, common_stock, common_stock_shares_outstanding)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18,
			$19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35,
			$36, $37, $38, $39, $40, $41)  ON CONFLICT DO NOTHING`,
			tickerID, ticker, 1, m.FiscalDateEnding, m.ReportedCurrency,
			m.TotalAssets, m.TotalCurrentAssets, m.CashAndCashEquivalentsAtCarryingValue,
			m.CashAndShortTermInvestments, m.Inventory, m.CurrentNetReceivables,
			m.TotalNonCurrentAssets, m.PropertyPlantEquipment,
			m.AccumulatedDepreciationAmortizationPPE, m.IntangibleAssets,
			m.IntangibleAssetsExcludingGoodwill, m.Goodwill, m.Investments,
			m.LongTermInvestments, m.ShortTermInvestments, m.OtherCurrentAssets,
			m.OtherNonCurrentAssets, m.TotalLiabilities, m.TotalCurrentLiabilities,
			m.CurrentAccountsPayable, m.DeferredRevenue, m.CurrentDebt, m.ShortTermDebt,
			m.TotalNonCurrentLiabilities, m.CapitalLeaseObligations, m.LongTermDebt,
			m.CurrentLongTermDebt, m.LongTermDebtNoncurrent, m.ShortLongTermDebtTotal,
			m.OtherCurrentLiabilities, m.OtherNonCurrentLiabilities, m.TotalShareholderEquity,
			m.TreasuryStock, m.RetainedEarnings, m.CommonStock, m.CommonStockSharesOutstanding)
		e.CheckDBInsert(err, result, ticker, "Balance_Sheets", m)
	}
}

func (ed EarningsData) Upload(db *sql.DB, ticker string, tickerID int, body []byte) {
	var m EarningsData
	err := json.Unmarshal(body, &m)
	CheckJSON(err, m)
	for _, m := range m.QuarterlyReports {
		result, err := db.Exec(`
INSERT INTO earnings (TickerID, TickerSymbol, datasource, fiscal_date_ending, reported_date, reportedEPS, estimatedEPS, surprise, surprise_percentage)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ON CONFLICT DO NOTHING`,
			tickerID, ticker, 1, m.FiscalDateEnding, m.ReportedDate,
			m.ReportedEPS, m.EstimatedEPS, m.Surprise,
			m.SurprisePercentage)
		e.CheckDBInsert(err, result, ticker, "Earnings", m)
	}
}
func (SO CashFlowStatements) Upload(db *sql.DB, ticker string, tickerID int, body []byte) {
	var m CashFlowStatements
	err := json.Unmarshal(body, &m)
	CheckJSON(err, m)
	for _, m := range m.QuarterlyReports {
		result, err := db.Exec(`INSERT INTO cash_flow_statements (TickerID, TickerSymbol, datasource, fiscal_date_ending, reported_currency,
				operating_cashflow, payments_for_operating_activities,
				proceeds_from_operating_activities, change_in_operating_liabilities,
				change_in_operating_assets, depreciation_depletion_and_amortization,
				capital_expenditures, change_in_receivables, change_in_inventory,
				profit_loss, cashflow_from_investment, cashflow_from_financing,
				proceeds_from_repayments_of_short_term_debt,
				payments_for_repurchase_of_common_stock,
				payments_for_repurchase_of_equity,
				payments_for_repurchase_of_preferred_stock, dividend_payout,
				dividend_payout_common_stock, dividend_payout_preferred_stock,
				proceeds_from_issuance_of_common_stock,
				proceeds_from_issuance_of_long_term_debt_and_capital_securities,
				proceeds_from_issuance_of_preferred_stock,
				proceeds_from_repurchase_of_equity,
				proceeds_from_sale_of_treasury_stock,
				change_in_cash_and_cash_equivalents, change_in_exchange_rate, net_income)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17,
				$18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32)  ON CONFLICT DO NOTHING`,
			tickerID, ticker, 1, m.FiscalDateEnding,
			m.ReportedCurrency, m.OperatingCashflow,
			m.PaymentsForOperatingActivities,
			m.ProceedsFromOperatingActivities,
			m.ChangeInOperatingLiabilities,
			m.ChangeInOperatingAssets,
			m.DepreciationDepletionAndAmortization,
			m.CapitalExpenditures, m.ChangeInReceivables,
			m.ChangeInInventory, m.ProfitLoss,
			m.CashflowFromInvestment, m.CashflowFromFinancing,
			m.ProceedsFromRepaymentsOfShortTermDebt,
			m.PaymentsForRepurchaseOfCommonStock,
			m.PaymentsForRepurchaseOfEquity,
			m.PaymentsForRepurchaseOfPreferredStock,
			m.DividendPayout, m.DividendPayoutCommonStock,
			m.DividendPayoutPreferredStock,
			m.ProceedsFromIssuanceOfCommonStock,
			m.ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet,
			m.ProceedsFromIssuanceOfPreferredStock,
			m.ProceedsFromRepurchaseOfEquity,
			m.ProceedsFromSaleOfTreasuryStock,
			m.ChangeInCashAndCashEquivalents,
			m.ChangeInExchangeRate, m.NetIncome)
		e.CheckDBInsert(err, result, ticker, "Cash_Flow_Statements", m)
	}
}

// Commodities and Economic Indicators - use same structure
// WTI, BRENT, nat gas, COPPER, ALUMINUM, WHEAT, CORN, COTTON, SUGAR, COFFEE
func (SO CommodityPrices) Upload(db *sql.DB, ticker string, tickerID int, body []byte) {
	var m CommodityPrices
	err := json.Unmarshal(body, &m)
	CheckJSON(err, m)
	for _, commodityData := range m.Data {
		date, _ := time.Parse("2006-01-02", commodityData.Date)
		e.Check(err)
		result, err := db.Exec(`
			INSERT INTO commodities (TickerID, TickerSymbol, date, value, datasource)
			VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING`,
			tickerID, ticker, date, commodityData.Value, 1)
		e.CheckDBInsert(err, result, ticker, "Commodities", m)
	}
}

func (SO IntradayOHLCVs) Upload(db *sql.DB, ticker string, tickerID int, body []byte) {
	var m IntradayOHLCVs
	err := json.Unmarshal(body, &m)
	CheckJSON(err, m)
	for time, m := range m.TimeSeries1min {
		result, err := db.Exec(`INSERT INTO intradayohlcvs (TickerID, TickerSymbol, timestamp, open, high, low, close, volume, datasource)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)  ON CONFLICT DO NOTHING`, tickerID, ticker, time, m.Open, m.High,
			m.Low, m.Close, m.Volume, 1) // 1= alphavantage
		e.CheckDBInsert(err, result, ticker, "IntradayOHLCVs", m)
	}
}

//bulk insert doesnt work
// case "APIs.IntradayOHLCVs":
// 	// timeseries1min comes out empty, idk why
// 	var m APIs.IntradayOHLCVs
// 	fmt.Printf("%v", m)
// 	err := decoder.Decode(&m)
// 	e.Check(err)

// 	if len(m.TimeSeries1min) == 0 {
// 		fmt.Println("m.TimeSeries1min is empty")
// 		return
// 	}
// 	valueStrings := make([]string, 0, len(m.TimeSeries1min))
// 	valueArgs := make([]interface{}, 0, len(m.TimeSeries1min)*6)
// 	i := 0
// 	fmt.Printf("Number of elements in m.TimeSeries1min: %d\n", len(m.TimeSeries1min))

// 	tickerID := GetTickerID(db, ticker)

// 	for _, b := range m.TimeSeries1min {
// 		fmt.Printf("Adding values: %v %v %v %v %v %v\n", tickerID, b.Open, b.High, b.Low, b.Close, b.Volume)
// 		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)", i*6+1, i*6+2, i*6+3, i*6+4, i*6+5, i*6+6))
// 		valueArgs = append(valueArgs, tickerID)
// 		valueArgs = append(valueArgs, b.Open)
// 		valueArgs = append(valueArgs, b.High)
// 		valueArgs = append(valueArgs, b.Low)
// 		valueArgs = append(valueArgs, b.Close)
// 		valueArgs = append(valueArgs, b.Volume)
// 		i++
// 	}

// 	stmt := fmt.Sprintf("INSERT INTO dailyOHLCVs (tickerID, open, high, low, close, volume, datasource) VALUES %s", strings.Join(valueStrings, ","))
// 	fmt.Print(stmt)
// 	fmt.Printf("%v", valueArgs)

// 	_, err = db.Exec(stmt, valueArgs...)
// 			e.CheckDBInsert(err, ticker, structType, m)
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

func (SO DailyOHLCVs) Upload(db *sql.DB, ticker string, tickerID int, body []byte) {
	var m DailyOHLCVs
	err := json.Unmarshal(body, &m)
	CheckJSON(err, m)
	for date, m := range m.TimeSeries {
		result, err := db.Exec(`INSERT INTO dailyOHLCVs (TickerID, TickerSymbol, date, open, high, low, close, volume, datasource)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)  ON CONFLICT DO NOTHING`, tickerID, ticker, date, m.Open, m.High, m.Low, m.Close, m.Volume, 1) // 1= alphavantage
		e.CheckDBInsert(err, result, ticker, "DailyOHLCVs", m)
	}
}

func (SO ForexPrices) Upload(db *sql.DB, ticker string, tickerID int, body []byte) {
	var m ForexPrices
	err := json.Unmarshal(body, &m)
	CheckJSON(err, m)
	for date, m := range m.TimeSeriesFX {
		result, err := db.Exec(`INSERT INTO forex (TickerID, TickerSymbol, date, open, high, low, close, datasource)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)  ON CONFLICT DO NOTHING`, tickerID, ticker, date, m.Open, m.High, m.Low, m.Close, 1)
		e.CheckDBInsert(err, result, ticker, "Forex", m)
	}
}
