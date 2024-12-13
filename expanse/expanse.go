package expanse

import (
	"context"
	"log"
	"os"

	bigquery "cloud.google.com/go/bigquery"
)

func getAllExpanse() {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("failed to create bigquery client: %v", err)
	}

	q := client.Query(`
		SELECT SUM(amount) 
		FROM 'my-budgety-1.expanse.expanse'
		WHERE date >= @firstOfMonthDate
		AND date <= @endOfMonthDate
	`)

	q.Parameters = []bigquery.QueryParameter{
		{Name: "firstOfMonthDate", Value: "2024-12-01"},
		{Name: "endOfMonthDate", Value: "2024-12-31"},
	}
}
