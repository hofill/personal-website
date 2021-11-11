package utils

import (
	"log"
	"time"
)

const (
	layoutUS = "January 2, 2006"
)

func TimeToLayoutUS(timeString string) time.Time {
	t, err := time.Parse(layoutUS, timeString)
	if err != nil {
		log.Println(err)
	}
	return t
}
