package utils

import (
	"strconv"
	"time"
)

func ConvertDateToUTC(dateStr string, _format string) (string, error) {
	// Parse the input date string based on the given format
	date, err := time.Parse(_format, dateStr)
	if err != nil {
		return "", err
	}

	// Convert the date to UTC
	utcDate := date.UTC()

	// Format the UTC date as a string
	utcDateStr := utcDate.Format(_format)

	return utcDateStr, nil
}

// func ConvertDataFrameSeries(dfSereis series.Series) []string {
// 	// Create an array of strings
// 	var data []string

// 	// Convert the dataframe series to a string array
// 	for _, value := range dfSereis.Values {
// 		data = append(data, value.(string))
// 	}

// 	return data
// }

// Contains checks if a string is present in a slice
func Contains(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

// Helper function to parse float64 from string
func ParseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

// Helper function to parse int from string
func ParseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func ParseInt64(s string) int64 {
	i, _ := strconv.Atoi(s)
	return int64(i)
}
