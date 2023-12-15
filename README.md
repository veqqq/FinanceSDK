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


### To Do:
- check #todos in the code

- track total API calls each day

- manage secrets better
- add testing

- go CLI get basic info from db? (how many rows etc.)
- change "importance" to "frequency"
- add something like "coverage" to determine if you should get dailies, intraday, financial docs...
    - makes confirming something was successfully updated difficult, must check many things. How to deal with failure midway?
    - intraday should only be gotten through this, entering e.g. "CLF 2021-04" in CLI will make a new ticker wih that date...
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

sudo docker exec -it financesdk_db_1 bash

- .env file in /e like this:{"X-RapidAPI-Key":"apikey"}


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
