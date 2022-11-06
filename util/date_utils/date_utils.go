package date_utils

import "time"

const (
	apiDateLayout = "01-02-2006T15:04:05Z"
)

// GetNow gets current time in UTC
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString converts current time into a string
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
