package expanse

import (
	"context"
	"log"

	bigquery "cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"

	envmanager "my-budgety-gRPC-server/envmanager"
)

func GetAllExpanse() float32 {
	ctx := context.Background()
	env_variables := envmanager.GetEnvVariables()
	client, err := bigquery.NewClient(ctx, env_variables.ProjectId)
	if err != nil {
		log.Fatalf("failed to create bigquery client: %v", err)
	}

	query_str :=
		`SELECT SUM(amount) ` +
			`FROM ` + env_variables.ProjectId + `.` + env_variables.ProjectName + `.expanse ` +
			`WHERE date >= @firstOfMonthDate
		AND date <= @endOfMonthDate`
	q := client.Query(query_str)

	q.Parameters = []bigquery.QueryParameter{
		{Name: "firstOfMonthDate", Value: "2024-12-01"},
		{Name: "endOfMonthDate", Value: "2024-12-31"},
	}

	it, err := q.Read(ctx)
	if err != nil {
		log.Fatalf("failed to read query results: %v", err)
	}

	var amount float64

	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("failed to iterate query results: %v", err)
			break
		}

		amount = values[0].(float64)
	}

	return float32(amount)
}
