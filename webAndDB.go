package main

import (
	"FinanceSDK/e"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// this file touches the DB
// both inserts and retrievals

// the DB
const (
	host     = "127.0.0.1" // "jsontosql_db_1" // "127.0.0.1" // use docker name if from docker, ip if not in container!
	port     = 5432
	user     = "postgres"
	password = "password2"
	dbname   = "financial_markets"
)

func UpdateLastUpdated(db *sql.DB, ticker string) {
	rows, err := db.Query("SELECT TickerSymbol FROM jobqueue") // #todo add coverage to all of this
	e.Check(err)
	defer rows.Close()
	var isthishere string
	for rows.Next() {
		if err := rows.Scan(&isthishere); err != nil {
			panic("uh oh")
		}
		if isthishere == ticker {
			_, err := db.Exec(`UPDATE tickers SET lastupdated = $1 WHERE tickersymbol = $2`,
				time.Now(), ticker)
			e.Check(err)
			_, err = db.Exec(`DELETE FROM jobqueue WHERE tickersymbol = $1`, ticker)
			e.Check(err)
		}
	}

}

func RemoveFromQueryQueue(db *sql.DB, QueryTicker string) {

	_, err := db.Exec(`DELETE FROM queryqueue WHERE QueryTicker = $1`, QueryTicker)
	e.Check(err)

}

func AddTickersToDB(db *sql.DB, jobs []Job) {
	err := db.Ping()
	e.Check(err)

	for _, job := range jobs {
		_, err = db.Exec(`INSERT INTO tickers (TickerSymbol, updatefrequency, lastupdated)
	VALUES ($1, $2, current_date - interval '2 months')`,
			job.TickerSymbol, job.updatefrequency)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				fmt.Println("Duplicate value not saved")
			} else {
				panic(err)
			}
		}
	}
}

func GetJobQueue(db *sql.DB) ([]Job, error) {
	rows, err := db.Query("SELECT TickerID, TickerSymbol, coverage FROM jobqueue") // #todo add coverage to all of this
	e.Check(err)
	defer rows.Close()
	var jobQueue []Job
	for rows.Next() {
		var job Job
		if err := rows.Scan(&job.TickerID, &job.TickerSymbol, &job.coverage); err != nil {
			return nil, err
		}
		jobQueue = append(jobQueue, job)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return jobQueue, nil
}

// jobqueue funcs:
type Job struct {
	// kept in sql
	TickerSymbol    string
	TickerID        int
	updatefrequency string
	lastupdated     time.Time
	coverage        string
	// added by programs
	QueryTicker string // everything should populate this, tickersymbol for listed ticker
	giventype   Uploader
	url         string
}

func UpdateJobQueue(db *sql.DB) {
	// check if last updated is null
	// check if last updated is further back than the updatefrequency period
	rows, err := db.Query("SELECT TickerID, TickerSymbol, updatefrequency, lastupdated, coverage FROM tickers")
	e.Check(err)
	defer rows.Close()

	var updatequeue []Job

	for rows.Next() {
		var job Job
		if err := rows.Scan(&job.TickerID, &job.TickerSymbol, &job.updatefrequency, &job.lastupdated, &job.coverage); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}

		switch job.updatefrequency {
		case "m": // more than a month
			if time.Since(job.lastupdated).Hours() > 24*30 {
				updatequeue = append(updatequeue, job)
			}
		case "q": // more than 3 months
			if time.Since(job.lastupdated).Hours() > 24*30*3 { // #todo 90 days =/= 3 months... statements reported dates etc. sigh.
				updatequeue = append(updatequeue, job) // statements should also become unique or have primary key combos, and inserts do nothing
			}
		}
	}

	var failedlist string
	var added string
	for _, job := range updatequeue {
		_, err = db.Exec(`INSERT INTO jobqueue (TickerID, TickerSymbol, coverage, lastupdated)
	VALUES ($1, $2, $3, $4)`,
			job.TickerID, job.TickerSymbol, job.coverage, job.lastupdated)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				failedlist += job.TickerSymbol + " "
			} else {
				panic(err)
			}
		}
		added += job.TickerSymbol + " "
	}
	if failedlist != "" {
		fmt.Print("These are already on job queue: " + failedlist + "\n\n")
	}
	if added != "" {
		fmt.Print("Added to job queue: " + added + "\n")
	}
	if added == "" && failedlist == "" {
		fmt.Print("The job queue is empty. Everything is up to date.")
	}
}

func ReportApiCall(db *sql.DB) { // #todo make it stop at 500?
	result, err := db.Exec("UPDATE ApiCalls SET CallCount = CallCount + 1 WHERE date = CURRENT_DATE AND datasource = 1")
	e.Check(err)
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		_, err := db.Exec("INSERT INTO ApiCalls (date, CallCount, datasource) VALUES (CURRENT_DATE, 1, 1)")
		e.Check(err)
	}
}

func coverageToQueryQueue(db *sql.DB, jobs []Job) {
	var queryqueue []Job
	for _, oldjob := range jobs {
		switch oldjob.coverage {
		case "statements+dailyOHLCVs": // only stocks
			newJob := Job{ // OHLCVs & financial statements
				TickerSymbol: oldjob.TickerSymbol,
				TickerID:     oldjob.TickerID,
			}
			newJob.QueryTicker = oldjob.TickerSymbol
			queryqueue = append(queryqueue, newJob)
			// financial statements
			newJob2 := Job{ // OVERVIEW
				TickerSymbol: oldjob.TickerSymbol,
				TickerID:     oldjob.TickerID,
			}
			newJob2.QueryTicker = "OVERVIEW" + " " + oldjob.TickerSymbol
			queryqueue = append(queryqueue, newJob2)
			// actual financial statements only updated quarterly
			if time.Since(oldjob.lastupdated).Hours() > 24*30*3 { // #todo 90 days =/= 3 months...
				newJob3 := Job{ // INCOME STATEMENT
					TickerSymbol: oldjob.TickerSymbol,
					TickerID:     oldjob.TickerID,
				}
				newJob3.QueryTicker = "INCOME_STATEMENT" + " " + oldjob.TickerSymbol
				queryqueue = append(queryqueue, newJob3)
				newJob4 := Job{ // BALANCE_SHEET
					TickerSymbol: oldjob.TickerSymbol,
					TickerID:     oldjob.TickerID,
				}
				newJob4.QueryTicker = "BALANCE_SHEET" + " " + oldjob.TickerSymbol
				queryqueue = append(queryqueue, newJob4)
				newJob5 := Job{ // CASH_FLOW
					TickerSymbol: oldjob.TickerSymbol,
					TickerID:     oldjob.TickerID,
				}
				newJob5.QueryTicker = "CASH_FLOW" + " " + oldjob.TickerSymbol
				queryqueue = append(queryqueue, newJob5)
				newJob6 := Job{ // EARNINGS
					TickerSymbol: oldjob.TickerSymbol,
					TickerID:     oldjob.TickerID,
				}
				newJob6.QueryTicker = "EARNINGS" + " " + oldjob.TickerSymbol
				queryqueue = append(queryqueue, newJob6)
			}
		case "intradayOHLCVs+statements+dailyOHLCVs":
		case "daily":
			newJob := Job{ // ETFs, Commodities, Forex, macro stuff (n.b. type constraint in table tickers, should only have 'daily')
				TickerSymbol: oldjob.TickerSymbol,
				TickerID:     oldjob.TickerID,
			}
			newJob.QueryTicker = oldjob.TickerSymbol
			queryqueue = append(queryqueue, newJob)
		}
	}
	fmt.Println(queryqueue)

	var failedlist string
	var added string
	for _, query := range queryqueue {
		_, err := db.Exec(`INSERT INTO queryqueue (TickerID, TickerSymbol, QueryTicker)
	VALUES ($1, $2, $3)`, query.TickerID, query.TickerSymbol, query.QueryTicker)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				failedlist += query.TickerSymbol + ", "
			} else {
				panic(err)
			}
		}
		added += query.QueryTicker + ", "
	}
	if failedlist != "" {
		fmt.Print("These are already on query queue: " + failedlist + "\n\n")
	}
	if added != "" {
		fmt.Print("Added to query queue: " + added + "\n")
	}
	if added == "" && failedlist == "" {
		fmt.Print("The query queue is empty. Everything is up to date.")
	}
}

func GetQueryQueue(db *sql.DB) (queryqueue []Job, err error) {
	rows, err := db.Query("SELECT TickerID, TickerSymbol, QueryTicker FROM queryqueue") // #todo add coverage to all of this
	e.Check(err)
	defer rows.Close()
	for rows.Next() {
		var job Job
		if err := rows.Scan(&job.TickerID, &job.TickerSymbol, &job.QueryTicker); err != nil {
			return nil, err
		}
		queryqueue = append(queryqueue, job)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return queryqueue, nil

}
