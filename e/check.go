package e

import (
	"fmt"
	"os"
)

func Check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}

// CheckDBInsert(err, ticker, structType, m)
func CheckDBInsert(err error, ticker, typeOfStructType string, m any) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Println(m)
		fmt.Println(ticker + " experienced an error")
		fmt.Println(typeOfStructType + " indicates intended sql table")
		fmt.Println("End of Error Message")
		os.Exit(-1)
	}
}
