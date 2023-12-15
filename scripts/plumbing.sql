-- jobqueue (human can add to this with cli, updater service checks tickers'
-- last updated etc. and overview's industry)
CREATE TABLE jobqueue (
    TickerID int REFERENCES tickers(TickerID),
    TickerSymbol varchar unique REFERENCES tickers(TickerSymbol),
    coverage varchar -- n = only ohcvls, a = accounting docs and ohcvls, i = intraday ohcvls
);

CREATE TABLE ApiCalls (
    date date,
    CallCount int,
    datasource int REFERENCES datasources(SourceID)
);