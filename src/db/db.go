package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

const DB_USER string = "postgres"
const DB_PASSWORD string = "228p_b"
const DB_PORT string = "5432"
const DB_NAME string = "ozon"

func Connect() {
	// str := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", DB_USER, DB_PASSWORD, DB_PORT, DB_NAME)
	var err error
	pool, err = pgxpool.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", DB_USER, DB_PASSWORD, DB_PORT, DB_NAME))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
}
func ClosePool() {
	pool.Close()
}
func GetPool() *pgxpool.Pool {
	return pool
}
