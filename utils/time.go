package utils

import "time"

const (
	layoutUS = "January 2, 2006"
)

func TimeToLayoutUS(timeString string) time.Time {
	t, _ := time.Parse(layoutUS, timeString)
	return t
}
