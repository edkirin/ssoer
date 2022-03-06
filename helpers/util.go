package helpers

import "time"

func GetUTCNow() time.Time {
	utc, _ := time.LoadLocation("UTC")
	return time.Now().In(utc)
}
