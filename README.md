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
    - how to check if insert was successful? Check before removing from jobqueue
- go CLI get basic info from db?
4
- manage secrets better
- bash install script
    - build go cli
    - docker-compose build and up -d (how to deal with different environments?)
- add testing
- add forex to jobqueue etc.
    - use this format: JPY:EUR
- deploy to oracle, turso
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

- remove nulls e.g. in commodities:   4 | 0001-01-01 | 0.000000000000 |          1
- corn etf and corn commodity conflict, same ticker (query builder takes corn to declare a func...)
    - but are they different?
    - research how BRENT and BNO, WHEAT and EAT, SUGAR and CANE, ALUMINIUM and JJUFF etc. differ
- how is e.g. (Sector) xle different from (commodity etf?) dbe

### Usage:
- `sudo docker-compose build`
- `sudo docker-compose up -d`

- *go build*

UPDATE tickers
SET lastupdated = current_date - interval '2 months'
WHERE tickerid NOT IN (SELECT DISTINCT tickerid FROM commodities);

UPDATE tickers
SET lastupdated = current_date - interval '2 months'
WHERE tickersymbol = 'MT';

SELECT DISTINCT TickerID FROM commodities;

sudo docker exec -it pgsql-dev bash

### Implementation
- Technologies:
    - Go
    - Docker-compose
    - Postgresql



    