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

- manage secrets better
- bash install script
    - build go cli
    - docker-compose build and up -d (how to deal with different environments?)
- add testing
- add forex to jobqueue etc.

- deploy to oracle, turso

- decouple from alphavantage api
    - add more apis
    - test alignment between different sources
- optimize bulkinsert method for postgres https://stackoverflow.com/questions/12206600/how-to-speed-up-insertion-performance-in-postgresql

- remove nulls e.g. in commodities:   4 | 0001-01-01 | 0.000000000000 |          1

Usage:
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



    