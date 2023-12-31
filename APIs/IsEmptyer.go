package APIs

import (
	"fmt"
	"os"
)

// Checks if JSON received from site was empty. Applies if:
// - ticker not supported by vendor
// - vendor is querried too fast and does not supply results
// IsEmptyer is implemented by the wrapper structs w/ original json format
func CheckJSON(err error, json IsEmptyer) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	if json.IsEmpty() {
		fmt.Println("The decoded JSON is empty or contains an empty map.")
		fmt.Println(json)
		// os.Exit(-1)  // #todo this trips falsely on quarterly?
	}
}

type IsEmptyer interface {
	IsEmpty() bool
}

func (c FinancialStatements) IsEmpty() bool { // this trips falsely on quarterly?
	return len(c.QuarterlyReports) == 0
}

type FinancialStatements struct { // statements embed this struct
	QuarterlyReports []any // satisfying the interface
}

func (c StockOverview) IsEmpty() bool {
	return len(c.Symbol) == 0
}

func (c CommodityPrices) IsEmpty() bool {
	return len(c.Data) == 0
}

func (d DailyOHLCVs) IsEmpty() bool {
	return len(d.TimeSeries) == 0
}

func (d IntradayOHLCVs) IsEmpty() bool {
	return len(d.TimeSeries1min) == 0
}

func (d ForexPrices) IsEmpty() bool {
	return len(d.TimeSeriesFX) == 0
}
