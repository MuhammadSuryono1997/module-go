package utils

import (
	"fmt"
	"time"
)

func TimeStamp() string {
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	return formatted
}

func CompareTime(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func TimeNow() string {
	return time.Now().Format(time.RFC3339)
}

func TimeAdd(minute time.Duration) string {
	return time.Now().Add(minute).Format(time.RFC3339)
}
