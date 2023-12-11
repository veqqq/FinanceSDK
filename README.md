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

- Postgres
    - postgres access funcs
    - finish schema
- add job queue to sql table
- go CLI to add to job queue (and get basic info from db?)
- go program to update based on job queue
    - check lastupdateds to populate jobqueue
    - then update
- manage secrets better

- bash install script
    - build go cli
    - docker-compose build and up -d (how to deal with different environments?)
- add testing

- deploy to oracle, turso

- decouple from alphavantage api
    - add more apis
    - test alignment between different sources

Usage:
- `sudo docker-compose build`
- `sudo docker-compose up -d`

- *go build*

### Implementation
- Technologies:
    - Go
    - Docker-compose
    - Postgresql



    