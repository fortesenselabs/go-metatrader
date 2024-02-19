package utils

import "time"

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

func Contains(slice []string, item string) bool {
	// Check if slice contains item
	return false
}
