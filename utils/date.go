package utils

import "time"

const (
	dbDateTimeLayout = "2006-01-02 15:04:05"
	apiDateLayout    = "2006-01-02T15:04:05Z"
)

//GetNow returns new time object
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowString gets a standard time string
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

//GetNowDBFormat gets a DB time string
func GetNowDBFormat() string {
	return GetNow().Format(dbDateTimeLayout)
}
