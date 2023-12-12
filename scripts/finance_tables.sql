-- -- Block all users by default
-- REVOKE ALL ON DATABASE pysecmaster FROM public;
-- REVOKE ALL ON SCHEMA public FROM public;

/*
Create an index on the Prices table for the
following columns as such (TickerId, AsOfDateTime, Price). This will allow you 
to efficiently search for a set of prices for a given ticker over a range of dates. 
A B-Tree index has a search time complexity of O(Log(n)) so it will be very fast
to find a subset of the data, even with billions of records in the table.
*/

/*
Track splits and dividends on a daily basis and delete and then bulk insert 
data for every symbol that needs to be changed.
*/

-- jobqueue (human can add to this, updater service checks tickers' last updated etc. and overview's industry)
  
CREATE TABLE tickers (
    TickerID serial primary key,
    TickerSymbol varchar unique,
    type varchar, -- stock, etf, macro, commodity? Manual labling?
    lastupdated date,
    importance varchar -- how often to update
    -- q = quarterly, m = monthly
);

Create TABLE jobqueue (
    TickerID int REFERENCES tickers(TickerID),
    Depth varchar -- n = only ohcvls, a = accounting docs and ohcvls, i = intraday ohcvls
);

CREATE TABLE datasources (
  SourceID serial primary key,
  SourceName varchar,
  SourceURL varchar
);

CREATE TABLE dailyOHLCVs (
    TickerID int REFERENCES tickers(TickerID),
    date date,
    open decimal(10, 2),
    high decimal(10, 2),
    low decimal(10, 2),
    close decimal(10, 2),
    volume decimal(11),
    datasource int REFERENCES datasources(SourceID),
    primary key (TickerID, date, datasource) -- makes sure each date tickerID combo is unique
);

CREATE TABLE intradayOHLCVs (
    TickerID int REFERENCES tickers(TickerID),
    timestamp timestamp,
    open decimal(10, 2),
    high decimal(10, 2),
    low decimal(10, 2),
    close decimal(10, 2),
    volume decimal(11),
    datasource int REFERENCES datasources(SourceID),
    primary key (TickerID, timestamp, datasource) -- makes sure each date tickerID combo is unique
);

-- 3 accounting docs + overview (only quarterly, not annual):
-- #todo unify the statements after testing a few hundred docs
-- to verify e.g. net_income will always be the same
-- descrepencies are very possible and missing them would
-- corrupt all data, so even if not finding them in the data, perhaps best to keep them isolated

CREATE TABLE stock_overviews (
    id int REFERENCES tickers(TickerID),
    symbol varchar REFERENCES tickers(TickerSymbol),
    asset_type varchar,
    name varchar,
    cik varchar,
    exchange varchar,
    currency varchar,
    country varchar,
    sector varchar,
    industry varchar,
    address varchar,
    fiscal_year_end varchar unique,
    latest_quarter varchar,
    market_capitalization decimal,
    ebitda decimal,
    pe_ratio decimal,
    peg_ratio decimal,
    book_value decimal,
    dividend_per_share decimal,
    dividend_yield decimal,
    eps decimal,
    revenue_per_share_ttm decimal,
    profit_margin decimal,
    operating_margin_ttm decimal,
    return_on_assets_ttm decimal,
    return_on_equity_ttm decimal,
    revenue_ttm decimal,
    gross_profit_ttm decimal,
    diluted_eps_ttm decimal,
    quarterly_earnings_growth_yoy decimal,
    quarterly_revenue_growth_yoy decimal,
    analyst_target_price decimal,
    trailing_pe decimal,
    forward_pe decimal,
    price_to_sales_ratio_ttm decimal,
    price_to_book_ratio decimal,
    ev_to_revenue decimal,
    ev_to_ebitda decimal,
    beta decimal,
    day_moving_average_50 decimal,
    day_moving_average_200 decimal,
    shares_outstanding decimal,
    dividend_date varchar,
    ex_dividend_date varchar
);

CREATE TABLE income_statements (
    id int REFERENCES tickers(TickerID),
    fiscal_date_ending date unique,
    reported_currency varchar,
    gross_profit decimal,
    total_revenue decimal,
    cost_of_revenue decimal,
    cost_of_goods_and_services_sold decimal,
    operating_income decimal,
    selling_general_and_administrative decimal,
    research_and_development decimal,
    operating_expenses decimal,
    investment_income_net decimal,
    net_interest_income decimal,
    interest_income decimal,
    interest_expense decimal,
    non_interest_income decimal,
    other_non_operating_income decimal,
    depreciation decimal,
    depreciation_and_amortization decimal,
    income_before_tax decimal,
    income_tax_expense decimal,
    interest_and_debt_expense decimal,
    net_income_from_continuing_operations decimal,
    comprehensive_income_net_of_tax decimal,
    ebit decimal,
    ebitda decimal,
    net_income decimal
);

CREATE TABLE balance_sheets (
    id int REFERENCES tickers(TickerID),
    fiscal_date_ending date unique,
    reported_currency varchar,
    total_assets decimal,
    total_current_assets decimal,
    cash_and_cash_equivalents_at_carrying_value decimal,
    cash_and_short_term_investments decimal,
    inventory decimal,
    current_net_receivables decimal,
    total_non_current_assets decimal,
    property_plant_equipment decimal,
    accumulated_depreciation_amortization_ppe decimal,
    intangible_assets decimal,
    intangible_assets_excluding_goodwill decimal,
    goodwill decimal,
    investments decimal,
    long_term_investments decimal,
    short_term_investments decimal,
    other_current_assets decimal,
    other_non_current_assets decimal,
    total_liabilities decimal,
    total_current_liabilities decimal,
    current_accounts_payable decimal,
    deferred_revenue decimal,
    current_debt decimal,
    short_term_debt decimal,
    total_non_current_liabilities decimal,
    capital_lease_obligations decimal,
    long_term_debt decimal,
    current_long_term_debt decimal,
    long_term_debt_noncurrent decimal,
    short_long_term_debt_total decimal,
    other_current_liabilities decimal,
    other_non_current_liabilities decimal,
    total_shareholder_equity decimal,
    treasury_stock decimal,
    retained_earnings decimal,
    common_stock decimal,
    common_stock_shares_outstanding decimal
);

CREATE TABLE cash_flow_statements (
    id int REFERENCES tickers(TickerID),
    fiscal_date_ending date unique,
    reported_currency varchar,
    operating_cashflow decimal,
    payments_for_operating_activities decimal,
    proceeds_from_operating_activities decimal,
    change_in_operating_liabilities decimal,
    change_in_operating_assets decimal,
    depreciation_depletion_and_amortization decimal,
    capital_expenditures decimal,
    change_in_receivables decimal,
    change_in_inventory decimal,
    profit_loss decimal,
    cashflow_from_investment decimal,
    cashflow_from_financing decimal,
    proceeds_from_repayments_of_short_term_debt decimal,
    payments_for_repurchase_of_common_stock decimal,
    payments_for_repurchase_of_equity decimal,
    payments_for_repurchase_of_preferred_stock decimal,
    dividend_payout decimal,
    dividend_payout_common_stock decimal,
    dividend_payout_preferred_stock decimal,
    proceeds_from_issuance_of_common_stock decimal,
    --  proceeds_from_issuance_of_long_term_debt_and_capital_securities_net is too long
    proceeds_from_issuance_of_long_term_debt_and_capital_securities decimal,
    proceeds_from_issuance_of_preferred_stock decimal,
    proceeds_from_repurchase_of_equity decimal,
    proceeds_from_sale_of_treasury_stock decimal,
    change_in_cash_and_cash_equivalents decimal,
    change_in_exchange_rate decimal,
    net_income decimal
);


-- Commodities and macro, the types are hard here
-- e.g.: sugar: value":"24.9216494133885 <- 15!
CREATE TABLE commodities ( -- commodities and macro indicators
  id int REFERENCES tickers(TickerID), -- these specific tickers have different formats
  date date,
  value decimal(16,12),
  datasource int REFERENCES datasources(SourceID)
);

-- #todo
-- make 2nd commodity table for ones like oil or bond rates with low precision requirements

CREATE TABLE forex (
    fromCurrency varchar,
    toCurrency varchar,
    date date,
    open decimal(10,7), -- unsure, yen is the main issue here
    high decimal(10,7), -- e.g. close":"0.00643
    low decimal(10,7),
    close decimal(10,7),
    datasource int REFERENCES datasources(SourceID)
);
  