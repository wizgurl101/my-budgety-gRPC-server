package utils

import "time"

var GetCurrentDate = func() time.Time {
	return time.Now()
}

func GetFirstDateOfCurrentMonth() string {
	currentDate := GetCurrentDate()
	currentMonthFirstDate := time.Date(currentDate.Year(), currentDate.Month(), 1, 0, 0, 0, 0, time.UTC)
	return currentMonthFirstDate.Format("2006-01-02")
}

func GetFirstDateOfNextMonth() string {
	currentDate := GetCurrentDate()
	nextMonthFirstDate := time.Date(currentDate.Year(), currentDate.Month()+1, 1, 0, 0, 0, 0, time.UTC)
	return nextMonthFirstDate.Format("2006-01-02")
}
