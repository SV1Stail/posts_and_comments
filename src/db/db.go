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

func GetTables() ([]string, error) {
	// Получаем пул соединений
	conn := GetPool()

	// Запрос для получения списка таблиц
	query := `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public'
	`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("unable to execute query: %v", err)
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, fmt.Errorf("unable to scan row: %v", err)
		}
		tables = append(tables, tableName)
	}

	// Проверка на ошибки после завершения цикла
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %v", rows.Err())
	}

	return tables, nil
}

