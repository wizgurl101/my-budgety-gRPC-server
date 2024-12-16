package utils

import "time"

var GetCurrentDate = func() time.Time {
	return time.Now()
}

func GetLastDateOfCurrentMonth() string {
	currentDate := GetCurrentDate()
	nextMonthFirstDate := time.Date(currentDate.Year(), currentDate.Month()+1, 1, 0, 0, 0, 0, time.UTC)
	lastDateOfCurrentMonth := nextMonthFirstDate.AddDate(0, -1, -1)
	return lastDateOfCurrentMonth.Format("2006-01-02")
}
