package helpers

import "time"

func TimeFormat(time time.Time) string {
	return time.Format("2006-02-01 01:02:03 PM")
}

func DateFormat(time time.Time) string {
	return time.Format("2006-02-01")
}
