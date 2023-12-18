# FinanceSDK
-----

- early alpha

### Goals
- save and compare multivendor financial data
- provide workspace to value securities and model overall market

##### Project Structure
- DB container
    - postgres 
    - UpdaterDaemon
- Go cli uploads desired tickers to jobqueue in DB
- value companies with [CompanyModels](https://github.com/veqqq/CompanyModels) and explore data


ALTER TABLE tickers
ADD CONSTRAINT etf_type_constraint
CHECK (
  (type = 'macro' OR type = 'commodity' OR type LIKE '%ETF%') AND coverage = 'daily'
);

ETFs, Commodities, Forex, macro


### To Do:
- api fails a lot, even at 1 call/10 secons. For 13 tickers, 12 overviews, 11 balance and earnings, 10 income, 9 cash.
- check if alphavantage daily etc. are dividend adjusted and how on earth do people deal with that when new divies come, esp. since i like energy and brazil...
- instead of os.exit when json empty, try going to the next thing? but stop if e.g. 3? are empty in a row?
- check #todos in the code
- manage secrets better
- add testing
- go CLI get basic info from db? (how many rows etc.)
- deploy to oracle, turso
    - for oracle, make updaterdaemon which does case 2 of userinput options (updating only)
- split commodities in 2, some require more left hand precision, others right hand. Can reduce data by a fair bit
    - highest
        - vol in daily 853446200
        - val in commodities 158461.000000000000
            - -36.980000000000 lol i remember that day WTI
- decouple from alphavantage api
    - add more apis
    - test alignment between different sources
- optimize bulkinsert method for postgres https://stackoverflow.com/questions/12206600/how-to-speed-up-insertion-performance-in-postgresql
    - mostly for fun, would overload API here
    - also not relevant for [CompanyModels](https://github.com/veqqq/CompanyModels) which can't retrieve and insert many more valuations than ticks...
- bash install script
    - build go cli
    - docker-compose build and up -d (how to deal with different environments?)
    - my current usage:
        - docker-compose up
        - go run .

### Usage:
- `sudo docker-compose build`
- `sudo docker-compose up -d`

- sudo docker exec -it financesdk_db_1 bash

UPDATE tickers
SET lastupdated = current_date - interval '2 months'
WHERE tickerid NOT IN (SELECT DISTINCT tickerid FROM commodities);

UPDATE tickers
SET lastupdated = current_date - interval '2 months'
WHERE tickersymbol = 'ZION';

SELECT DISTINCT TickerID FROM commodities;

SELECT TickerSymbol
FROM tickers
WHERE type LIKE '%stock%'
AND NOT EXISTS (
    SELECT 1
    FROM dailyOHLCVs
    WHERE dailyOHLCVs.TickerID = tickers.TickerID
);

UPDATE tickers
SET lastupdated = current_date - interval '4 months'
WHERE (type LIKE '%stock%' or type LIKE '%ETF%')
AND NOT EXISTS (
    SELECT 1
    FROM dailyOHLCVs
    WHERE dailyOHLCVs.TickerID = tickers.TickerID
);

- .env file in /e like this:{"X-RapidAPI-Key":"apikey"}


ALTER TABLE tickers ADD CONSTRAINT check_tickers_type_coverage
CHECK (
  (type NOT LIKE '%stock%' AND coverage = 'daily') OR
  (type LIKE '%stock%' AND coverage <> 'daily')
);

UPDATE tickers
SET coverage = 'daily'
WHERE type LIKE '%ETF%' OR type = 'commodity' OR type = 'Forex' OR type = 'macro';

### Implementation
- Technologies:
    - Go
    - Docker-compose
    - Postgresql

- will containerize later
    - DB (backup layer)
        - sync DBs (a write writes to 2, others query DBs to see if they have more recent updates than themselves)
    - check DB queue and update tickers (or in DB container?)
    - not containerized:
        - CLI to update tickers/jobs
