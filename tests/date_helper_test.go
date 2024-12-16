package tests

import (
	"my-budgety-gRPC-server/utils"
	"testing"
	"time"
)

func TestGetFirstDateOfNextMonth(t *testing.T) {
	t.Run("Given the date 2024-12-15 Then the date of the first of the next month is returned", func(t *testing.T) {
		original_get_current_date := utils.GetCurrentDate

		// Restore the original GetCurrentDate function after the test
		defer func() { utils.GetCurrentDate = original_get_current_date }()

		utils.GetCurrentDate = func() time.Time {
			return time.Date(2024, 12, 15, 0, 0, 0, 0, time.UTC)
		}

		expected := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
		actual := utils.GetFirstDateOfNextMonth()
		if actual != expected {
			t.Errorf("expected %s but got %s", expected, actual)
		}
	})
}

func TestGetFirstDateOfCurrentMonth(t *testing.T) {
	t.Run("Given the date 2024-12-15 Then the date of the first of the month is returned", func(t *testing.T) {
		original_get_current_date := utils.GetCurrentDate

		// Restore the original GetCurrentDate function after the test
		defer func() { utils.GetCurrentDate = original_get_current_date }()

		utils.GetCurrentDate = func() time.Time {
			return time.Date(2024, 12, 15, 0, 0, 0, 0, time.UTC)
		}

		expected := time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
		actual := utils.GetFirstDateOfCurrentMonth()
		if actual != expected {
			t.Errorf("expected %s but got %s", expected, actual)
		}
	})
}
