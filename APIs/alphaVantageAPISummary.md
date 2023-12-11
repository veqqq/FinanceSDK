

### Summary of Alpha Vantage API

https://www.alphavantage.co/documentation/

- ignores uninteresting things like Quote Endpoint/GLOBAL_QUOTE
- ignores technical indicators
- ✅ means incorporated into API, ❌ means waiting, none means no intention to add but still worth adding to this summary


### Required Parameters - some have further parameters shown next line
- symbol
    - use - instead of period for prefered shares e.g. PBR-A
- apikey
    - get here with email: https://www.alphavantage.co/support/#api-key


- function - this gives you fundemental company data, forex etc.
    - ✅ INCOME_STATEMENT: https://www.alphavantage.co/query?function=INCOME_STATEMENT&symbol=IBM&apikey=demo
    - ✅ BALANCE_SHEET
    - ✅ OVERVIEW - company info, financial ratios etc. refreshed on earnings day
    - ✅ CASH_FLOW - annual and quartlery cash flow, normalized yields to GAAP and IFRS
    - ✅ EARNINGS - 

    - ✅ TIME_SERIES_INTRADAY - OHLCV (open, high, low, close, volume)
        - interval - 1min, 5min, 15min, 30min, 60min - required
        - adjusted - true if adjusted by historical split and dividend events, false
        - extended_hours - default true, false
        - month - default current month, format: month=2009-01, since 2000-01
        - outputsize - compact 100 data points, full month
        - datatype - default json, csv
    - ✅ TIME_SERIES_DAILY - OHLCV by day
        - outputsize - compact 100 data points, full 20+ years
        - datatype
        - n.b. adjusted is TIME_SERIES_DAILY_ADJUSTED which is premium
    - TIME_SERIES_WEEKLY_ADJUSTED or TIME_SERIES_WEEKLY
    - TIME_SERIES_MONTHLY_ADJUSTED or TIME_SERIES_MONTHLY


    - CURRENCY_EXCHANGE_RATE - realtime rate for pair
        - to_currency, from_currency - currency list: https://www.alphavantage.co/physical_currency_list/ required
    - ✅ FX_DAILY - daily (timestamp, open, high, low, close)
        - from_symbol, to_symbol
        - output size - default compact, full 20 years
    - FX_WEEKLY, ? FX_MONTHLY

    - ✅ WTI, BRENT, NATURAL_GAS - oil baby for 20 years - this block is the same but for
        - interval - default monthly, daily, weekly      - the intervals, but always use
    - ✅ COPPER, ALUMINUM, WHEAT, CORN, COTTON, SUGAR, COFFEEl - daily so lol
        - interval - default monthly, quarterly, annually
    - ✅ ALL_COMMODITIES - basket price of all commodities
        - interval - default monthly, quarterly, annually


    - ✅ REAL_GDP, REAL_GDP_PER_CAPITA - of US
        - interval - default annual, quarterly
    - ✅ TREASURY_YIELD
        - interval - default monthly, daily, weekly
        - maturity - default 10year, 3month, 2year,5,year,7year, 10 year, 30year
    - ✅ FEDERAL_FUNDS_RATE
        - interval - default monthly, daily, weekly
    - ✅ CPI - consumer price index
        - interval - default monthly, semiannual
    - ✅ INFLATION - annual inflation in consumer prices
    - ✅ RETAIL_SALES - monthly Advance Retail Sales: Retail Trade
    - ✅ DURABLES - monthly manufacturers' new orders of durables
    - ✅ UNEMPLOYMENT - monthly
    - ✅ NONFARM_PAYROLL


    - LISTING_STATUS - gives list of all active or unactive tickers
        - date - YYYY-MM-DD after 2010-01-01
        - state - active/delisted
    - EARNINGS_CALENDAR - expected earnings in next 3, 6, 12 months
        - symbol - default is all, can specify only 1 company
        - horizon - default is 3month, also 6month, 12month
    - IPO_CALENDAR - list of all IPOs in next 3 months


    - ❌ NEWS_SENTIMENT - will output article summaries, stock sentiment etc.!
        - tickers - `tickers=COIN,CRYPTO:BTC,FOREX:USD` will filter for articles that simultaneously mention them
        - topics - blockchain, earnings, ipo, mergers_and_acquisitions, financial_markets, economy_fiscal, economy_monetary, economy_marco, energy_transportation, finance, life_sciences, manufacturing, real_estate, retail_wholesale, technology. `topics=technology,ipo` will filter for articles that simultaneously cover technology and IPO
        - time_from, time_to - YYYYMMDDTHHMM e.g. time_from=20220410T0130
        - sort - default LATEST, EARLIEST, RELEVANCE
        - limit - default 50, 1000
    - ✅ TOP_GAINTERS_LOSERS - top 20 gainers, losers, most actively traded tickers for the day



---------

### Fundemental Data

- Provides:
- overview
    - SYMBOL
    - Sector
    - Industry
    - FiscalYearEnd
    - MarketCapitalization
    - EBITDA
    - PERatio
    - PEGRatio
    - BookValue
    - DividendPerShare
    - DividendYield
    - EPS
    - RevenuePerShareTTM - TTM = trailing 12 months
    - ProfitMargin
    - OperatingMarginTTM
    - ReturnOnAssetsTTM
    - ReturnOnEquityTTM
    - RevenueTTM
    - GrossProfitTTM
    - DilutedEPSTTM
    - QuarterlyEarningsGrowthYOY
    - QuarterlyRevenueGrowthYOY
    - AnalystTargetPrice
    - TrailingPE
    - FowardPE
    - PriceToSalesRatioTTM
    - PriceToBookRatio
    - EVToRevenue
    - EVToEBITDA
    - Beta
    - 52WeekHigh
    - 52WeekLow
    - 50DayMovingAverage
    - 200DayMovingAverage
    - SharesOutStanding
    - DividendDate
    - ExDividendDate

- income_statement - yearly, quarterly
    - grossProfit
    - totalRevenue
    - costOfRevenue
    - costofGoodsAndServicesSold
    - operatingIncome
    - sellingGeneralAndAdministrative
    - researchAndDevelopment
    - operatingExpenses
    - investmentIncomeNet
    - netInterestIncome
    - interestIncome
    - interestExpense
    - nonInterestIncome
    - otherNonOperatingIncome
    - depreciation
    - depreciationgAndamortization
    - incomeBeforeTax
    - incomeTaxExpense
    - interestAndDebtExpense
    - netIncomeFromContinuingOperations
    - comprehensiveIncomeNetOfTax
    - ebit
    - ebitda
    - netIncome

- balance_sheet - yearly, quarterly
    - totalAssets
    - totalCurrentassets
    - cashAndCashEquivalentsAtCarryingValue
    - cashAndShortTermInvestments
    - inventory
    - currentNetReceivables
    - totalNonCurrentAssets
    - propertyPlantEquipment
    - accumulatedDepreciationAmortizationPPE
    - intangibleAssets
    - intangibleAssetsExcludingGoodwill
    - goodwill
    - investments
    - longTermInvestments
    - shortTermInvestments
    - otherCurrentAssets
    - otherNonCurrentAssets
    - totalLiabilities
    - totalCurrentLiabilities
    - currentAccountsPayable
    - deferredRevenue
    - currentDebt
    - shortTermDebt
    - totalNonCurrentLiabilities
    - capitalLeaseObligations
    - longTermDebt
    - currentLongTermDebt
    - longTermDebtNoncurrent
    - shortLongTermDebtTotal
    - otherCurrentLiabilities
    - otherNonCurrentLiabilities
    - totalShareholderEquity
    - treasuryStock
    - retainedEarnings
    - commonStock
    - commonStockSharesOutstanding

- cash flow - yearly, quarterly
    - operatingCashflow
    - paymentsForOperatingActivities
    - proceedsFromOperatingActivities
    - changeInOperatingLiabilities
    - changeInOperatingAssets
    - depreciatingDepletionAndAmortization
    - capitalExpenditures
    - changeInReceivables
    - changeInInventory
    - profitLoss
    - cashflowFromInvestment
    - cashflowFromFinancing
    - proceedsFromRepaymentsOfShortTermDebt
    - paymentsForRepurchaseOfCommonStock
    - paymentsForRepurchaseofEquity
    - paymentsForRepurchaseofPreferredStock
    - dividendPayout
    - dividendPayoutCommonStock
    - dividendPayoutPreferredStock
    - proceedsFromIssuanceOfCommonStock
    - proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet
    - proceedsFromIssuanceOfPreferredStock
    - proceedsFromRepurchaseOfEquity
    - proceedsFromSaleOfTreasuryStock
    - changeInCashAndCashEquivalents
    - changeInExchangeRate ???
    - netIncome

- earnings - quarterly (better than annual)
    - reported eps 
    - estimates eps
    - surprise
    - surprisepercentage