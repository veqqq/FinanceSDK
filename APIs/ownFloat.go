package APIs

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

// Custom Float64 to handle nulls etc. when unmarshaling
type ownFloat64 struct {
	Val   float64
	Valid bool
}

func (of *ownFloat64) UnmarshalJSON(data []byte) error {
	var rawValue interface{}
	if err := json.Unmarshal(data, &rawValue); err != nil {
		return err
	}

	switch v := rawValue.(type) {
	case float64:
		of.Val = v
		of.Valid = true
	case string:
		if v == "None" {
			of.Valid = false
		} else if v == "-" {
			of.Valid = false

		} else {
			value, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return err
			}
			of.Val = value
			of.Valid = true
		}
	default:
		return fmt.Errorf("unexpected value type for ownFloat64: %T", v)
	}

	return nil
}

func (of ownFloat64) MarshalJSON() ([]byte, error) {
	if !of.Valid {
		return json.Marshal(0)
	}
	return json.Marshal(fmt.Sprint(int64(of.Val)))
}

func (of ownFloat64) Value() (driver.Value, error) { // sql db
	return float64(of.Val), nil
}

func (of *ownFloat64) Scan(value interface{}) error {
	switch v := value.(type) {
	case float64:
		of.Val = v
		of.Valid = true
	case []byte:
		f, err := strconv.ParseFloat(string(v), 64)
		if err != nil {
			return err
		}
		of.Val = f
		of.Valid = true
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
	return nil
}
