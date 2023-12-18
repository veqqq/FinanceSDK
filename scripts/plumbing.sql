-- jobqueue (human can add to this with cli, updater service checks tickers'
-- last updated etc. and overview's industry)
CREATE TABLE jobqueue (
    TickerID int REFERENCES tickers(TickerID),
    TickerSymbol varchar unique REFERENCES tickers(TickerSymbol),
    coverage varchar, -- intradayOHLCVs+statements+dailyOHLCVs, statements+dailyOHLCVs, daily
    lastupdated date

);

-- queryqueue has the "url"s fed into querybuilder which will produce jobs for individual things
CREATE TABLE queryqueue (
    TickerID int REFERENCES tickers(TickerID),
    TickerSymbol varchar REFERENCES tickers(TickerSymbol),
    QueryTicker varchar,
    primary key (TickerID, TickerSymbol, QueryTicker)
);

-- ReportApiCall() increments after each call
CREATE TABLE ApiCalls (
    date date,
    CallCount int,
    datasource int REFERENCES datasources(SourceID)
);