package e

import (
	"database/sql"
	"fmt"
	"os"
)

func Check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}

// e.CheckDBInsert(err, result, ticker, structType, m)
func CheckDBInsert(err error, result sql.Result, ticker, typeOfStructType string, m any) {
	if err != nil {
		fmt.Println(m)
		fmt.Fprintln(os.Stderr, err)
		fmt.Println(ticker + " experienced an error")
		fmt.Println(typeOfStructType + " indicates intended sql table")

		fmt.Println("End of Error Message")
		os.Exit(-1)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// Handle the error, log it, and potentially return an error response
		fmt.Println(ticker + " experienced an error inserting rows")
		fmt.Println("Error getting rows affected: ", err)
		fmt.Println(typeOfStructType + " indicates intended sql table")
	}

	// Check if any rows were affected
	if rowsAffected == 0 {
		// No rows were affected, which might indicate an issue
		fmt.Println(m)
		fmt.Println(ticker + " experienced an error inserting rows")
		fmt.Println("No rows affected after insert")
		fmt.Println(typeOfStructType + " indicates intended sql table")
	}
}
