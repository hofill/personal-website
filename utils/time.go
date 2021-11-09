package utils

import "time"

func GetDateFormatted(t time.Time) string {
	// format time by us layout, january 2nd 2006
	return t.Format("January 2, 2006")
}
