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
	_, err := db.Exec(`
    UPDATE tickers
    SET lastupdated = $1
    WHERE tickersymbol = $2
`, time.Now(), ticker)
	e.Check(err)
}

func RemoveFromJobQueue(db *sql.DB, ticker string) {
	_, err := db.Exec(`DELETE FROM JobQueue WHERE TickerSymbol = $1`, ticker)
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

// jobqueue funcs:
type Job struct {
	TickerSymbol    string
	TickerID        int
	updatefrequency string
}

func GetJobQueue(db *sql.DB) ([]Job, error) {
	rows, err := db.Query("SELECT TickerID, TickerSymbol FROM jobqueue") // #todo add coverage to all of this
	e.Check(err)
	defer rows.Close()
	var jobQueue []Job
	for rows.Next() {
		var job Job
		if err := rows.Scan(&job.TickerID, &job.TickerSymbol); err != nil {
			return nil, err
		}
		jobQueue = append(jobQueue, job)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return jobQueue, nil
}

func UpdateJobQueue(db *sql.DB) {
	// check if last updated is null
	// check if last updated is further back than the updatefrequency period
	rows, err := db.Query("SELECT TickerID, TickerSymbol, updatefrequency, lastupdated FROM tickers")
	e.Check(err)
	defer rows.Close()

	var updatequeue []Job

	for rows.Next() {
		var job Job
		var lastUpdated time.Time

		if err := rows.Scan(&job.TickerID, &job.TickerSymbol, &job.updatefrequency, &lastUpdated); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}

		switch job.updatefrequency {
		case "m": // more than a month
			if time.Since(lastUpdated).Hours() > 24*30 {
				updatequeue = append(updatequeue, job)
			}
		case "q": // more than 3 months
			if time.Since(lastUpdated).Hours() > 24*30*3 {
				updatequeue = append(updatequeue, job)
			}
		}
	}

	var failedlist string
	var added string
	for _, job := range updatequeue {
		_, err = db.Exec(`INSERT INTO jobqueue (TickerID, TickerSymbol)
	VALUES ($1, $2)`,
			job.TickerID, job.TickerSymbol)
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
		fmt.Print("These are already on queue: " + failedlist + "\n\n")
	}
	if added != "" {
		fmt.Print("Added to queue: " + added + "\n")
	}
	if added == "" && failedlist == "" {
		fmt.Print("The job queue is empty. Everything is up to date.")
	}
}

func ReportApiCall(db *sql.DB) {
	result, err := db.Exec("UPDATE ApiCalls SET CallCount = CallCount + 1 WHERE date = CURRENT_DATE")
	e.Check(err)
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		_, err := db.Exec("INSERT INTO ApiCalls (date, CallCount, datasource) VALUES (CURRENT_DATE, 1, 1)")
		e.Check(err)
	}
}
