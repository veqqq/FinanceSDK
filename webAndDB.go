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
		_, err = db.Exec(`INSERT INTO tickers (TickerSymbol, Importance, lastupdated)
	VALUES ($1, $2, current_date - interval '2 months')`,
			job.TickerSymbol, job.Importance)
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
	TickerSymbol string
	TickerID     int
	Importance   string
}

// #todo I can move to TickerID into this to send the job to others?
func GetJobQueue(db *sql.DB) ([]Job, error) {
	rows, err := db.Query("SELECT TickerID, TickerSymbol FROM jobqueue") // #todo add depth to all of this
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
	// check if last updated is further back than the importance period
	rows, err := db.Query("SELECT TickerID, TickerSymbol, importance, lastupdated FROM tickers")
	e.Check(err)
	defer rows.Close()

	var updatequeue []Job
	fmt.Print("Added to queue: ")

	for rows.Next() {
		var job Job
		var lastUpdated time.Time

		if err := rows.Scan(&job.TickerID, &job.TickerSymbol, &job.Importance, &lastUpdated); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}

		switch job.Importance {
		case "m":
			if time.Since(lastUpdated).Hours() > 24*30 { // more than a month
				// Add logic for updating job queue
				fmt.Printf("%s, ", job.TickerSymbol)
				updatequeue = append(updatequeue, job) // Add job to the slice
			}
		case "q":
			if time.Since(lastUpdated).Hours() > 24*30*3 { // more than a quarter
				// Add logic for updating job queue
				fmt.Printf("%s, ", job.TickerSymbol)
				updatequeue = append(updatequeue, job) // Add job to the slice
			}
		}
	}
	fmt.Print("\n\n")

	var failedlist string
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
	}
	fmt.Print("These are already on queue: " + failedlist + "\n\n")
}
