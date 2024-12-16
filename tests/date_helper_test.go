package tests

import (
	"my-budgety-gRPC-server/utils"
	"testing"
	"time"
)

func TestGetLastDateOfCurrentMonth(t *testing.T) {
	t.Run("Given the current date Then the date of the last of the month is returned", func(t *testing.T) {
		original_get_current_date := utils.GetCurrentDate

		// Restore the original GetCurrentDate function after the test
		defer func() { utils.GetCurrentDate = original_get_current_date }()

		utils.GetCurrentDate = func() time.Time {
			return time.Date(2024, 12, 15, 0, 0, 0, 0, time.UTC)
		}

		expected := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
		actual := utils.GetLastDateOfCurrentMonth()
		if actual != expected {
			t.Errorf("expected %s but got %s", expected, actual)
		}
	})

	// t.Run("Given a leap year February Then the date of the last of the month is returned", func(t *testing.T) {
	// 	leapYear := 2020
	// 	expected := time.Date(leapYear, 3, 0, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
	// 	actual := utils.GetLastDateOfCurrentMonth()
	// 	if actual != expected {
	// 		t.Errorf("expected %s but got %s", expected, actual)
	// 	}
	// })

	// t.Run("Given a non-leap year February Then the date of the last of the month is returned", func(t *testing.T) {
	// 	nonLeapYear := 2021
	// 	expected := time.Date(nonLeapYear, 3, 0, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
	// 	actual := utils.GetLastDateOfCurrentMonth()
	// 	if actual != expected {
	// 		t.Errorf("expected %s but got %s", expected, actual)
	// 	}
	// })
}
