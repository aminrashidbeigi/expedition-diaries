package main

import (
	"database/sql"

	"example.com/history-travelers/storage/queries"
	_ "github.com/lib/pq"
)

func GetQueries() (*queries.Queries, error) {
	db, err := sql.Open("postgres", "user=postgres dbname=travels")
	if err != nil {
		return nil, err
	}
	queries_storage := queries.New(db)
	if err != nil {
		return nil, err
	}
	return queries_storage, nil
}
